package contract

import (
	"StakeBackend/internal/abi"
	"StakeBackend/internal/config"
	"StakeBackend/internal/pkg/logger"
	"StakeBackend/internal/redis"
	"context"
	"encoding/json"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	ethereumABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var (
	Client         *ethclient.Client
	MiningContract *abi.Abi
	Ctx            = context.Background()
	ContractABI    *ethereumABI.ABI
)

func InitContract() {
	err := Retry("RPC连接", func() error {
		client, err := ethclient.Dial(config.GlobalConfig.Chain.RPC)
		if err != nil {
			return err
		}
		Client = client
		return nil
	})
	if err != nil {
		logger.Logger.Fatal("RPC连接失败", zap.Error(err))
	}

	loadContractABI()

	miningAddr := common.HexToAddress(config.GlobalConfig.Chain.MiningProxy)
	contract, err := abi.NewAbi(miningAddr, Client)
	if err != nil {
		logger.Logger.Fatal("合约绑定失败", zap.Error(err))
	}
	MiningContract = contract

	logger.Logger.Info("合约连接成功")
}

func loadContractABI() {
	abiFile, err := os.ReadFile("./internal/abi/StakeMining.abi")
	if err != nil {
		logger.Logger.Fatal("读取ABI文件失败", zap.Error(err))
	}

	parsedABI, err := ethereumABI.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		logger.Logger.Fatal("解析ABI失败", zap.Error(err))
	}
	ContractABI = &parsedABI
}

func GetUserStakeAmount(userAddr string) (*big.Int, error) {
	var amount *big.Int
	err := Retry("查询质押金额", func() error {
		info, e := MiningContract.UserStakes(nil, common.HexToAddress(userAddr))
		amount = info.Amount
		return e
	})
	return amount, err
}

func GetPendingReward(userAddr string) (*big.Int, error) {
	var reward *big.Int
	err := Retry("查询待领取收益", func() error {
		var e error
		reward, e = MiningContract.GetPendingReward(nil, common.HexToAddress(userAddr))
		return e
	})
	return reward, err
}

