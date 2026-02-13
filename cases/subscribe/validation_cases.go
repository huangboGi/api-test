package subscribe

import (
	"github.com/shopspring/decimal"
)

// ValidationCase 申购验证测试用例
type ValidationCase struct {
	TestCase
	Input    ValidationInput
	Expected ValidationExpect
}

// ValidationCases 申购验证测试用例表
var ValidationCases = []ValidationCase{
	// ========== 参数验证 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_001",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "coin为空应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
			},
		},
		Input: ValidationInput{
			Coin:         "", // 空
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "coin",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_002",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "coin不存在应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
			},
		},
		Input: ValidationInput{
			Coin:         "NONEXISTCOIN",
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "币种",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_003",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "币种未上架应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"币种已创建但未上架",
			},
		},
		Input: ValidationInput{
			CoinOff:      true,
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "币种",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_004",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "specValue为空应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
			},
		},
		Input: ValidationInput{
			SpecValue:    "", // 空字符串
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "specValue",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_005",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "specValue不存在应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
			},
		},
		Input: ValidationInput{
			SpecValue:    9999,
			DeadlineType: 1,
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "规格",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_006",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "规格已下架应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"规格已创建但未上架",
			},
		},
		Input: ValidationInput{
			SpecOff:      true,
			DeadlineType: 1,
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "规格",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_007",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "产品已下架应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已创建但未上架",
			},
		},
		Input: ValidationInput{
			ProductOff:   true,
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "产品",
		},
	},
	// ========== 金额验证 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_008",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "申购金额为0应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
			},
		},
		Input: ValidationInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(0),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "申购金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_009",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "申购金额为负数应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
			},
		},
		Input: ValidationInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(-1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "申购金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_010",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "申购金额小于最小申购额应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
				"最小申购额为1000",
			},
		},
		Input: ValidationInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(500),
			MinVol:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "最小",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_011",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "申购金额超过个人额度应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
				"个人额度为10000",
			},
		},
		Input: ValidationInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(15000),
			MinVol:       decimal.NewFromInt(100),
			PersonQuota:  decimal.NewFromInt(10000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "个人额度",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_012",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "申购金额超过产品总额度应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
				"总额度100000，已售50000",
			},
		},
		Input: ValidationInput{
			SpecValue:   -1,
			DeadlineType: 0,
			Volume:      decimal.NewFromInt(60000),
			MinVol:      decimal.NewFromInt(100),
			TotalQuota:  decimal.NewFromInt(100000),
			SoldAmount:  decimal.NewFromInt(50000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "总额度",
		},
	},
	// ========== 参数类型验证 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_013",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "逆向测试",
			Title:    "deadlineType值非法应失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
			},
		},
		Input: ValidationInput{
			SpecValue:    -1,
			DeadlineType: 2, // 非法值
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "期限类型",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_014",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "逆向测试",
			Title:    "deadlineType与产品不匹配应失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"活期产品已上架",
			},
		},
		Input: ValidationInput{
			SpecValue:    -1,
			DeadlineType: 1, // 定期，与活期产品不匹配
			Volume:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "期限类型",
		},
	},
	// ========== 边界测试 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_015",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "边界测试",
			Title:    "申购金额边界值测试-最小申购额",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
				"最小申购额为1000",
			},
		},
		Input: ValidationInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
			MinVol:       decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	// ========== 日限额边界测试 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_016",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "边界测试",
			Title:    "申购金额等于日限额应成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
				"日限额为10000",
				"用户当日已申购0",
			},
		},
		Input: ValidationInput{
			SpecValue:   -1,
			DeadlineType: 0,
			Volume:      decimal.NewFromInt(10000),
			MinVol:      decimal.NewFromInt(100),
			DailyLimit:  decimal.NewFromInt(10000),
		},
		Expected: ValidationExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_017",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "边界测试",
			Title:    "申购金额超过日限额应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
				"日限额为10000",
				"用户当日已申购0",
			},
		},
		Input: ValidationInput{
			SpecValue:   -1,
			DeadlineType: 0,
			Volume:      decimal.NewFromInt(10001),
			MinVol:      decimal.NewFromInt(100),
			DailyLimit:  decimal.NewFromInt(10000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "日限额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_018",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "边界测试",
			Title:    "当日已申购金额+本次申购超过日限额应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
				"日限额为10000",
				"用户当日已申购8000",
			},
		},
		Input: ValidationInput{
			SpecValue:   -1,
			DeadlineType: 0,
			Volume:      decimal.NewFromInt(3000),
			MinVol:      decimal.NewFromInt(100),
			DailyLimit:  decimal.NewFromInt(10000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "日限额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_019",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "边界测试",
			Title:    "当日已申购金额+本次申购等于日限额应成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"产品已上架",
				"日限额为10000",
				"用户当日已申购8000",
			},
		},
		Input: ValidationInput{
			SpecValue:   -1,
			DeadlineType: 0,
			Volume:      decimal.NewFromInt(2000),
			MinVol:      decimal.NewFromInt(100),
			DailyLimit:  decimal.NewFromInt(10000),
		},
		Expected: ValidationExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	// ========== 用户开户状态验证 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_020",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "用户未开户应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"用户未完成开户",
				"产品已上架",
			},
		},
		Input: ValidationInput{
			SpecValue:   -1,
			DeadlineType: 0,
			Volume:      decimal.NewFromInt(1000),
			MinVol:      decimal.NewFromInt(100),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "开户",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_021",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "逆向测试",
			Title:    "用户开户审核中应失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"用户开户状态为审核中",
				"产品已上架",
			},
		},
		Input: ValidationInput{
			SpecValue:   -1,
			DeadlineType: 0,
			Volume:      decimal.NewFromInt(1000),
			MinVol:      decimal.NewFromInt(100),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "审核",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_VAL_022",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "逆向测试",
			Title:    "用户开户被拒绝应失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"用户开户状态为被拒绝",
				"产品已上架",
			},
		},
		Input: ValidationInput{
			SpecValue:   -1,
			DeadlineType: 0,
			Volume:      decimal.NewFromInt(1000),
			MinVol:      decimal.NewFromInt(100),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "开户",
		},
	},
}

// GetValidationCaseByID 根据ID获取用例
func GetValidationCaseByID(caseID string) *ValidationCase {
	for i := range ValidationCases {
		if ValidationCases[i].CaseID == caseID {
			return &ValidationCases[i]
		}
	}
	return nil
}
