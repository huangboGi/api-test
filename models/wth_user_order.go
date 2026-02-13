package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// WthUserOrder 用户订单模型
type WthUserOrder struct {
	ID                uint            `gorm:"column:id;primarykey"`
	OrderNo           string          `gorm:"column:order_no;type:varchar(50);uniqueIndex:uk_order_no;comment:订单编号"`
	UserID            uint            `gorm:"column:user_id;type:bigint;index:idx_user_id;comment:用户ID"`
	ProductID         uint            `gorm:"column:product_id;type:bigint;index:idx_product_id;comment:产品ID"`
	CoinKey           string          `gorm:"column:coin_key;type:varchar(50);comment:币种标识"`
	InvestAmount      string          `gorm:"column:invest_amount;type:varchar(50);comment:投资金额"`
	OrderStatus       int8            `gorm:"column:order_status;type:tinyint;default:0;comment:订单状态 0-待确认 1-确认中 2-已确认 3-已取消"`
	Status            int8            `gorm:"column:status;type:tinyint;default:0;comment:状态 0-持仓中 1-已赎回 2-已到期"`
	Volume            decimal.Decimal `gorm:"column:volume;type:decimal(20,4);comment:剩余金额"`
	InterestStartDate int64           `gorm:"column:interest_start_date;type:bigint;comment:计息开始时间"`
	InterestEndDate   int64           `gorm:"column:interest_end_date;type:bigint;comment:计息结束时间"`
	OpenSub           int8            `gorm:"column:open_sub;type:tinyint;default:1;comment:是否开启自动申购 0-否 1-是"`
	CreatedAt         time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time       `gorm:"column:updated_at;autoUpdateTime"`
}

// TableName 指定表名
func (WthUserOrder) TableName() string {
	return "wth_user_order"
}
