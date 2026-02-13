package models

import "time"

// ConfigLanguage 多语言配置表
type ConfigLanguage struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Type      int       `gorm:"column:type;type:tinyint;default:1" json:"type"`           // 类型：1-股票，2-币种，9-站内信
	ConfigKey string    `gorm:"column:config_key;type:varchar(100)" json:"config_key"`    // 配置key
	LangKey   string    `gorm:"column:lang_key;type:varchar(50)" json:"lang_key"`         // 语言key
	Content   string    `gorm:"column:content;type:text" json:"content"`                  // 多语言文本
	Meta      string    `gorm:"column:meta;type:varchar(100)" json:"meta"`                // 描述
	SubType   int       `gorm:"column:sub_type;type:tinyint" json:"sub_type"`             // 细分类型
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

// TableName 指定表名
func (ConfigLanguage) TableName() string {
	return "config_language"
}

// 语言类型常量
const (
	LanguageTypeCoin  = 2 // 币种类型
	LanguageTypeStock = 1 // 股票类型
	LanguageTypeMsg   = 9 // 站内信类型
)
