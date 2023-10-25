package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	timeout := os.Getenv("DB_TIMEOUT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, name, timeout)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogMode()),
	})

	if err != nil {
		panic("数据库连接失败, error: " + err.Error())
	}

	sqlDB, _ := DB.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func gormLogMode() logger.LogLevel {
	var logModes = map[string]logger.LogLevel{
		"Silent": logger.Silent,
		"Error":  logger.Error,
		"Warn":   logger.Warn,
		"Info":   logger.Info,
	}

	gormLogMode := logModes[os.Getenv("GORM_LOG_MODE")]
	if gormLogMode == 0 {
		gormLogMode = logger.Info
	}

	return gormLogMode
}
