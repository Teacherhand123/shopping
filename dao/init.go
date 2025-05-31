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
	// 只需要用这一行代码创建一个 gorm.DB 实例（db），然后通过 dbresolver 插件配置主库和从库的连接即可。
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
	sqlDB.SetMaxIdleConns(20)  // 设置连接池数 设置为20，意味着即使没有请求，最多保留20个连接在池子里，方便下次请求时能快速复用，减少频繁创建和销毁连接的开销。
	sqlDB.SetMaxOpenConns(100) // 打开连接数 设置为100，能保证高并发时最多只有100个连接，不会因为连接数太多把数据库“撑爆”。
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	// 主从配置
	// dbresolver 插件会自动帮你分配主库或从库
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connRead)},                       // 写操作 应该使用connWrite，但因为实际没做主从分离，所以这里用connRead
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},
		}))

	// db := dao.NewDBClient(ctx)
	// db.Find(&users)      // 自动走从库
	// db.Create(&user)     // 自动走主库

	// 开始构造表 首次启动需要
	// Migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
