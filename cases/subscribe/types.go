package subscribe

import (
	"github.com/shopspring/decimal"
)

// TestCase 申购测试用例基类
type TestCase struct {
	CaseID       string
	Module       string
	Priority     string
	Type         string
	Title        string
	Tags         []string
	PreCondition []string
}

// SubscribeInput 申购输入
type SubscribeInput struct {
	Coin         string // 币种，空表示自动创建
	SpecValue    int
	DeadlineType int
	Volume       decimal.Decimal
	MinVol       decimal.Decimal
	MaxVol       decimal.Decimal
	ProductOff   bool // 产品是否下架
	SpecOff      bool // 规格是否下架
}

// SubscribeExpect 申购预期结果
type SubscribeExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
	DBCheck        DBCheckPoints
}

// DBCheckPoints 数据库验证点
type DBCheckPoints struct {
	OrderCreated   bool
	OrderStatus    int
	VolumeMatch    bool
	BalanceChanged bool
	HisCreated     bool
}

// ValidationInput 验证测试输入
type ValidationInput struct {
	Coin            string // 空字符串表示空，"INVALID"表示不存在，空表示自动创建
	SpecValue       interface{} // 可以是 int 或 string
	DeadlineType    interface{} // 可以是 int 或无效值
	Volume          decimal.Decimal
	MinVol          decimal.Decimal
	MaxVol          decimal.Decimal
	PersonQuota     decimal.Decimal // 个人额度
	TotalQuota      decimal.Decimal // 产品总额度
	SoldAmount      decimal.Decimal // 已售金额
	DailyLimit      decimal.Decimal // 单日限额
	DailyUsed       decimal.Decimal // 当日已用额度
	ProductOff      bool            // 产品是否下架
	SpecOff         bool            // 规格是否下架
	CoinOff         bool            // 币种是否下架
	// 开户状态相关
	AccountStatus   int    // 开户状态: 0-未开户, 1-审核中, 2-已开户, 3-被拒绝
}

// ValidationExpect 验证测试预期
type ValidationExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// SecurityInput 和 SecurityExpect 已移至 security_cases.go
