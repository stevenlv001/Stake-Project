package config

import (
	"github.com/spf13/viper"
	"log"
)

// 全局配置对象
var GlobalConfig Config

// Config 总配置结构体
type Config struct {
	App       AppConfig       `mapstructure:"app"`
	MySQL     MySQLConfig     `mapstructure:"mysql"`
	Redis     RedisConfig     `mapstructure:"redis"`
	JWT       JWTConfig       `mapstructure:"jwt"`
	Chain     ChainConfig     `mapstructure:"chain"`
	Indexer   IndexerConfig   `mapstructure:"indexer"`
	RateLimit RateLimitConfig `mapstructure:"rate_limit"` 
}

// AppConfig 应用基础配置
type AppConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// MySQLConfig MySQL配置
type MySQLConfig struct {
	DSN     string `mapstructure:"dsn"`
	MaxIdle int    `mapstructure:"max_idle"`
	MaxOpen int    `mapstructure:"max_open"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JWTConfig JWT认证配置
type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

// ChainConfig 区块链RPC/合约配置
type ChainConfig struct {
	RPC         string `mapstructure:"rpc"`
	StakeToken  string `mapstructure:"stake_token"`
	RewardToken string `mapstructure:"reward_token"`
	MiningProxy string `mapstructure:"mining_proxy"`
	StartBlock  uint64 `mapstructure:"start_block"`
}

type IndexerConfig struct {
	BatchSize      int `mapstructure:"batch_size"`
	BatchWaitSeconds int `mapstructure:"batch_wait_seconds"`
}

type RateLimitConfig struct {
	MaxRequests  int `mapstructure:"max_requests"` 
	WindowSeconds int `mapstructure:"window_seconds"` 
}

// InitConfig 初始化配置文件
func InitConfig() {
	// 配置文件名称
	viper.SetConfigName("config")
	// 配置文件类型
	viper.SetConfigType("yaml")
	// 配置文件路径
	viper.AddConfigPath("./configs")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 解析配置到全局对象
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	log.Println("配置文件初始化成功")
}