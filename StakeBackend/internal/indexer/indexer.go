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

// ==================== 批量队列全局变量（线程安全）====================
var (
	eventQueue    []EventBatchItem           // 事件批量队列
	blockMap      map[uint64]model.BlockSync // 区块同步记录
	queueMutex    sync.Mutex
	lastFlushTime time.Time     // 上次批量入库时间
	batchSize     int           // 批量大小
	batchWaitTime time.Duration // 批量等待时间
)

// EventBatchItem 批量队列存储结构
type EventBatchItem struct {
	ChainEvent model.ChainEvent
	UserAddr   string
	Amount     *big.Int
	IsStake    bool
	IsReward   bool
	Reward     *big.Int
}

// ==================== 启动索引器（批量版）====================
func StartIndexer() {
	logger.Logger.Info("索引器启动")

	// 初始化参数
	batchSize = config.GlobalConfig.Indexer.BatchSize
	batchWaitTime = time.Duration(config.GlobalConfig.Indexer.BatchWaitSeconds) * time.Second
	lastFlushTime = time.Now()
	eventQueue = make([]EventBatchItem, 0, batchSize)
	blockMap = make(map[uint64]model.BlockSync)

	currentBlock := config.GlobalConfig.Chain.StartBlock
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	// 独立协程：定时检查批量入库条件（超时触发）
	go batchFlushTicker()

	// 主循环：同步区块 + 解析事件入队列
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

			// 解析事件 → 加入批量队列
			if err := parseEventsToQueue(currentBlock); err != nil {
				logger.Logger.Error("事件解析入队失败", zap.Uint64("block", currentBlock), zap.Error(err))
				break
			}

			currentBlock++
			config.GlobalConfig.Chain.StartBlock = currentBlock
		}
	}
}

// ==================== 定时检查：超时自动批量入库 ====================
func batchFlushTicker() {
	checkTicker := time.NewTicker(1 * time.Second)
	defer checkTicker.Stop()

	for range checkTicker.C {
		queueMutex.Lock()
		// 触发条件1：超时无写入
		if time.Since(lastFlushTime) > batchWaitTime && len(eventQueue) > 0 {
			logger.Logger.Info("批量入库触发：超时自动提交")
			flushBatch()
		}
		queueMutex.Unlock()
	}
}

// ==================== 解析事件 → 加入批量队列 ====================
func parseEventsToQueue(blockNumber uint64) error {
	block, err := contract.Client.BlockByNumber(contract.Ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return err
	}

	// 记录区块（用于链重组校验）
	blockMap[blockNumber] = model.BlockSync{
		BlockNumber: blockNumber,
		BlockHash:   block.Hash().Hex(),
		ParentHash:  block.ParentHash().Hex(),
	}

	// 获取合约事件日志
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

	// 事件解析入队
	for _, vLog := range logs {
		item := parseEventToBatchItem(vLog)
		if item != nil {
			eventQueue = append(eventQueue, *item)
		}
	}

	// 触发条件2：达到批量数量上限
	if len(eventQueue) >= batchSize {
		logger.Logger.Info("批量入库触发：数量达到阈值", zap.Int("count", len(eventQueue)))
		flushBatch()
	}

	return nil
}

// ==================== 解析单条事件为批量结构体 ====================
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
	}
	return nil
}

// ==================== 批量入库（事务原子提交）====================
func flushBatch() {
	if len(eventQueue) == 0 {
		return
	}

	var events []model.ChainEvent
	var blockSyncs []model.BlockSync

	// 事务批量写入
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		events = make([]model.ChainEvent, 0, len(eventQueue))
		blockSyncs = make([]model.BlockSync, 0, len(blockMap))

		// 1. 组装批量事件
		for _, item := range eventQueue {
			events = append(events, item.ChainEvent)
		}

		// 2. 组装批量区块记录
		for _, b := range blockMap {
			blockSyncs = append(blockSyncs, b)
		}

		// 3. 批量插入事件
		if len(events) > 0 {
			if err := tx.CreateInBatches(events, 50).Error; err != nil {
				return err
			}
		}

		// 4. 批量插入区块同步记录
		if len(blockSyncs) > 0 {
			if err := tx.CreateInBatches(blockSyncs, 50).Error; err != nil {
				return err
			}
		}

		// 5. 批量更新用户质押数据
		batchUpdateUserStake(tx, eventQueue)

		return nil
	})

	if err != nil {
		logger.Logger.Error("入库失败", zap.Error(err))
		return
	}

	// 清空队列 + 重置时间
	eventQueue = make([]EventBatchItem, 0, batchSize)
	clear(blockMap)
	lastFlushTime = time.Now()

	logger.Logger.Info("入库成功", zap.Int("event_count", len(events)))
}

// ==================== 更新用户质押数据 ====================
func batchUpdateUserStake(tx *gorm.DB, queue []EventBatchItem) {
	userMap := make(map[string]*model.UserStake)

	// 归集用户数据
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
			// 新用户
			if item.IsStake {
				stake.UserAddress = item.UserAddr
				stake.StakeAmount = amount.String()
				stake.UpdateTime = uint64(time.Now().Unix())
				tx.Create(stake)
			}
		} else {
			// 更新用户
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

// ==================== 链重组处理====================
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

		// 清空队列 + 回滚数据库
		queueMutex.Lock()
		eventQueue = make([]EventBatchItem, 0, batchSize)
		clear(blockMap)
		queueMutex.Unlock()

		db.DB.Where("block_number >= ?", blockNumber).Delete(&model.ChainEvent{})
		db.DB.Where("block_number >= ?", blockNumber).Delete(&model.BlockSync{})
		config.GlobalConfig.Chain.StartBlock = blockNumber
	}
	return nil
}

// ==================== 优雅关机：强制刷新队列 ====================
func FlushOnShutdown() {
	queueMutex.Lock()
	defer queueMutex.Unlock()
	logger.Logger.Info("服务关闭，强制刷新批量队列")
	flushBatch()
}
