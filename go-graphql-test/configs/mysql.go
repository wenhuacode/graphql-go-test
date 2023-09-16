package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	dbFactory *gorm.DB
	once      sync.Once
)

func NewMySQLOptions() *MySQLConfig {
	return &MySQLConfig{
		Host:                  "127.0.0.1",
		Port:                  3306,
		Username:              "rooy",
		Password:              "Mybabysjk888888.,",
		Database:              "",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifetime: time.Duration(10) * time.Second,
		LogLevel:              1, // Silent
	}
}

// GetPostgresConfig returns PostgresConfig object
func GetDB() (*gorm.DB, error) {
	vfg := NewMySQLOptions()
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			vfg.Username,
			vfg.Password,
			vfg.Host,
			vfg.Port,
			vfg.Database)
		//希望大家自己可以去封装logger
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second,                   // 慢 SQL 阈值
				LogLevel:                  logger.LogLevel(vfg.LogLevel), // 日志级别
				IgnoreRecordNotFoundError: true,                          // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,                         // 禁用彩色打印
			},
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			return
		}
		sqlDB, _ := db.DB()
		dbFactory = db

		sqlDB.SetMaxOpenConns(vfg.MaxOpenConnections)
		sqlDB.SetMaxIdleConns(vfg.MaxIdleConnections)
		sqlDB.SetConnMaxLifetime(vfg.MaxConnectionLifetime)
	})

	return dbFactory, nil
}
