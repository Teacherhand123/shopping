package dao

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	// 设置日志
	var ormLogger logger.Interface
	// if gin.Mode() == "debug" {
	// 	ormLogger = logger.Default.LogMode(logger.Info) // 日志级别为 Info
	// } else {
	// 	ormLogger = logger.Default
	// }

	ormLogger = logger.Default.LogMode(logger.Info)

	// 开启数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      connRead,
		DefaultStringSize:        256,  // string类型字段默认长度
		DisableDatetimePrecision: true, // 禁止datetime精度，mysql5.6之前的数据库不支持
		DontSupportRenameIndex:   true, // 重命名索引，就要把索引删了再重建，mysql 5.7不支持
		DontSupportRenameColumn:  true, // 用change重命名列，mysql8之前数据库不支持
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("DB连接出错: ", err)
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池数
	sqlDB.SetMaxOpenConns(100) // 打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	// 主从配置
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connRead)},                       // 写操作
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},
		}))

	// 开始构造表 首次启动需要
	Migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
