package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type dbHelper struct {
}

var (
	DB       = initDB()
	DBHelper = &dbHelper{}
)

// initDB 初始化DB
func initDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  Config.DBDSN,
		PreferSimpleProtocol: true, // 禁用隐式准备语句用法
	}), &gorm.Config{
		SkipDefaultTransaction:                   true, // 禁用默认事务
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
	})
	if err != nil {
		panic("InitDB Open err: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("InitDB db.DB err: " + err.Error())
	}
	// 数据库空闲连接池最大值
	sqlDB.SetMaxIdleConns(Config.DBMaxIdleConns)
	// 数据库连接池最大值
	sqlDB.SetMaxOpenConns(Config.DBMaxOpenConns)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(Config.DBConnMaxLifetimeHours) * time.Hour)
	return db
}

// Page 返回分页查询条件limit和offset,
func (helper *dbHelper) Page(pageNo int, pageSize int) (int, int) {
	limit := pageSize
	offset := pageSize * (pageNo - 1)
	return limit, offset
}