func StakeTxData(amount *big.Int) ([]byte, error) {
	data, err := ContractABI.Pack("stake", amount)
	if err != nil {
		logger.Logger.Error("构造质押交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func UnstakeTxData(amount *big.Int) ([]byte, error) {
	data, err := ContractABI.Pack("unstake", amount)
	if err != nil {
		logger.Logger.Error("构造解质押交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func ClaimRewardTxData() ([]byte, error) {
	data, err := ContractABI.Pack("claimReward")
	if err != nil {
		logger.Logger.Error("构造领取收益交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// AddBlacklistTxData 构造添加黑名单交易数据
func AddBlacklistTxData(account common.Address) ([]byte, error) {
	data, err := ContractABI.Pack("addBlacklist", account)
	if err != nil {
		logger.Logger.Error("构造添加黑名单交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// RemoveBlacklistTxData 构造移除黑名单交易数据
func RemoveBlacklistTxData(account common.Address) ([]byte, error) {
	data, err := ContractABI.Pack("removeBlacklist", account)
	if err != nil {
		logger.Logger.Error("构造移除黑名单交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// PauseTxData 构造暂停合约交易数据
func PauseTxData() ([]byte, error) {
	data, err := ContractABI.Pack("pause")
	if err != nil {
		logger.Logger.Error("构造暂停交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// UnpauseTxData 构造恢复合约交易数据
func UnpauseTxData() ([]byte, error) {
	data, err := ContractABI.Pack("unpause")
	if err != nil {
		logger.Logger.Error("构造恢复交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// UpdateStakeLimitsTxData 构造更新质押限额交易数据
func UpdateStakeLimitsTxData(min, max *big.Int) ([]byte, error) {
	data, err := ContractABI.Pack("updateStakeLimits", min, max)
	if err != nil {
		logger.Logger.Error("构造更新限额交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

// SetRewardRateTxData 构造设置收益费率交易数据
func SetRewardRateTxData(rate *big.Int) ([]byte, error) {
	data, err := ContractABI.Pack("setRewardRate", rate)
	if err != nil {
		logger.Logger.Error("构造设置费率交易数据失败", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func EstimateGas(to string, data []byte) (uint64, error) {
	toAddr := common.HexToAddress(to)
	msg := ethereum.CallMsg{
		To:   &toAddr,
		Data: data,
	}
	gas, err := Client.EstimateGas(Ctx, msg)
	if err != nil {
		logger.Logger.Warn("Gas估算失败，使用默认值", zap.Error(err))
		return 300000, nil
	}
	return gas, nil
}

func GetStakeTokenAddress() string {
	return config.GlobalConfig.Chain.StakeToken
}

func GetMiningContractAddress() string {
	return config.GlobalConfig.Chain.MiningProxy
}

// IsPaused 查询合约是否暂停
func IsPaused() (bool, error) {
	var paused bool
	err := Retry("查询合约暂停状态", func() error {
		var e error
		paused, e = MiningContract.Paused(nil)
		return e
	})
	return paused, err
}

// GetStakeMinAmount 查询最小质押金额
func GetStakeMinAmount() (*big.Int, error) {
	var amount *big.Int
	err := Retry("查询最小质押金额", func() error {
		var e error
		amount, e = MiningContract.StakeMinAmount(nil)
		return e
	})
	return amount, err
}

// GetStakeMaxAmount 查询最大质押金额
func GetStakeMaxAmount() (*big.Int, error) {
	var amount *big.Int
	err := Retry("查询最大质押金额", func() error {
		var e error
		amount, e = MiningContract.StakeMaxAmount(nil)
		return e
	})
	return amount, err
}

// GetRewardRate 查询收益费率
func GetRewardRate() (*big.Int, error) {
	var rate *big.Int
	err := Retry("查询收益费率", func() error {
		var e error
		rate, e = MiningContract.RewardRate(nil)
		return e
	})
	return rate, err
}

// GetStakeToken 查询质押代币地址
func GetStakeToken() (string, error) {
	var addr string
	err := Retry("查询质押代币地址", func() error {
		token, e := MiningContract.StakeToken(nil)
		addr = token.Hex()
		return e
	})
	return addr, err
}

// GetRewardToken 查询收益代币地址
func GetRewardToken() (string, error) {
	var addr string
	err := Retry("查询收益代币地址", func() error {
		token, e := MiningContract.RewardToken(nil)
		addr = token.Hex()
		return e
	})
	return addr, err
}

// ContractStatus 合约状态结构
type ContractStatus struct {
	Paused      bool   `json:"paused"`
	MinAmount   string `json:"min_amount"`
	MaxAmount   string `json:"max_amount"`
	RewardRate  string `json:"reward_rate"`
	StakeToken  string `json:"stake_token"`
	RewardToken string `json:"reward_token"`
}

// GetContractStatusWithCache 带缓存的合约状态查询
func GetContractStatusWithCache() (*ContractStatus, error) {
	cacheKey := "contract:status"
	
	// 尝试从缓存获取
	if cached, err := redis.GetCache(cacheKey); err == nil && cached != "" {
		var status ContractStatus
		if err := json.Unmarshal([]byte(cached), &status); err == nil {
			logger.Logger.Debug("合约状态缓存命中")
			return &status, nil
		}
	}
	
	// 缓存未命中，从合约查询
	paused, err := IsPaused()
	if err != nil {
		return nil, err
	}
	
	minAmount, err := GetStakeMinAmount()
	if err != nil {
		return nil, err
	}
	
	maxAmount, err := GetStakeMaxAmount()
	if err != nil {
		return nil, err
	}
	
	rewardRate, err := GetRewardRate()
	if err != nil {
		return nil, err
	}
	
	stakeToken, err := GetStakeToken()
	if err != nil {
		return nil, err
	}
	
	rewardToken, err := GetRewardToken()
	if err != nil {
		return nil, err
	}
	
	status := &ContractStatus{
		Paused:      paused,
		MinAmount:   minAmount.String(),
		MaxAmount:   maxAmount.String(),
		RewardRate:  rewardRate.String(),
		StakeToken:  stakeToken,
		RewardToken: rewardToken,
	}
	
	// 写入缓存（5分钟）
	if jsonData, err := json.Marshal(status); err == nil {
		redis.SetCache(cacheKey, jsonData)
	}
	
	return status, nil
}

// InvalidateContractStatusCache 清除合约状态缓存
func InvalidateContractStatusCache() {
	redis.DelCache("contract:status")
	logger.Logger.Info("合约状态缓存已清除")
}

// IsBlacklistedWithCache 带缓存的黑名单检查
func IsBlacklistedWithCache(address string) (bool, error) {
	cacheKey := "blacklist:" + address
	
	// 尝试从缓存获取
	if cached, err := redis.GetCache(cacheKey); err == nil {
		if cached == "1" {
			return true, nil
		} else if cached == "0" {
			return false, nil
		}
	}
	
	// 缓存未命中，从合约查询
	blacklisted, err := MiningContract.Blacklist(nil, common.HexToAddress(address))
	if err != nil {
		return false, err
	}
	
	// 写入缓存（10分钟）
	cacheValue := "0"
	if blacklisted {
		cacheValue = "1"
	}
	redis.RDB.Set(redis.Ctx, cacheKey, cacheValue, 10*time.Minute)
	
	return blacklisted, nil
}
