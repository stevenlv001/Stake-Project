package api

import (
	"math/big"
	"time"

	apperr "StakeBackend/internal/pkg/errors"

	"StakeBackend/internal/contract"
	"StakeBackend/internal/db"
	"StakeBackend/internal/middleware"
	"StakeBackend/internal/model"
	"StakeBackend/internal/pkg/logger"
	"StakeBackend/internal/pkg/response"
	"StakeBackend/internal/redis"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AdminBlacklistRequest 黑名单管理请求
type AdminBlacklistRequest struct {
	Address string `json:"address" binding:"required"`
}

// AdminStakeLimitsRequest 质押限额请求
type AdminStakeLimitsRequest struct {
	MinAmount string `json:"min_amount" binding:"required"`
	MaxAmount string `json:"max_amount" binding:"required"`
}

// AdminRewardRateRequest 收益费率请求
type AdminRewardRateRequest struct {
	Rate string `json:"rate" binding:"required"`
}

// AddBlacklist godoc
// @Summary 添加黑名单
// @Description 将指定地址加入黑名单，禁止其质押操作
// @Tags 管理员
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body AdminBlacklistRequest true "黑名单地址"
// @Success 200 {object} response.Response
// @Router /api/admin/blacklist/add [post]
func AddBlacklist(c *gin.Context) {
	var req AdminBlacklistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 验证地址格式
	if !common.IsHexAddress(req.Address) {
		response.BadRequest(c, "无效的以太坊地址")
		return
	}

	txData, err := contract.AddBlacklistTxData(common.HexToAddress(req.Address))
	if err != nil {
		logger.Logger.Error("构造添加黑名单交易数据失败", zap.Error(err))
		response.ServerError(c, "构造交易数据失败")
		return
	}

	response.Success(c, buildTxResponse(txData))
	
	// 清除相关缓存
	contract.InvalidateContractStatusCache()
	redis.DelCache("blacklist:" + req.Address)
}

// RemoveBlacklist godoc
// @Summary 移除黑名单
// @Description 将指定地址从黑名单中移除
// @Tags 管理员
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body AdminBlacklistRequest true "黑名单地址"
// @Success 200 {object} response.Response
// @Router /api/admin/blacklist/remove [post]
func RemoveBlacklist(c *gin.Context) {
	var req AdminBlacklistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 验证地址格式
	if !common.IsHexAddress(req.Address) {
		response.BadRequest(c, "无效的以太坊地址")
		return
	}

	txData, err := contract.RemoveBlacklistTxData(common.HexToAddress(req.Address))
	if err != nil {
		logger.Logger.Error("构造移除黑名单交易数据失败", zap.Error(err))
		response.ServerError(c, "构造交易数据失败")
		return
	}

	response.Success(c, buildTxResponse(txData))
	
	// 清除相关缓存
	contract.InvalidateContractStatusCache()
	redis.DelCache("blacklist:" + req.Address)
}

// PauseContract godoc
// @Summary 暂停合约
// @Description 紧急暂停合约，停止所有质押操作
// @Tags 管理员
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /api/admin/pause [post]
func PauseContract(c *gin.Context) {
	txData, err := contract.PauseTxData()
	if err != nil {
		logger.Logger.Error("构造暂停交易数据失败", zap.Error(err))
		response.ServerError(c, "构造交易数据失败")
		return
	}

	response.Success(c, buildTxResponse(txData))
	
	// 清除合约状态缓存
	contract.InvalidateContractStatusCache()
}

// UnpauseContract godoc
// @Summary 恢复合约
// @Description 恢复已暂停的合约
// @Tags 管理员
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /api/admin/unpause [post]
func UnpauseContract(c *gin.Context) {
	txData, err := contract.UnpauseTxData()
	if err != nil {
		logger.Logger.Error("构造恢复交易数据失败", zap.Error(err))
		response.ServerError(c, "构造交易数据失败")
		return
	}

	response.Success(c, buildTxResponse(txData))
	
	// 清除合约状态缓存
	contract.InvalidateContractStatusCache()
}

// UpdateStakeLimits godoc
// @Summary 更新质押限额
// @Description 调整最小和最大质押金额限制
// @Tags 管理员
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body AdminStakeLimitsRequest true "质押限额"
// @Success 200 {object} response.Response
// @Router /api/admin/limits/update [post]
func UpdateStakeLimits(c *gin.Context) {
	var req AdminStakeLimitsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	minAmount, ok := new(big.Int).SetString(req.MinAmount, 10)
	if !ok {
		response.BadRequest(c, "最小金额格式错误")
		return
	}

	maxAmount, ok := new(big.Int).SetString(req.MaxAmount, 10)
	if !ok {
		response.BadRequest(c, "最大金额格式错误")
		return
	}

	// 验证限额合理性
	if minAmount.Cmp(maxAmount) >= 0 {
		response.BadRequest(c, "最小金额必须小于最大金额")
		return
	}

	txData, err := contract.UpdateStakeLimitsTxData(minAmount, maxAmount)
	if err != nil {
		logger.Logger.Error("构造更新限额交易数据失败", zap.Error(err))
		response.ServerError(c, "构造交易数据失败")
		return
	}

	response.Success(c, buildTxResponse(txData))
	
	// 清除合约状态缓存
	contract.InvalidateContractStatusCache()
}

// UpdateRewardRate godoc
// @Summary 更新收益费率
// @Description 调整每秒收益速率
// @Tags 管理员
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body AdminRewardRateRequest true "收益费率"
// @Success 200 {object} response.Response
// @Router /api/admin/rate/update [post]
func UpdateRewardRate(c *gin.Context) {
	var req AdminRewardRateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	rate, ok := new(big.Int).SetString(req.Rate, 10)
	if !ok {
		response.BadRequest(c, "费率格式错误")
		return
	}

	txData, err := contract.SetRewardRateTxData(rate)
	if err != nil {
		logger.Logger.Error("构造更新费率交易数据失败", zap.Error(err))
		response.ServerError(c, "构造交易数据失败")
		return
	}

	response.Success(c, buildTxResponse(txData))
	
	// 清除合约状态缓存
	contract.InvalidateContractStatusCache()
}

// GetContractStatus godoc
// @Summary 查询合约状态
// @Description 获取合约当前状态（是否暂停、质押限额等）
// @Tags 管理员
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /api/admin/status [get]
func GetContractStatus(c *gin.Context) {
	status, err := contract.GetContractStatusWithCache()
	if err != nil {
		logger.Logger.Error("查询合约状态失败", zap.Error(err))
		response.ServerError(c, "查询失败")
		return
	}

	response.Success(c, status)
}
