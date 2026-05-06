package api

import (
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"StakeBackend/internal/contract"
	"StakeBackend/internal/db"
	"StakeBackend/internal/middleware"
	"StakeBackend/internal/model"
	"StakeBackend/internal/redis"
)

var baseRepo = &db.BaseRepo{}

// Login 用户登录（生成Token）
func Login(c *gin.Context) {
	userAddress := c.Query("address")
	if userAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "地址不能为空"})
		return
	}

	token, err := middleware.GenerateToken(userAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "token生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetStakeInfo 查询质押信息
func GetStakeInfo(c *gin.Context) {
	userAddr := c.GetString("user_address")
	cacheKey := "stake:info:" + userAddr

	// 查缓存
	cacheData, err := redis.GetCache(cacheKey)
	if err == nil {
		if cacheData == "null" {
			c.JSON(http.StatusOK, gin.H{"data": model.UserStake{}})
			return
		}
		var res model.UserStake
		_ = json.Unmarshal([]byte(cacheData), &res)
		c.JSON(http.StatusOK, gin.H{"data": res})
		return
	}

	// 加锁防击穿
	if !redis.TryLock(cacheKey) {
		c.JSON(http.StatusOK, gin.H{"msg": "系统繁忙"})
		c.Abort()
		return
	}
	defer redis.Unlock(cacheKey)

	// 查数据库
	var stake model.UserStake
	if err := db.DB.Where("user_address = ?", userAddr).First(&stake).Error; err != nil {
		redis.SetNullCache(cacheKey)
		c.JSON(http.StatusOK, gin.H{"data": model.UserStake{}})
		return
	}

	// 回写缓存
	jsonData, _ := json.Marshal(stake)
	redis.SetCache(cacheKey, jsonData)

	c.JSON(http.StatusOK, gin.H{"data": stake})
}

// GetPendingReward 查询待领取收益
func GetPendingReward(c *gin.Context) {
	userAddr := c.GetString("user_address")
	reward, err := contract.GetPendingReward(userAddr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{"pending_reward": reward.String()},
	})
}