package api

import (
	"encoding/json"
	"math/big"
	"time"

	"StakeBackend/internal/contract"
	"StakeBackend/internal/db"
	"StakeBackend/internal/middleware"
	"StakeBackend/internal/model"
	"StakeBackend/internal/pkg/logger"
	"StakeBackend/internal/pkg/response"
	"StakeBackend/internal/pkg/txtracker"
	"StakeBackend/internal/redis"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var baseRepo = &db.BaseRepo{}

type StakeRequest struct {
	Amount string `json:"amount" binding:"required"`
}

type TxDataResponse struct {
	To       string `json:"to"`
	Data     string `json:"data"`
	Value    string `json:"value"`
	GasLimit uint64 `json:"gas_limit"`
}

type TxStatusResponse struct {
	TxHash        string `json:"tx_hash"`
	Status        string `json:"status"`
	Confirmations int    `json:"confirmations"`
	BlockNumber   uint64 `json:"block_number"`
}

// Login godoc
// @Summary 用户登录
// @Description 用户登录并生成JWT Token
// @Tags 认证
// @Param address query string true "用户钱包地址"
// @Success 200 {object} response.Response{data=map[string]string}
// @Router /api/login [get]
func Login(c *gin.Context) {
	userAddress := c.Query("address")
	if userAddress == "" {
		response.BadRequest(c, "地址不能为空")
		return
	}

	token, err := middleware.GenerateToken(userAddress)
	if err != nil {
		response.ServerError(c, "token生成失败")
		return
	}

	response.Success(c, gin.H{"token": token})
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录并生成JWT Token（演示用，生产环境应使用签名验证）
// @Tags 管理员
// @Param admin_id query string true "管理员ID"
// @Param role query string true "角色 (admin/super_admin)"
// @Success 200 {object} response.Response{data=map[string]string}
// @Router /api/admin/login [get]
func AdminLogin(c *gin.Context) {
	adminID := c.Query("admin_id")
	role := c.Query("role")

	if adminID == "" || role == "" {
		response.BadRequest(c, "参数不能为空")
		return
	}

	// 验证角色
	if role != "admin" && role != "super_admin" {
		response.BadRequest(c, "无效的角色")
		return
	}

	// TODO: 生产环境应该验证管理员身份（如签名验证、数据库查询等）
	token, err := middleware.GenerateAdminToken(adminID, role)
	if err != nil {
		response.ServerError(c, "token生成失败")
		return
	}

	response.Success(c, gin.H{"token": token})
}

// GetStakeInfo godoc
// @Summary 查询质押信息
// @Description 获取当前用户的质押信息
// @Tags 质押
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.UserStake}
// @Router /api/stake/info [get]
func GetStakeInfo(c *gin.Context) {
	userAddr := c.GetString("user_address")
	cacheKey := "stake:info:" + userAddr

	cacheData, err := redis.GetCache(cacheKey)
	if err == nil {
		if cacheData == "null" {
			response.Success(c, model.UserStake{})
			return
		}
		var res model.UserStake
		_ = json.Unmarshal([]byte(cacheData), &res)
		response.Success(c, res)
		return
	}

	if !redis.TryLock(cacheKey) {
		response.Fail(c, 503, "系统繁忙")
		c.Abort()
		return
	}
	defer redis.Unlock(cacheKey)

	var stake model.UserStake
	if err := db.DB.Where("user_address = ?", userAddr).First(&stake).Error; err != nil {
		redis.SetNullCache(cacheKey)
		response.Success(c, model.UserStake{})
		return
	}

	jsonData, _ := json.Marshal(stake)
	redis.SetCache(cacheKey, jsonData)

	response.Success(c, stake)
}

// GetPendingReward godoc
// @Summary 查询待领取收益
// @Description 获取当前用户的待领取挖矿收益
// @Tags 质押
// @Security BearerAuth
// @Success 200 {object} response.Response{data=map[string]string}
// @Router /api/stake/reward [get]
func GetPendingReward(c *gin.Context) {
	userAddr := c.GetString("user_address")
	reward, err := contract.GetPendingReward(userAddr)
	if err != nil {
		response.ServerError(c, "查询失败")
		return
	}

	response.Success(c, gin.H{"pending_reward": reward.String()})
}

// DoStake godoc
// @Summary 质押代币
// @Description 构造质押交易数据，供前端钱包签名
// @Tags 质押
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.StakeRequest true "质押请求"
// @Success 200 {object} response.Response{data=TxDataResponse}
// @Router /api/stake/do [post]
func DoStake(c *gin.Context) {
	var req StakeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		response.BadRequest(c, "金额格式错误")
		return
	}

	txData, err := contract.StakeTxData(amount)
	if err != nil {
		response.ServerError(c, "构造交易数据失败")
		return
	}

	gasLimit, err := contract.EstimateGas(contract.GetMiningContractAddress(), txData)
	if err != nil {
		logger.Logger.Warn("Gas估算失败，使用默认值", zap.Error(err))
		gasLimit = 300000 // 使用保守的默认值
	}

	response.Success(c, gin.H{
		"to":        contract.GetMiningContractAddress(),
		"data":      common.Bytes2Hex(txData),
		"value":     "0",
		"gas_limit": gasLimit,
	})
}

