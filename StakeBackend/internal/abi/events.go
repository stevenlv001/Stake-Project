package abi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 事件签名哈希常量
var (
	// AbiStakedEventID Staked事件签名哈希
	AbiStakedEventID = crypto.Keccak256Hash([]byte("Staked(address,uint256,uint256)"))
	
	// AbiUnstakedEventID Unstaked事件签名哈希
	AbiUnstakedEventID = crypto.Keccak256Hash([]byte("Unstaked(address,uint256,uint256)"))
	
	// AbiRewardClaimedEventID RewardClaimed事件签名哈希
	AbiRewardClaimedEventID = crypto.Keccak256Hash([]byte("RewardClaimed(address,uint256,uint256)"))

	// 事件名称常量
	EventStaked        = "Staked"
	EventUnstaked      = "Unstaked"
	EventRewardClaimed = "RewardClaimed"
)

// ParseEventByTopic 根据日志主题解析事件类型
func ParseEventByTopic(topic common.Hash) string {
	switch topic.Hex() {
	case AbiStakedEventID.Hex():
		return EventStaked
	case AbiUnstakedEventID.Hex():
		return EventUnstaked
	case AbiRewardClaimedEventID.Hex():
		return EventRewardClaimed
	default:
		return "Unknown"
	}
}