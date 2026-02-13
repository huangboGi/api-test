package subscribe

import (
	"github.com/shopspring/decimal"
)

// ConcurrentCase 申购并发测试用例
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

// ConcurrentTestData 并发测试数据
type ConcurrentTestData struct {
	SpecValue       int
	DeadlineType    int // 0-活期, 1-定期
	TotalQuota      decimal.Decimal
	PersonQuota     decimal.Decimal
	DailyLimit      decimal.Decimal
	MinVol          decimal.Decimal
	ConcurrentCount int
	VolumePerUser   decimal.Decimal
	SameUser        bool // 是否同一用户并发
	BaseUserID      int
}

// ConcurrentExpect 并发测试预期结果
type ConcurrentExpect struct {
	MaxSuccessCount     int
	SoldAmountNotExceed decimal.Decimal
	OrderCount          int
	OrderVolume         decimal.Decimal
	SoldAmountCorrect   bool
}

// ConcurrentCases 申购并发测试用例表
var ConcurrentCases = []ConcurrentCase{
	{
		CaseID:   "WTH_SUB_CON_001",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发申购定期不应超额度",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"币种、规格、产品已创建并上架",
			"产品总额度为100000",
		},
		TestData: ConcurrentTestData{
			DeadlineType:    1,
			TotalQuota:      decimal.NewFromInt(100000),
			PersonQuota:     decimal.NewFromInt(10000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 20,
			VolumePerUser:   decimal.NewFromInt(6000),
			SameUser:        false,
			BaseUserID:      10000,
		},
		Expect: ConcurrentExpect{
			MaxSuccessCount:     16,
			SoldAmountNotExceed: decimal.NewFromInt(100000),
			SoldAmountCorrect:   true,
		},
	},
	{
		CaseID:   "WTH_SUB_CON_002",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发申购活期不应超额度",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"产品总额度为100000",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			TotalQuota:      decimal.NewFromInt(100000),
			PersonQuota:     decimal.NewFromInt(10000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 20,
			VolumePerUser:   decimal.NewFromInt(6000),
			SameUser:        false,
			BaseUserID:      10100,
		},
		Expect: ConcurrentExpect{
			MaxSuccessCount:     16,
			SoldAmountNotExceed: decimal.NewFromInt(100000),
			SoldAmountCorrect:   true,
		},
	},
	{
		CaseID:   "WTH_SUB_CON_003",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发申购不应超日限额",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"日限额为100000",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			TotalQuota:      decimal.NewFromInt(1000000),
			PersonQuota:     decimal.NewFromInt(10000),
			DailyLimit:      decimal.NewFromInt(100000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 20,
			VolumePerUser:   decimal.NewFromInt(6000),
			SameUser:        false,
			BaseUserID:      10200,
		},
		Expect: ConcurrentExpect{
			MaxSuccessCount: 16,
			SoldAmountCorrect: true,
		},
	},
	{
		CaseID:   "WTH_SUB_CON_004",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发申购不应超个人额度",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"个人额度为10000",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			TotalQuota:      decimal.NewFromInt(1000000),
			PersonQuota:     decimal.NewFromInt(10000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 5,
			VolumePerUser:   decimal.NewFromInt(3000),
			SameUser:        true,
		},
		Expect: ConcurrentExpect{
			MaxSuccessCount: 3,
		},
	},
	{
		CaseID:   "WTH_SUB_CON_005",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发申购sold_amount不应超",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"产品总额度为100000",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			TotalQuota:      decimal.NewFromInt(100000),
			PersonQuota:     decimal.NewFromInt(10000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 20,
			VolumePerUser:   decimal.NewFromInt(6000),
			SameUser:        false,
			BaseUserID:      10300,
		},
		Expect: ConcurrentExpect{
			SoldAmountNotExceed: decimal.NewFromInt(100000),
		},
	},
	{
		CaseID:   "WTH_SUB_CON_006",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "并发测试",
		Title:    "并发活期申购累加正确",
		Tags:     []string{"P1", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"活期产品已创建并上架",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			TotalQuota:      decimal.NewFromInt(100000),
			PersonQuota:     decimal.NewFromInt(10000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 5,
			VolumePerUser:   decimal.NewFromInt(1000),
			SameUser:        true,
		},
		Expect: ConcurrentExpect{
			OrderCount:        1,
			OrderVolume:       decimal.NewFromInt(5000),
			SoldAmountCorrect: true,
		},
	},
	{
		CaseID:   "WTH_SUB_CON_007",
		Module:   "用户理财订单",
		Priority: "极高",
		Type:     "并发测试",
		Title:    "并发竞态条件应防止",
		Tags:     []string{"P0", "concurrent"},
		PreCondition: []string{
			"已登录系统",
			"产品总额度为20000",
		},
		TestData: ConcurrentTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			TotalQuota:      decimal.NewFromInt(20000),
			PersonQuota:     decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentCount: 100,
			VolumePerUser:   decimal.NewFromInt(300),
			SameUser:        false,
			BaseUserID:      10400,
		},
		Expect: ConcurrentExpect{
			SoldAmountNotExceed: decimal.NewFromInt(20000),
			SoldAmountCorrect:   true,
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
