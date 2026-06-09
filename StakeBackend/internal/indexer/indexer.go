package indexer

import (
	"StakeBackend/internal/abi"
	"StakeBackend/internal/config"
	"StakeBackend/internal/contract"
	"StakeBackend/internal/db"
	"StakeBackend/internal/model"
	"StakeBackend/internal/pkg/logger"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	eventQueue    []EventBatchItem
	adminQueue    []AdminBatchItem
	blockMap      map[uint64]model.BlockSync
	queueMutex    sync.Mutex
	lastFlushTime time.Time
	batchSize     int
	batchWaitTime time.Duration
)

type EventBatchItem struct {
	ChainEvent        model.ChainEvent
	UserAddr          string
	Amount            *big.Int
	IsStake           bool
	IsReward          bool
	Reward            *big.Int
	Account           string
	MinLimit          *big.Int
	MaxLimit          *big.Int
	IsBlacklistAdd    bool
	IsBlacklistRemove bool
	IsStakeLimits     bool
}

type AdminBatchItem struct {
	AdminOp           model.AdminOperation
	Account           string
	MinLimit          *big.Int
	MaxLimit          *big.Int
	IsBlacklistAdd    bool
	IsBlacklistRemove bool
	IsStakeLimits     bool
}

func StartIndexer() {
	logger.Logger.Info("索引器启动")

	batchSize = config.GlobalConfig.Indexer.BatchSize
	batchWaitTime = time.Duration(config.GlobalConfig.Indexer.BatchWaitSeconds) * time.Second
	lastFlushTime = time.Now()
	eventQueue = make([]EventBatchItem, 0, batchSize)
	adminQueue = make([]AdminBatchItem, 0, batchSize)
	blockMap = make(map[uint64]model.BlockSync)

	currentBlock := config.GlobalConfig.Chain.StartBlock
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	go batchFlushTicker()

	for range ticker.C {
		latestBlock, err := contract.Client.BlockNumber(contract.Ctx)
		if err != nil {
			logger.Logger.Error("获取最新区块失败", zap.Error(err))
			continue
		}

		for currentBlock <= latestBlock {
			if err := handleReorg(currentBlock); err != nil {
				logger.Logger.Error("链重组处理失败", zap.Uint64("block", currentBlock), zap.Error(err))
				break
			}

			if err := parseEventsToQueue(currentBlock); err != nil {
				logger.Logger.Error("事件解析入队失败", zap.Uint64("block", currentBlock), zap.Error(err))
				break
			}

			currentBlock++
			config.GlobalConfig.Chain.StartBlock = currentBlock
		}
	}
}

func batchFlushTicker() {
	checkTicker := time.NewTicker(1 * time.Second)
	defer checkTicker.Stop()

	for range checkTicker.C {
		queueMutex.Lock()
		if time.Since(lastFlushTime) > batchWaitTime && (len(eventQueue) > 0 || len(adminQueue) > 0) {
			logger.Logger.Info("批量入库触发：超时自动提交")
			flushBatch()
		}
		queueMutex.Unlock()
	}
}

func parseEventsToQueue(blockNumber uint64) error {
	block, err := contract.Client.BlockByNumber(contract.Ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return err
	}

	blockMap[blockNumber] = model.BlockSync{
		BlockNumber: blockNumber,
		BlockHash:   block.Hash().Hex(),
		ParentHash:  block.ParentHash().Hex(),
	}

	contractAddr := common.HexToAddress(config.GlobalConfig.Chain.MiningProxy)
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(blockNumber),
		ToBlock:   new(big.Int).SetUint64(blockNumber),
		Addresses: []common.Address{contractAddr},
	}
	logs, err := contract.Client.FilterLogs(contract.Ctx, query)
	if err != nil {
		return err
	}

	queueMutex.Lock()
	defer queueMutex.Unlock()

	for _, vLog := range logs {
		item := parseEventToBatchItem(vLog)
		if item != nil {
			if item.IsAdminEvent() {
				adminQueue = append(adminQueue, item.ToAdminBatchItem(vLog))
			} else {
				eventQueue = append(eventQueue, *item)
			}
		}
	}

	if len(eventQueue) >= batchSize || len(adminQueue) >= batchSize {
		logger.Logger.Info("批量入库触发：数量达到阈值", zap.Int("user_events", len(eventQueue)), zap.Int("admin_events", len(adminQueue)))
		flushBatch()
	}

	return nil
}

func (item *EventBatchItem) IsAdminEvent() bool {
	return item.IsBlacklistAdd || item.IsBlacklistRemove || item.IsStakeLimits
}

