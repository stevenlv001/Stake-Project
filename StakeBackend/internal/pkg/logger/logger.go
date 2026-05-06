package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"log"
)

var Logger *zap.Logger

func InitLogger() {
	// 日志切割
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writer, os.Stdout),
		zap.InfoLevel,
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	log.Println("日志系统初始化成功")
	
}