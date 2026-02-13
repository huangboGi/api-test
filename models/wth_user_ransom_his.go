package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// WthUserRansomHis 用户赎回历史模型
type WthUserRansomHis struct {
	ID           uint            `gorm:"column:id;primarykey"`
	OrderID      uint            `gorm:"column:order_id;type:bigint;index:idx_order_id;comment:订单ID"`
	UserID       uint            `gorm:"column:user_id;type:bigint;index:idx_user_id;comment:用户ID"`
	RansomVolume decimal.Decimal `gorm:"column:ransom_volume;type:decimal(20,4);comment:赎回金额"`
	Interest     decimal.Decimal `gorm:"column:interest;type:decimal(20,4);comment:利息"`
	Status       int8            `gorm:"column:status;type:tinyint;default:0;comment:状态"`
	CreatedAt    time.Time       `gorm:"column:created_at;autoCreateTime"`
}

// TableName 指定表名
func (WthUserRansomHis) TableName() string {
	return "wth_user_ransom_his"
}
