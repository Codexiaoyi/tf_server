package repository

import (
	"fmt"
	"os"
	"tfserver/config"
	"tfserver/internal/model"
	"tfserver/pkg/log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

//初始化数据库上下文
func InitDbContext() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: false,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		//退出程序
		os.Exit(1)
	}

	// 迁移数据表
	_ = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}, &model.Account{}, &model.Team{}, &model.Member{}, &model.UserAlbum{}, &model.UserAlbumMedia{})

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

//开启事务
func BeginTransaction(db *gorm.DB, process func(tx *gorm.DB) error) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := process(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error

	log.ErrorLog("Transaction error:", err.Error())
	return err
}