func (item *EventBatchItem) ToAdminBatchItem(vLog types.Log) AdminBatchItem {
	op := model.AdminOperation{
		BlockNumber: vLog.BlockNumber,
		BlockHash:   vLog.BlockHash.Hex(),
		TxHash:      vLog.TxHash.Hex(),
		EventTime:   vLog.BlockNumber,
	}

	if item.IsBlacklistAdd {
		op.EventType = "BlacklistAdded"
		op.TargetUser = item.Account
	} else if item.IsBlacklistRemove {
		op.EventType = "BlacklistRemoved"
		op.TargetUser = item.Account
	} else if item.IsStakeLimits {
		op.EventType = "StakeLimitsUpdated"
		op.MinLimit = item.MinLimit.String()
		op.MaxLimit = item.MaxLimit.String()
	}

	return AdminBatchItem{
		AdminOp:           op,
		Account:           item.Account,
		MinLimit:          item.MinLimit,
		MaxLimit:          item.MaxLimit,
		IsBlacklistAdd:    item.IsBlacklistAdd,
		IsBlacklistRemove: item.IsBlacklistRemove,
		IsStakeLimits:     item.IsStakeLimits,
	}
}

func parseEventToBatchItem(vLog types.Log) *EventBatchItem {
	switch vLog.Topics[0].Hex() {
	case abi.AbiStakedEventID.Hex():
		event, _ := contract.MiningContract.ParseStaked(vLog)
		return &EventBatchItem{
			ChainEvent: model.ChainEvent{
				BlockNumber: vLog.BlockNumber,
				BlockHash:   vLog.BlockHash.Hex(),
				TxHash:      vLog.TxHash.Hex(),
				EventType:   "Staked",
				User:        event.User.Hex(),
				Amount:      event.Amount.String(),
				EventTime:   event.Time.Uint64(),
			},
			UserAddr: event.User.Hex(),
			Amount:   event.Amount,
			IsStake:  true,
		}

	case abi.AbiUnstakedEventID.Hex():
		event, _ := contract.MiningContract.ParseUnstaked(vLog)
		return &EventBatchItem{
			ChainEvent: model.ChainEvent{
				BlockNumber: vLog.BlockNumber,
				BlockHash:   vLog.BlockHash.Hex(),
				TxHash:      vLog.TxHash.Hex(),
				EventType:   "Unstaked",
				User:        event.User.Hex(),
				Amount:      event.Amount.String(),
				EventTime:   event.Time.Uint64(),
			},
			UserAddr: event.User.Hex(),
			Amount:   event.Amount,
			IsStake:  false,
		}

	case abi.AbiRewardClaimedEventID.Hex():
		event, _ := contract.MiningContract.ParseRewardClaimed(vLog)
		return &EventBatchItem{
			ChainEvent: model.ChainEvent{
				BlockNumber: vLog.BlockNumber,
				BlockHash:   vLog.BlockHash.Hex(),
				TxHash:      vLog.TxHash.Hex(),
				EventType:   "RewardClaimed",
				User:        event.User.Hex(),
				Amount:      event.Reward.String(),
				EventTime:   event.Time.Uint64(),
			},
			UserAddr: event.User.Hex(),
			IsReward: true,
			Reward:   event.Reward,
		}

	case abi.AbiBlacklistAddedEventID.Hex():
		event, _ := contract.MiningContract.ParseBlacklistAdded(vLog)
		return &EventBatchItem{
			ChainEvent: model.ChainEvent{
				BlockNumber: vLog.BlockNumber,
				BlockHash:   vLog.BlockHash.Hex(),
				TxHash:      vLog.TxHash.Hex(),
				EventType:   "BlacklistAdded",
				User:        event.Account.Hex(),
				EventTime:   vLog.BlockNumber,
			},
			Account:        event.Account.Hex(),
			IsBlacklistAdd: true,
		}

	case abi.AbiBlacklistRemovedEventID.Hex():
		event, _ := contract.MiningContract.ParseBlacklistRemoved(vLog)
		return &EventBatchItem{
			ChainEvent: model.ChainEvent{
				BlockNumber: vLog.BlockNumber,
				BlockHash:   vLog.BlockHash.Hex(),
				TxHash:      vLog.TxHash.Hex(),
				EventType:   "BlacklistRemoved",
				User:        event.Account.Hex(),
				EventTime:   vLog.BlockNumber,
			},
			Account:           event.Account.Hex(),
			IsBlacklistRemove: true,
		}

	case abi.AbiStakeLimitsUpdatedEventID.Hex():
		event, _ := contract.MiningContract.ParseStakeLimitsUpdated(vLog)
		return &EventBatchItem{
			ChainEvent: model.ChainEvent{
				BlockNumber: vLog.BlockNumber,
				BlockHash:   vLog.BlockHash.Hex(),
				TxHash:      vLog.TxHash.Hex(),
				EventType:   "StakeLimitsUpdated",
				Amount:      event.Min.String(),
				ExtraData:   event.Max.String(),
				EventTime:   vLog.BlockNumber,
			},
			MinLimit:      event.Min,
			MaxLimit:      event.Max,
			IsStakeLimits: true,
		}
	}
	return nil
}

