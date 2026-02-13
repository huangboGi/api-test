package models

import (
	"github.com/shopspring/decimal"
)

// WthUserSubscribeHis 用户申购历史表
type WthUserSubscribeHis struct {
	ID        int64           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Coin      string          `gorm:"column:coin;size:50" json:"coin"`                // 币种
	CoinKey   string          `gorm:"column:coin_key;size:50" json:"coinKey"`         // 币种多语言key
	UserID    int64           `gorm:"column:user_id;index" json:"userId"`             // 用户id
	ProductID int64           `gorm:"column:product_id;index" json:"productId"`       // 申购产品id
	OrderID   int64           `gorm:"column:order_id;index" json:"orderId"`           // 订单id
	Volume    decimal.Decimal `gorm:"column:volume;type:decimal(20,8)" json:"volume"` // 申购数量
	Type      int             `gorm:"column:type;default:0" json:"type"`              // 状态：0-普通；1-自动申购
	CTime     int64           `gorm:"column:ctime" json:"createdAt"`                  // 创建时间（Unix时间戳）
	UTime     int64           `gorm:"column:utime" json:"updatedAt"`                  // 更新时间（Unix时间戳）
}

// TableName 指定表名
func (WthUserSubscribeHis) TableName() string {
	return "wth_user_subscribe_his"
}
