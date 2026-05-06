package model

import (
	"StakeBackend/internal/db"
)

// UserStake 用户质押数据
type UserStake struct {
	db.BaseModel
	UserAddress string `gorm:"size:42;index;unique" json:"user_address"`
	StakeAmount string `gorm:"type:varchar(255)" json:"stake_amount"`
	RewardDebt  string `gorm:"type:varchar(255)" json:"reward_debt"`
	UpdateTime  uint64 `json:"update_time"`
}

// ChainEvent 链上事件
type ChainEvent struct {
	db.BaseModel
	BlockNumber uint64 `gorm:"index" json:"block_number"`
	BlockHash   string `gorm:"size:66;index" json:"block_hash"`
	TxHash      string `gorm:"size:66;uniqueIndex" json:"tx_hash"`
	EventType   string `gorm:"size:20" json:"event_type"`
	User        string `gorm:"size:42" json:"user"`
	Amount      string `gorm:"type:varchar(255)" json:"amount"`
	EventTime   uint64 `json:"event_time"`
}

// BlockSync 区块同步记录（处理链重组）
type BlockSync struct {
	db.BaseModel
	BlockNumber uint64 `gorm:"primaryKey" json:"block_number"`
	BlockHash   string `gorm:"size:66" json:"block_hash"`
	ParentHash  string `gorm:"size:66" json:"parent_hash"`
}