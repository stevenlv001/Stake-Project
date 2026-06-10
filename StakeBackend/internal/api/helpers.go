package api

import (
	"StakeBackend/internal/contract"
	"StakeBackend/internal/pkg/logger"
	"StakeBackend/internal/pkg/txtracker"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)


func buildTxResponse(txData []byte) map[string]interface{} {
	gasLimit, err := contract.EstimateGas(contract.GetMiningContractAddress(), txData)
	if err != nil {
		logger.Logger.Warn("Gas估算失败，使用默认值", zap.Error(err))
		gasLimit = 300000
	}

	return map[string]interface{}{
		"to":        contract.GetMiningContractAddress(),
		"data":      common.Bytes2Hex(txData),
		"value":     "0",
		"gas_limit": gasLimit,
	}
}


func formatTxStatus(record *txtracker.TxRecord) TxStatusResponse {
	statusStr := "pending"
	if record.Status == 1 {
		statusStr = "confirmed"
	} else if record.Status == 2 {
		statusStr = "failed"
	}

	return TxStatusResponse{
		TxHash:      record.TxHash,
		Status:      statusStr,
		BlockNumber: record.BlockNumber,
	}
}
