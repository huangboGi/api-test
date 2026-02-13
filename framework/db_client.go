package framework

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"my_stonks_api_tests/config"
)

// DBClient 数据库测试客户端（只读）
type DBClient struct {
	*gorm.DB
}

var (
	dbInstance *DBClient
	once       sync.Once
)

// NewDBClient 创建数据库测试客户端（只读）
func NewDBClient() *DBClient {
	once.Do(func() {
		dsn := config.Cfg.GetDSN()
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			// 禁用写操作
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   false,
		})
		if err != nil {
			panic(fmt.Sprintf("连接数据库失败: %v", err))
		}

		// 如果配置为只读模式，添加拦截器
		if config.Cfg.DBReadOnly {
			// 拦截所有写操作，在运行时panic
			db.Callback().Create().Before("gorm:create").Register("gorm:ensure_readonly", func(d *gorm.DB) {
				panic("❌ 数据库处于只读模式！不允许执行 CREATE 操作")
			})

			db.Callback().Update().Before("gorm:update").Register("gorm:ensure_readonly", func(d *gorm.DB) {
				panic("❌ 数据库处于只读模式！不允许执行 UPDATE 操作")
			})

			db.Callback().Delete().Before("gorm:delete").Register("gorm:ensure_readonly", func(d *gorm.DB) {
				panic("❌ 数据库处于只读模式！不允许执行 DELETE 操作")
			})

			fmt.Printf("【数据库】连接成功: %s (只读模式 - 写操作将被拦截)\n", config.Cfg.DBName)
		} else {
			fmt.Printf("【数据库】连接成功: %s (读写模式)\n", config.Cfg.DBName)
		}

		dbInstance = &DBClient{
			DB: db,
		}
	})

	return dbInstance
}

// Query 根据条件查询（只读）
func (db *DBClient) Query(model interface{}, condition string, args ...interface{}) error {
	return db.Where(condition, args...).First(model).Error
}

// QueryByID 根据ID查询（只读）
func (db *DBClient) QueryByID(model interface{}, id uint) error {
	return db.First(model, id).Error
}

// QueryList 查询列表（只读）
func (db *DBClient) QueryList(model interface{}, condition string, args ...interface{}) error {
	return db.Where(condition, args...).Find(model).Error
}

// GetByCondition 根据条件查询单条记录（只读）
func (db *DBClient) GetByCondition(model interface{}, condition string, args ...interface{}) error {
	return db.Where(condition, args...).First(model).Error
}

// Exists 检查记录是否存在（只读）
func (db *DBClient) Exists(model interface{}, condition string, args ...interface{}) (bool, error) {
	var count int64
	err := db.Model(model).Where(condition, args...).Count(&count).Error
	return count > 0, err
}

// GetCount 获取记录数（只读）
func (db *DBClient) GetCount(model interface{}, condition string, args ...interface{}) (int64, error) {
	var count int64
	err := db.Model(model).Where(condition, args...).Count(&count).Error
	return count, err
}

// LogSQL 打印SQL（调试用）
func (db *DBClient) LogSQL(enable bool) {
	if enable {
		db.DB = db.DB.Debug()
	}
}

// EnsureReadOnly 确保数据库处于只读模式（如果尝试写操作会panic）
func (db *DBClient) EnsureReadOnly() {
	if !config.Cfg.DBReadOnly {
		return
	}

	// 拦截所有写操作
	db.Callback().Create().Before("gorm:create").Register("gorm:ensure_readonly", func(db *gorm.DB) {
		panic("❌ 数据库处于只读模式！不允许执行 CREATE 操作")
	})

	db.Callback().Update().Before("gorm:update").Register("gorm:ensure_readonly", func(db *gorm.DB) {
		panic("❌ 数据库处于只读模式！不允许执行 UPDATE 操作")
	})

	db.Callback().Delete().Before("gorm:delete").Register("gorm:ensure_readonly", func(db *gorm.DB) {
		panic("❌ 数据库处于只读模式！不允许执行 DELETE 操作")
	})
}
