package txtracker

import (
	"StakeBackend/internal/contract"
	"StakeBackend/internal/pkg/logger"
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type TxStatus int

const (
	TxStatusPending   TxStatus = 0
	TxStatusConfirmed TxStatus = 1
	TxStatusFailed    TxStatus = 2
)

type TxRecord struct {
	TxHash      string
	Status      TxStatus
	BlockNumber uint64
	Error       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var (
	tracker     *TransactionTracker
	trackerOnce sync.Once
)

type TransactionTracker struct {
	txMap        map[string]*TxRecord
	mutex        sync.RWMutex
	ctx          context.Context
	cancel       context.CancelFunc
	pollInterval time.Duration
}

func GetTracker() *TransactionTracker {
	trackerOnce.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		tracker = &TransactionTracker{
			txMap:        make(map[string]*TxRecord),
			ctx:          ctx,
			cancel:       cancel,
			pollInterval: 3 * time.Second,
		}
	})
	return tracker
}

func (t *TransactionTracker) Start() {
	logger.Logger.Info("交易状态追踪器启动")
	go t.pollLoop()
	go t.cleanupLoop()
}

func (t *TransactionTracker) Stop() {
	logger.Logger.Info("交易状态追踪器关闭")
	t.cancel()
}

// cleanupLoop 定期清理已完成的交易记录（防止内存泄漏）
func (t *TransactionTracker) cleanupLoop() {
	cleanupTicker := time.NewTicker(1 * time.Hour)
	defer cleanupTicker.Stop()

	for {
		select {
		case <-t.ctx.Done():
			return
		case <-cleanupTicker.C:
			t.cleanupOldRecords()
		}
	}
}

// cleanupOldRecords 清理已完成超过24小时的交易记录
func (t *TransactionTracker) cleanupOldRecords() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	cutoff := time.Now().Add(-24 * time.Hour)
	cleanedCount := 0

	for hash, record := range t.txMap {
		// 只清理非pending状态且超过24小时的记录
		if record.Status != TxStatusPending && record.UpdatedAt.Before(cutoff) {
			delete(t.txMap, hash)
			cleanedCount++
		}
	}

	if cleanedCount > 0 {
		logger.Logger.Info("清理过期交易记录",
			zap.Int("cleaned_count", cleanedCount),
			zap.Int("remaining_count", len(t.txMap)),
		)
	}
}

func (t *TransactionTracker) TrackTx(txHash string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if _, exists := t.txMap[txHash]; exists {
		return
	}

	t.txMap[txHash] = &TxRecord{
		TxHash:    txHash,
		Status:    TxStatusPending,
		CreatedAt: time.Now(),
	}

	logger.Logger.Info("交易开始追踪", zap.String("tx_hash", txHash))
}

func (t *TransactionTracker) GetStatus(txHash string) (*TxRecord, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	record, exists := t.txMap[txHash]
	return record, exists
}

func (t *TransactionTracker) pollLoop() {
	ticker := time.NewTicker(t.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-t.ctx.Done():
			return
		case <-ticker.C:
			t.checkPendingTxs()
		}
	}
}

func (t *TransactionTracker) checkPendingTxs() {
	t.mutex.Lock()
	var pendingHashes []string
	for hash, record := range t.txMap {
		if record.Status == TxStatusPending {
			pendingHashes = append(pendingHashes, hash)
		}
	}
	t.mutex.Unlock()

	for _, txHash := range pendingHashes {
		t.checkTxStatus(txHash)
	}
}

func (t *TransactionTracker) checkTxStatus(txHash string) {
	ctx := context.Background()

	_, isPending, err := contract.Client.TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		logger.Logger.Warn("查询交易失败", zap.String("tx_hash", txHash), zap.Error(err))
		return
	}

	if isPending {
		logger.Logger.Debug("交易等待打包", zap.String("tx_hash", txHash))
		return
	}

	receipt, err := contract.Client.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		logger.Logger.Error("获取交易收据失败", zap.String("tx_hash", txHash), zap.Error(err))
		return
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	record, exists := t.txMap[txHash]
	if !exists {
		return
	}

	record.BlockNumber = receipt.BlockNumber.Uint64()
	record.UpdatedAt = time.Now()

	if receipt.Status == 1 {
		record.Status = TxStatusConfirmed
		logger.Logger.Info("交易确认成功",
			zap.String("tx_hash", txHash),
			zap.Uint64("block_number", record.BlockNumber),
		)
	} else {
		record.Status = TxStatusFailed
		logger.Logger.Error("交易执行失败",
			zap.String("tx_hash", txHash),
			zap.Uint64("block_number", record.BlockNumber),
		)
	}
}

func (t *TransactionTracker) WaitForConfirmation(txHash string, timeout time.Duration) (*TxRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(t.pollInterval)
	defer ticker.Stop()

	deadline := time.Now().Add(timeout)

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			record, exists := t.GetStatus(txHash)
			if !exists {
				continue
			}

			if time.Now().After(deadline) {
				return record, context.DeadlineExceeded
			}

			if record.Status != TxStatusPending {
				return record, nil
			}

			t.checkTxStatus(txHash)
		}
	}
}

func WatchTx(txHash string) {
	GetTracker().TrackTx(txHash)
}

func GetTxStatus(txHash string) (*TxRecord, bool) {
	return GetTracker().GetStatus(txHash)
}

func WaitTxConfirmed(txHash string, timeout time.Duration) (*TxRecord, error) {
	return GetTracker().WaitForConfirmation(txHash, timeout)
}

func GetConfirmedTxBlock(txHash string) (uint64, error) {
	record, exists := GetTxStatus(txHash)
	if !exists {
		return 0, nil
	}

	if record.Status != TxStatusConfirmed {
		return 0, nil
	}

	return record.BlockNumber, nil
}

func GetBlockTxs(blockNumber uint64) ([]string, error) {
	ctx := context.Background()

	block, err := contract.Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}

	txHashes := make([]string, 0, len(block.Transactions()))
	for _, tx := range block.Transactions() {
		txHashes = append(txHashes, tx.Hash().Hex())
	}

	return txHashes, nil
}
