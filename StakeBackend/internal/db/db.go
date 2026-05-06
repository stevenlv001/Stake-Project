package db

import (
	"StakeBackend/internal/config"
	"StakeBackend/internal/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"go.uber.org/zap"
)

var DB *gorm.DB

// BaseModel 基础模型
type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// IBaseRepo 通用CRUD接口
type IBaseRepo interface {
	Create(data interface{}) error
	Update(data interface{}) error
	Delete(data interface{}) error
	FindByID(id uint, data interface{}) error
	FindList(condition interface{}, list interface{}, page, size int) error
}

type BaseRepo struct{}

// InitMySQL 初始化数据库
func InitMySQL() {
	dsn := config.GlobalConfig.MySQL.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		logger.Logger.Fatal("MySQL连接失败", zap.Error(err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.GlobalConfig.MySQL.MaxIdle)
	sqlDB.SetMaxOpenConns(config.GlobalConfig.MySQL.MaxOpen)
	DB = db

	logger.Logger.Info("MySQL连接成功")
}

// Create 通用创建
func (b *BaseRepo) Create(data interface{}) error {
	return DB.Create(data).Error
}

// Update 通用更新
func (b *BaseRepo) Update(data interface{}) error {
	return DB.Save(data).Error
}

// Delete 通用删除
func (b *BaseRepo) Delete(data interface{}) error {
	return DB.Delete(data).Error
}

// FindByID ID查询
func (b *BaseRepo) FindByID(id uint, data interface{}) error {
	return DB.First(data, id).Error
}

// FindList 分页查询
func (b *BaseRepo) FindList(condition interface{}, list interface{}, page, size int) error {
	return DB.Where(condition).Offset((page - 1) * size).Limit(size).Find(list).Error
}
