package redeem

import (
	"github.com/shopspring/decimal"
)

// TestCase 赎回测试用例基类
type TestCase struct {
	CaseID       string
	Module       string
	Priority     string
	Type         string
	Title        string
	Tags         []string
	PreCondition []string
}

// RedeemInput 赎回输入
type RedeemInput struct {
	SpecValue       int
	DeadlineType    int
	SubscribeVolume decimal.Decimal // 申购金额
	RedeemVolume    decimal.Decimal // 赎回金额
	MinVol          decimal.Decimal
	MaxVol          decimal.Decimal
	InvalidOrderNo  bool // 是否使用无效订单号
	OrderCompleted  bool // 订单是否已完成
}

// RedeemExpect 赎回预期结果
type RedeemExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
	DBCheck        DBCheckPoints
}

// DBCheckPoints 赎回数据库验证点
type DBCheckPoints struct {
	BalanceChanged bool
	OrderCompleted bool
	VolumeRemain   decimal.Decimal // 剩余持仓
}

// ValidationInput 赎回验证测试输入
type ValidationInput struct {
	OrderNo         string // 空表示空，"INVALID"表示不存在，空表示自动创建
	RedeemVolume    decimal.Decimal
	SubscribeVolume decimal.Decimal
	OtherUserOrder  bool // 是否使用其他用户订单
	OrderCompleted  bool // 订单是否已完成
	DailyLimit      decimal.Decimal // 单日限额
}

// ValidationExpect 赎回验证测试预期
type ValidationExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// SecurityInput 和 SecurityExpect 已移至 security_cases.go
