package contract

import (
	"StakeBackend/internal/abi"
	"StakeBackend/internal/config"
	"StakeBackend/internal/pkg/logger"
	"context"
	"math/big"
	"os"
	"strings"

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