func flushBatch() {
	if len(eventQueue) == 0 && len(adminQueue) == 0 {
		return
	}

	var events []model.ChainEvent
	var adminOps []model.AdminOperation
	var blockSyncs []model.BlockSync

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		events = make([]model.ChainEvent, 0, len(eventQueue))
		adminOps = make([]model.AdminOperation, 0, len(adminQueue))
		blockSyncs = make([]model.BlockSync, 0, len(blockMap))

		for _, item := range eventQueue {
			events = append(events, item.ChainEvent)
		}

		for _, item := range adminQueue {
			adminOps = append(adminOps, item.AdminOp)
		}

		for _, b := range blockMap {
			blockSyncs = append(blockSyncs, b)
		}

		if len(events) > 0 {
			if err := tx.CreateInBatches(events, 50).Error; err != nil {
				return err
			}
		}

		if len(adminOps) > 0 {
			if err := tx.CreateInBatches(adminOps, 50).Error; err != nil {
				return err
			}
		}

		if len(blockSyncs) > 0 {
			if err := tx.CreateInBatches(blockSyncs, 50).Error; err != nil {
				return err
			}
		}

		batchUpdateUserStake(tx, eventQueue)

		return nil
	})

	if err != nil {
		logger.Logger.Error("入库失败", zap.Error(err))
		return
	}

	eventQueue = make([]EventBatchItem, 0, batchSize)
	adminQueue = make([]AdminBatchItem, 0, batchSize)
	clear(blockMap)
	lastFlushTime = time.Now()

	logger.Logger.Info("入库成功", zap.Int("user_events", len(events)), zap.Int("admin_events", len(adminOps)))
}

func batchUpdateUserStake(tx *gorm.DB, queue []EventBatchItem) {
	userMap := make(map[string]*model.UserStake)

	for _, item := range queue {
		if item.IsReward {
			continue
		}
		if _, ok := userMap[item.UserAddr]; !ok {
			var stake model.UserStake
			tx.Where("user_address = ?", item.UserAddr).First(&stake)
			userMap[item.UserAddr] = &stake
		}

		stake := userMap[item.UserAddr]
		amount := item.Amount

		if stake.ID == 0 {
			if item.IsStake {
				stake.UserAddress = item.UserAddr
				stake.StakeAmount = amount.String()
				stake.UpdateTime = uint64(time.Now().Unix())
				tx.Create(stake)
			}
		} else {
			current := new(big.Int)
			current.SetString(stake.StakeAmount, 10)
			if item.IsStake {
				current.Add(current, amount)
			} else {
				current.Sub(current, amount)
			}
			stake.StakeAmount = current.String()
			stake.UpdateTime = uint64(time.Now().Unix())
			tx.Save(stake)
		}
	}
}

func handleReorg(blockNumber uint64) error {
	var syncBlock model.BlockSync
	if err := db.DB.Where("block_number = ?", blockNumber).First(&syncBlock).Error; err != nil {
		return nil
	}

	chainBlock, err := contract.Client.BlockByNumber(contract.Ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return err
	}

	if syncBlock.BlockHash != chainBlock.Hash().Hex() {
		logger.Logger.Warn("检测到链重组，回滚数据", zap.Uint64("block", blockNumber))

		queueMutex.Lock()
		eventQueue = make([]EventBatchItem, 0, batchSize)
		adminQueue = make([]AdminBatchItem, 0, batchSize)
		clear(blockMap)
		queueMutex.Unlock()

		db.DB.Where("block_number >= ?", blockNumber).Delete(&model.ChainEvent{})
		db.DB.Where("block_number >= ?", blockNumber).Delete(&model.AdminOperation{})
		db.DB.Where("block_number >= ?", blockNumber).Delete(&model.BlockSync{})
		config.GlobalConfig.Chain.StartBlock = blockNumber
	}
	return nil
}

func FlushOnShutdown() {
	queueMutex.Lock()
	defer queueMutex.Unlock()
	logger.Logger.Info("服务关闭，强制刷新批量队列")
	flushBatch()
}
