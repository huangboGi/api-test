package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// WthUserOrderInterest 用户订单利息模型
type WthUserOrderInterest struct {
	ID        uint            `gorm:"column:id;primarykey"`
	OrderID   uint            `gorm:"column:order_id;type:bigint;index:idx_order_id;comment:订单ID"`
	UserID    uint            `gorm:"column:user_id;type:bigint;index:idx_user_id;comment:用户ID"`
	Interest  decimal.Decimal `gorm:"column:interest;type:decimal(20,4);comment:利息"`
	Status    int8            `gorm:"column:status;type:tinyint;default:0;comment:状态"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime"`
}

// TableName 指定表名
func (WthUserOrderInterest) TableName() string {
	return "wth_user_order_interest"
}
