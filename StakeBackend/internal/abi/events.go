package abi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AbiStakedEventID         = crypto.Keccak256Hash([]byte("Staked(address,uint256,uint256)"))
	AbiUnstakedEventID       = crypto.Keccak256Hash([]byte("Unstaked(address,uint256,uint256)"))
	AbiRewardClaimedEventID   = crypto.Keccak256Hash([]byte("RewardClaimed(address,uint256,uint256)"))
	AbiBlacklistAddedEventID  = crypto.Keccak256Hash([]byte("BlacklistAdded(address)"))
	AbiBlacklistRemovedEventID = crypto.Keccak256Hash([]byte("BlacklistRemoved(address)"))
	AbiStakeLimitsUpdatedEventID = crypto.Keccak256Hash([]byte("StakeLimitsUpdated(uint256,uint256)"))

	EventStaked           = "Staked"
	EventUnstaked         = "Unstaked"
	EventRewardClaimed    = "RewardClaimed"
	EventBlacklistAdded   = "BlacklistAdded"
	EventBlacklistRemoved = "BlacklistRemoved"
	EventStakeLimitsUpdated = "StakeLimitsUpdated"
)

func ParseEventByTopic(topic common.Hash) string {
	switch topic.Hex() {
	case AbiStakedEventID.Hex():
		return EventStaked
	case AbiUnstakedEventID.Hex():
		return EventUnstaked
	case AbiRewardClaimedEventID.Hex():
		return EventRewardClaimed
	case AbiBlacklistAddedEventID.Hex():
		return EventBlacklistAdded
	case AbiBlacklistRemovedEventID.Hex():
		return EventBlacklistRemoved
	case AbiStakeLimitsUpdatedEventID.Hex():
		return EventStakeLimitsUpdated
	default:
		return "Unknown"
	}
}