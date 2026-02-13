package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// WthProduct 理财产品模型
type WthProduct struct {
	ID             int64           `gorm:"column:id;primarykey;autoIncrement"`
	ProductKey     string          `gorm:"column:product_key;type:varchar(50);uniqueIndex:uk_product_key;comment:产品标识"`
	ProductName    string          `gorm:"column:product_name;type:varchar(200);comment:产品名称"`
	Coin           string          `gorm:"column:coin;type:varchar(50);comment:币种"`
	CoinKey        string          `gorm:"column:coin_key;type:varchar(50);comment:币种标识"`
	ClassifyKey    string          `gorm:"column:classify_key;type:varchar(50);comment:分类标识"`
	SpecValue      int             `gorm:"column:spec_value;type:int;comment:规格值"`
	DeadlineType   int             `gorm:"column:deadline_type;type:int;default:0;comment:期限类型 0-活期 1-定期"`
	AnnualAte      decimal.Decimal `gorm:"column:annual_ate;type:decimal(10,4);comment:年化收益率"`
	Tag            string          `gorm:"column:tag;type:varchar(100);comment:标签"`
	MinVol         decimal.Decimal `gorm:"column:min_vol;type:decimal(20,4);comment:最小起购金额"`
	UseQuotaTotal  decimal.Decimal `gorm:"column:use_quota_total;type:decimal(20,4);comment:总配额"`
	PersonQuota    decimal.Decimal `gorm:"column:person_quota;type:decimal(20,4);comment:个人配额"`
	SoldAmount     decimal.Decimal `gorm:"column:sold_amount;type:decimal(20,4);default:0;comment:已售金额"`
	OpenSub        int             `gorm:"column:open_sub;type:int;default:1;comment:是否可申购 0-否 1-是"`
	ShelvesStatus  int             `gorm:"column:shelves_status;type:int;default:0;comment:上架状态 0-下架 1-上架"`
	ExtraAnnualAte decimal.Decimal `gorm:"column:extra_annual_ate;type:decimal(10,4);comment:额外年化收益率"`
	DailyMaximum   decimal.Decimal `gorm:"column:daily_maximum;type:decimal(20,4);comment:每日最大申购金额"`
	Sort           int             `gorm:"column:sort;type:int;default:0;comment:排序"`
	CreatedAt      time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time       `gorm:"column:updated_at;autoUpdateTime"`
}

// TableName 指定表名
func (WthProduct) TableName() string {
	return "wth_product"
}
