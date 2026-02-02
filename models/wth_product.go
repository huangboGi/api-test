package models

import (
	"time"
)

// WthProduct 理财产品模型
type WthProduct struct {
	ID          uint      `gorm:"column:id;primarykey"`
	ProductKey  string    `gorm:"column:product_key;type:varchar(50);uniqueIndex:uk_product_key;comment:产品标识"`
	ProductName string    `gorm:"column:product_name;type:varchar(200);comment:产品名称"`
	CoinKey     string    `gorm:"column:coin_key;type:varchar(50);comment:币种标识"`
	Status      int8      `gorm:"column:status;type:tinyint;default:1;comment:状态 0-下架 1-上架"`
	Sort        int       `gorm:"column:sort;type:int;default:0;comment:排序"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

// TableName 指定表名
func (WthProduct) TableName() string {
	return "wth_product"
}
