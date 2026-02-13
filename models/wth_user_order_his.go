package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// WthUserOrderHis 用户订单历史模型
type WthUserOrderHis struct {
	ID           uint            `gorm:"column:id;primarykey"`
	OrderID      uint            `gorm:"column:order_id;type:bigint;index:idx_order_id;comment:订单ID"`
	UserID       uint            `gorm:"column:user_id;type:bigint;index:idx_user_id;comment:用户ID"`
	ProductID    uint            `gorm:"column:product_id;type:bigint;comment:产品ID"`
	InvestAmount decimal.Decimal `gorm:"column:invest_amount;type:decimal(20,4);comment:投资金额"`
	Status       int8            `gorm:"column:status;type:tinyint;default:0;comment:状态"`
	CreatedAt    time.Time       `gorm:"column:created_at;autoCreateTime"`
}

// TableName 指定表名
func (WthUserOrderHis) TableName() string {
	return "wth_user_order_his"
}
