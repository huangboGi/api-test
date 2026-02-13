package models

import (
	"time"
)

// WthCoinConfig 币种配置模型（与主项目保持一致）
type WthCoinConfig struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Coin      string    `gorm:"column:coin;type:varchar(50);not null;uniqueIndex:uk_coin" json:"coin"` // 币种标识(核心字段)
	CoinKey   string    `gorm:"column:coin_key;type:varchar(50);index:idx_coin_key" json:"coinKey"`    // 币种多语言key
	Shelves   int       `gorm:"column:shelves;type:int;default:0;comment:0下架1上架" json:"shelves"`       // 0下架1上架
	PricePre  int       `gorm:"column:price_pre;type:int;default:2" json:"pricePre"`                   // 价格精度
	Tag       string    `gorm:"column:tag;type:varchar(50)" json:"tag"`                                // 标签
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

// TableName 指定表名
func (WthCoinConfig) TableName() string {
	return "wth_coin_config"
}
