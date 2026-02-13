package redeem

import (
	"github.com/shopspring/decimal"
)

// ConcurrentCase 赎回并发测试用例
type ConcurrentCase struct {
	CaseID       string
	Module       string
	Priority     string
	Type         string
	Title        string
	Tags         []string
	PreCondition []string
	TestData     ConcurrentTestData
	Expect       ConcurrentExpect
}

// ConcurrentTestData 赎回并发测试数据
type ConcurrentTestData struct {
	SpecValue         int
	DeadlineType      int
	SubscribeVolume   decimal.Decimal
	DailyLimit        decimal.Decimal
	MinVol            decimal.Decimal
	ConcurrentCount   int
	RedeemVolume      decimal.Decimal
	FullRedeem        bool
	CreateMultiOrders bool
	OrderCount        int
	MixedOperation    bool
}

// ConcurrentExpect 赎回并发测试预期结果
type ConcurrentExpect struct {
	SuccessCount          int
	MaxSuccessCount       int
	SoldAmountNotNegative bool
	BalanceNotNegative    bool
	OrderStatusComplete   bool
	VolumeRemainCorrect   bool
}

// ConcurrentCases 赎回并发测试用例表
var ConcurrentCases = []ConcurrentCase{
	{
		CaseID:   "WTH_RED_CON_001",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发赎回定期不应重复赎回",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"已创建定期产品并上架",
			"用户已申购定期产品",
		},
		TestData: ConcurrentTestData{
			DeadlineType:    1,
			SubscribeVolume: decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 10,
			RedeemVolume:    decimal.NewFromInt(1000),
			FullRedeem:      true,
		},
		Expect: ConcurrentExpect{
			SuccessCount:        1,
			OrderStatusComplete: true,
		},
	},
	{
		CaseID:   "WTH_RED_CON_002",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发全额赎回活期不应重复赎回",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"已创建活期产品并上架",
			"用户已申购活期产品",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 10,
			RedeemVolume:    decimal.NewFromInt(1000),
			FullRedeem:      true,
		},
		Expect: ConcurrentExpect{
			SuccessCount:        1,
			OrderStatusComplete: true,
		},
	},
	{
		CaseID:   "WTH_RED_CON_003",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发部分赎回活期不应超额度",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"用户已申购1000活期产品",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 5,
			RedeemVolume:    decimal.NewFromInt(400),
			FullRedeem:      false,
		},
		Expect: ConcurrentExpect{
			MaxSuccessCount:     2,
			VolumeRemainCorrect: true,
		},
	},
	{
		CaseID:   "WTH_RED_CON_004",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发赎回不应超日限额",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"用户已申购多个订单",
		},
		TestData: ConcurrentTestData{
			SpecValue:         -1,
			DeadlineType:      0,
			SubscribeVolume:   decimal.NewFromInt(1000),
			DailyLimit:        decimal.NewFromInt(3000),
			MinVol:            decimal.NewFromInt(100),
			ConcurrentCount:   5,
			RedeemVolume:      decimal.NewFromInt(1000),
			FullRedeem:        true,
			CreateMultiOrders: true,
			OrderCount:        5,
		},
		Expect: ConcurrentExpect{
			MaxSuccessCount: 3,
		},
	},
	{
		CaseID:   "WTH_RED_CON_005",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发赎回sold_amount不应为负",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"产品已售金额为1000",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 10,
			RedeemVolume:    decimal.NewFromInt(1000),
			FullRedeem:      true,
		},
		Expect: ConcurrentExpect{
			SoldAmountNotNegative: true,
		},
	},
	{
		CaseID:   "WTH_RED_CON_006",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发竞态条件应防止",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"用户有订单",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(5000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 20,
			RedeemVolume:    decimal.NewFromInt(500),
			MixedOperation:  true,
		},
		Expect: ConcurrentExpect{
			BalanceNotNegative:    true,
			SoldAmountNotNegative: true,
		},
	},
}

// GetConcurrentCaseByID 根据ID获取并发用例
func GetConcurrentCaseByID(caseID string) *ConcurrentCase {
	for i := range ConcurrentCases {
		if ConcurrentCases[i].CaseID == caseID {
			return &ConcurrentCases[i]
		}
	}
	return nil
}
