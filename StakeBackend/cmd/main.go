package main

import (
	"os"
	"os/signal"
	"syscall"

	"StakeBackend/internal/api"
	"StakeBackend/internal/config"
	"StakeBackend/internal/contract"
	"StakeBackend/internal/db"
	"StakeBackend/internal/indexer"
	"StakeBackend/internal/middleware"
	"StakeBackend/internal/model"
	"StakeBackend/internal/pkg/logger"
	"StakeBackend/internal/redis"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// ==================== 1. 初始化核心组件 ====================
	config.InitConfig()
	logger.InitLogger()

	// MySQL 初始化 + 自动建表
	db.InitMySQL()
	err := db.DB.AutoMigrate(
		&model.UserStake{},
		&model.ChainEvent{},
		&model.BlockSync{},
	)
	if err != nil {
		logger.Logger.Fatal("数据库迁移失败", zap.Error(err))
	}

	// Redis 初始化
	redis.InitRedis()

	// 区块链合约客户端初始化
	contract.InitContract()

	// ==================== 2. 启动批量事件索引器 ====================
	go indexer.StartIndexer()
	logger.Logger.Info("批量事件索引器已在后台运行")

	// ==================== 3. Gin 路由与中间件 ====================
	r := gin.Default()

	// 全局 API 限流中间件
	r.Use(middleware.RateLimitMiddleware())

	// 公共接口
	r.GET("/api/login", api.Login)

	// 私有接口（需要 JWT 鉴权）
	authGroup := r.Group("/api")
	authGroup.Use(middleware.JWTMiddleware())
	{
		authGroup.GET("/stake/info", api.GetStakeInfo)
		authGroup.GET("/stake/reward", api.GetPendingReward)
	}

	// ==================== 4. 关机 ====================
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quitChan
		logger.Logger.Info("开始优雅关闭服务...")

		// 关闭前强制刷新索引器队列，防止数据丢失
		indexer.FlushOnShutdown()

		logger.Logger.Info("所有批量数据已入库，服务安全关闭")
		os.Exit(0)
	}()

	// ==================== 5. 启动 HTTP 服务 ====================
	port := config.GlobalConfig.App.Port
	logger.Logger.Info("HTTP 服务启动成功", zap.String("port", port))
	if err := r.Run(":" + port); err != nil {
		logger.Logger.Fatal("服务启动失败", zap.Error(err))
	}
}
