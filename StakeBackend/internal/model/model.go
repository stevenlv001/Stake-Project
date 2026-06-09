package model

import (
	"StakeBackend/internal/db"
)

type UserStake struct {
	db.BaseModel
	UserAddress         string `gorm:"size:42;index;unique" json:"user_address"`
	StakeAmount         string `gorm:"type:varchar(255)" json:"stake_amount"`
	RewardDebt          string `gorm:"type:varchar(255)" json:"reward_debt"`
	UpdateTime          uint64 `json:"update_time"`
	LastRewardUpdatedAt uint64 `json:"last_reward_updated_at"` // 上次收益更新时间戳，用于离线计算待领取收益
	TotalStaked         string `gorm:"type:varchar(255);default:'0'" json:"total_staked"`         // 累计质押金额
	TotalRewardsClaimed string `gorm:"type:varchar(255);default:'0'" json:"total_rewards_claimed"` // 累计领取收益
}

type ChainEvent struct {
	db.BaseModel
	BlockNumber uint64 `gorm:"index" json:"block_number"`
	BlockHash   string `gorm:"size:66;index" json:"block_hash"`
	TxHash      string `gorm:"size:66;uniqueIndex" json:"tx_hash"`
	EventType   string `gorm:"size:30" json:"event_type"`
	User        string `gorm:"size:42" json:"user"`
	Amount      string `gorm:"type:varchar(255)" json:"amount"`
	ExtraData   string `gorm:"type:text" json:"extra_data"` 
	EventTime   uint64 `json:"event_time"`
}

type BlockSync struct {
	db.BaseModel
	BlockNumber uint64 `gorm:"primaryKey" json:"block_number"`
	BlockHash   string `gorm:"size:66" json:"block_hash"`
	ParentHash  string `gorm:"size:66" json:"parent_hash"`
}

type AdminOperation struct {
	db.BaseModel
	BlockNumber uint64 `gorm:"index" json:"block_number"`
	BlockHash   string `gorm:"size:66" json:"block_hash"`
	TxHash      string `gorm:"size:66;uniqueIndex" json:"tx_hash"`
	EventType   string `gorm:"size:30" json:"event_type"`
	TargetUser  string `gorm:"size:42" json:"target_user"`
	MinLimit    string `gorm:"type:varchar(255)" json:"min_limit"`
	MaxLimit    string `gorm:"type:varchar(255)" json:"max_limit"`
	EventTime   uint64 `json:"event_time"`
}
