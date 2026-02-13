package models

import (
	"time"
)

// WthSpec 规格模型
type WthSpec struct {
	ID            int64      `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	SpecValue     int        `gorm:"column:spec_value;not null" json:"specValue"`
	SpecName      string     `gorm:"column:spec_name;size:50" json:"specName"`
	SpecKey       string     `gorm:"column:spec_key;size:50" json:"specKey"`
	DeadlineType  int        `gorm:"column:deadline_type;default:0" json:"deadlineType"`
	ShelvesStatus int        `gorm:"column:shelves_status;default:0" json:"shelvesStatus"`
	Remark        string     `gorm:"column:remark;size:500" json:"remark"`
	CreatedAt     time.Time  `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deletedAt,omitempty"`
}

// TableName 指定表名
func (WthSpec) TableName() string {
	return "wth_spec"
}