// DoUnstake godoc
// @Summary 解质押
// @Description 构造解质押交易数据，供前端钱包签名
// @Tags 质押
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.StakeRequest true "解质押请求"
// @Success 200 {object} response.Response{data=TxDataResponse}
// @Router /api/stake/unstake [post]
func DoUnstake(c *gin.Context) {
	var req StakeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		response.BadRequest(c, "金额格式错误")
		return
	}

	txData, err := contract.UnstakeTxData(amount)
	if err != nil {
		response.ServerError(c, "构造交易数据失败")
		return
	}

	gasLimit, err := contract.EstimateGas(contract.GetMiningContractAddress(), txData)
	if err != nil {
		logger.Logger.Warn("Gas估算失败，使用默认值", zap.Error(err))
		gasLimit = 300000 // 使用保守的默认值
	}

	response.Success(c, gin.H{
		"to":        contract.GetMiningContractAddress(),
		"data":      common.Bytes2Hex(txData),
		"value":     "0",
		"gas_limit": gasLimit,
	})
}

// DoClaimReward godoc
// @Summary 领取收益
// @Description 构造领取收益交易数据，供前端钱包签名
// @Tags 质押
// @Security BearerAuth
// @Success 200 {object} response.Response{data=TxDataResponse}
// @Router /api/stake/claim [post]
func DoClaimReward(c *gin.Context) {
	txData, err := contract.ClaimRewardTxData()
	if err != nil {
		response.ServerError(c, "构造交易数据失败")
		return
	}

	gasLimit, err := contract.EstimateGas(contract.GetMiningContractAddress(), txData)
	if err != nil {
		logger.Logger.Warn("Gas估算失败，使用默认值", zap.Error(err))
		gasLimit = 300000 // 使用保守的默认值
	}

	response.Success(c, gin.H{
		"to":        contract.GetMiningContractAddress(),
		"data":      common.Bytes2Hex(txData),
		"value":     "0",
		"gas_limit": gasLimit,
	})
}

// TrackTxStatus godoc
// @Summary 查询交易状态
// @Description 查询已提交交易的状态
// @Tags 交易追踪
// @Security BearerAuth
// @Param tx_hash query string true "交易哈希"
// @Success 200 {object} response.Response{data=TxStatusResponse}
// @Router /api/tx/status [get]
func TrackTxStatus(c *gin.Context) {
	txHash := c.Query("tx_hash")
	if txHash == "" {
		response.BadRequest(c, "交易哈希不能为空")
		return
	}

	record, exists := txtracker.GetTxStatus(txHash)
	if !exists {
		response.Fail(c, 404, "交易未被追踪")
		return
	}

	statusStr := "pending"
	if record.Status == 1 {
		statusStr = "confirmed"
	} else if record.Status == 2 {
		statusStr = "failed"
	}

	response.Success(c, TxStatusResponse{
		TxHash:      record.TxHash,
		Status:      statusStr,
		BlockNumber: record.BlockNumber,
	})
}

// WaitTxConfirm godoc
// @Summary 等待交易确认
// @Description 等待交易确认（同步接口，最多等待30秒）
// @Tags 交易追踪
// @Security BearerAuth
// @Param tx_hash query string true "交易哈希"
// @Success 200 {object} response.Response{data=TxStatusResponse}
// @Router /api/tx/wait [get]
func WaitTxConfirm(c *gin.Context) {
	txHash := c.Query("tx_hash")
	if txHash == "" {
		response.BadRequest(c, "交易哈希不能为空")
		return
	}

	record, err := txtracker.WaitTxConfirmed(txHash, 30*time.Second)
	if err != nil {
		response.ServerError(c, "等待确认超时: "+err.Error())
		return
	}

	statusStr := "pending"
	if record.Status == 1 {
		statusStr = "confirmed"
	} else if record.Status == 2 {
		statusStr = "failed"
	}

	response.Success(c, TxStatusResponse{
		TxHash:      record.TxHash,
		Status:      statusStr,
		BlockNumber: record.BlockNumber,
	})
}
