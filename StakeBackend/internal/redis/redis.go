package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"StakeBackend/internal/config"
	"StakeBackend/internal/pkg/logger"
	"time"

	"go.uber.org/zap"
)

var (
	RDB        *redis.Client
	Ctx        = context.Background()
	nullExpire = 60 * time.Second
	lockExpire = 10 * time.Second
	cacheExpire = 5 * time.Minute
)

// InitRedis 初始化Redis
func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Addr,
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		logger.Logger.Fatal("Redis连接失败", zap.Error(err))
	}

	RDB = rdb
	logger.Logger.Info("Redis连接成功")
}

// SetCache 写入缓存
func SetCache(key string, value interface{}) error {
	return RDB.Set(Ctx, key, value, cacheExpire).Err()
}

// SetNullCache 缓存空值（防穿透）
func SetNullCache(key string) error {
	return RDB.Set(Ctx, key, "null", nullExpire).Err()
}

// TryLock 互斥锁（防击穿）
func TryLock(key string) bool {
	ok, _ := RDB.SetNX(Ctx, key+":lock", 1, lockExpire).Result()
	return ok
}

// Unlock 释放锁
func Unlock(key string) {
	RDB.Del(Ctx, key+":lock")
}

// GetCache 获取缓存
func GetCache(key string) (string, error) {
	return RDB.Get(Ctx, key).Result()
}

// DelCache 删除缓存
func DelCache(key string) error {
	return RDB.Del(Ctx, key).Err()
}