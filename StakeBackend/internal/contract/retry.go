package contract

import (
	"StakeBackend/internal/pkg/logger"
	"time"

	"go.uber.org/zap"
)

const (
	maxRetries = 3
	baseDelay  = 1 * time.Second
)

// Retry 合约调用重试（3次+指数退避）
func Retry(operation string, fn func() error) error {
	var err error
	delay := baseDelay

	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		if i == maxRetries-1 {
			break
		}

		logger.Logger.Warn("合约调用重试",
			zap.String("操作", operation),
			zap.Int("次数", i+1),
			zap.Duration("延迟", delay),
			zap.Error(err),
		)
		time.Sleep(delay)
		delay *= 2
	}

	logger.Logger.Error("合约调用失败", zap.String("操作", operation), zap.Error(err))
	return err
}