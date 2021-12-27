package model

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB 数据库链接单例
var DB *gorm.DB

// InitDatabase 初始化mysql链接
func InitDatabase(connect string) {
	// Logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		// DSN: "root:GXB@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local",
		DSN:               connect,
		DefaultStringSize: 171, //utf8mb4
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		DisableForeignKeyConstraintWhenMigrating: true,      // 使用逻辑外键，即代码里自动处理外键关系，可提升速度
		Logger:                                   newLogger, // logger
	})

	if err != nil {
		panic(err)
	}

	DB = db

	//设置连接池
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(20)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour) // time.Second * 30

	migration()
}

// migration 执行数据迁移
func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4")

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserInfo{})

	DB.AutoMigrate(&Address{})
	DB.AutoMigrate(&Order{})

	DB.AutoMigrate(&AdminInfo{})

	DB.AutoMigrate(&Truck{})
}
