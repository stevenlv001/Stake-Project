package contract

import (
	"StakeBackend/internal/abi"
	"StakeBackend/internal/config"
	"StakeBackend/internal/pkg/logger"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var (
	Client         *ethclient.Client
	MiningContract *abi.Abi 
	Ctx            = context.Background()
)

// InitContract 初始化合约连接（带重试）
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

	// 绑定合约
	miningAddr := common.HexToAddress(config.GlobalConfig.Chain.MiningProxy)
	contract, err := abi.NewAbi(miningAddr, Client)
	if err != nil {
		logger.Logger.Fatal("合约绑定失败", zap.Error(err))
	}
	MiningContract = contract

	logger.Logger.Info("合约连接成功")
}

// GetUserStakeAmount 查询用户质押数量
func GetUserStakeAmount(userAddr string) (*big.Int, error) {
	var amount *big.Int
	err := Retry("查询质押金额", func() error {
		info, e := MiningContract.UserStakes(nil, common.HexToAddress(userAddr))
		amount = info.Amount
		return e
	})
	return amount, err
}

// GetPendingReward 查询待领取收益
func GetPendingReward(userAddr string) (*big.Int, error) {
	var reward *big.Int
	err := Retry("查询待领取收益", func() error {
		var e error
		reward, e = MiningContract.GetPendingReward(nil, common.HexToAddress(userAddr))
		return e
	})
	return reward, err
}
