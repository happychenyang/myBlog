package clint

import (
	"fmt"
	config2 "go-blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

/**
go 的访问权限
变量名称，函数名称，常量名称 首字母大写是可以被其他包访问的
首字母大写为公有  首字母小写为私有
*/
var Db *gorm.DB

func init() {
	var configObj config2.MysqlConfig
	configInfo := configObj.GetMysqlConfig()
	//日志设置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //is 写入
		logger.Config{
			SlowThreshold:             time.Second,   //超时上线
			LogLevel:                  logger.Silent, //日志级别
			IgnoreRecordNotFoundError: true,          //忽略未找到的的记录
			Colorful:                  false,         //禁用颜色
		},
	)
	//链接数据库
	var dsnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", configInfo.Username,
		configInfo.Password,
		configInfo.Hostname,
		configInfo.Port,
		configInfo.Database,
		configInfo.Charset,
		configInfo.ParseTime,
		configInfo.Loc)
	var err error
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsnStr,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}

	sqlDB, err := Db.DB()
	//设置空闲连接池中链接的最大数量
	sqlDB.SetMaxIdleConns(configInfo.MaxIdleConn)
	//设置打开数据库的最大链接数
	sqlDB.SetMaxOpenConns(configInfo.MaxOpenConn)
	//设置链接池可服用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func Close() {
	sqlDB, err := Db.DB()
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	closeErr := sqlDB.Close()
	if closeErr != nil {
		return
	}
}
