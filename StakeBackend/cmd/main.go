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
	"StakeBackend/internal/pkg/txtracker"
	"StakeBackend/internal/redis"

	_ "StakeBackend/docs" // Swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Stake Mining API
// @version 1.0
// @description 质押挖矿系统API文档，包含用户接口和管理员接口
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	config.InitConfig()
	logger.InitLogger()

	err := db.DB.AutoMigrate(
		&model.UserStake{},
		&model.ChainEvent{},
		&model.BlockSync{},
		&model.AdminOperation{},
	)
	if err != nil {
		logger.Logger.Fatal("数据库迁移失败", zap.Error(err))
	}

	redis.InitRedis()
	contract.InitContract()

	txtracker.GetTracker().Start()
	logger.Logger.Info("交易状态追踪器已启动")

	go indexer.StartIndexer()
	logger.Logger.Info("批量事件索引器已在后台运行")

	r := gin.Default()
	r.Use(middleware.RateLimitMiddleware())

	// Swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/api/login", api.Login)
	r.GET("/api/admin/login", api.AdminLogin) // 管理员登录

	// 用户接口（需要JWT认证）
	authGroup := r.Group("/api")
	authGroup.Use(middleware.JWTMiddleware())
	{
		authGroup.GET("/stake/info", api.GetStakeInfo)
		authGroup.GET("/stake/reward", api.GetPendingReward)
		authGroup.POST("/stake/do", api.DoStake)
		authGroup.POST("/stake/unstake", api.DoUnstake)
		authGroup.POST("/stake/claim", api.DoClaimReward)
		authGroup.GET("/tx/status", api.TrackTxStatus)
		authGroup.GET("/tx/wait", api.WaitTxConfirm)
	}

	// 管理员接口（需要管理员JWT认证）
	adminGroup := r.Group("/api/admin")
	adminGroup.Use(middleware.AdminJWTMiddleware())
	{
		// 黑名单管理
		adminGroup.POST("/blacklist/add", api.AddBlacklist)
		adminGroup.POST("/blacklist/remove", api.RemoveBlacklist)

		// 合约控制
		adminGroup.POST("/pause", api.PauseContract)
		adminGroup.POST("/unpause", api.UnpauseContract)

		// 参数调整
		adminGroup.POST("/limits/update", api.UpdateStakeLimits)
		adminGroup.POST("/rate/update", api.UpdateRewardRate)

		// 状态查询
		adminGroup.GET("/status", api.GetContractStatus)
	}

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quitChan
		logger.Logger.Info("开始优雅关闭服务...")

		txtracker.GetTracker().Stop()
		indexer.FlushOnShutdown()

		logger.Logger.Info("所有批量数据已入库，服务安全关闭")
		os.Exit(0)
	}()

	port := config.GlobalConfig.App.Port
	logger.Logger.Info("HTTP 服务启动成功", zap.String("port", port))
	if err := r.Run(":" + port); err != nil {
		logger.Logger.Fatal("服务启动失败", zap.Error(err))
	}
}
