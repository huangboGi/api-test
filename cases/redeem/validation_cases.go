package redeem

import (
	"github.com/shopspring/decimal"
)

// ValidationCase 赎回验证测试用例
type ValidationCase struct {
	TestCase
	Input    ValidationInput
	Expected ValidationExpect
}

// ValidationCases 赎回验证测试用例表
var ValidationCases = []ValidationCase{
	// ========== 订单号验证 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_001",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "orderId为空应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
			},
		},
		Input: ValidationInput{
			OrderNo:      "", // 空
			RedeemVolume: decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "orderId",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_002",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "orderId不存在应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
			},
		},
		Input: ValidationInput{
			OrderNo:      "INVALID",
			RedeemVolume: decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "订单",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_003",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "订单已完成应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"用户有已完成状态的订单",
			},
		},
		Input: ValidationInput{
			OrderCompleted: true,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:   decimal.NewFromInt(1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "完成",
		},
	},
	// ========== 金额验证 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_004",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "赎回金额为0应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"用户有订单",
			},
		},
		Input: ValidationInput{
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:   decimal.NewFromInt(0),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "赎回金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_005",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "赎回金额为负数应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"用户有订单",
			},
		},
		Input: ValidationInput{
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:   decimal.NewFromInt(-1000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "赎回金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_006",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "赎回金额超过订单金额应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"用户有订单1000",
			},
		},
		Input: ValidationInput{
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:   decimal.NewFromInt(1500),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "超过",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_007",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "赎回金额超过单日限额应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"用户有订单",
				"今日已赎回80000，单日限额100000",
			},
		},
		Input: ValidationInput{
			SubscribeVolume: decimal.NewFromInt(10000),
			RedeemVolume:   decimal.NewFromInt(30000),
			DailyLimit:     decimal.NewFromInt(100000),
		},
		Expected: ValidationExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "单日限额",
		},
	},
	// ========== 边界测试 ==========
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_008",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "边界测试",
			Title:    "赎回金额边界值测试-最小值",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"用户有订单",
			},
		},
		Input: ValidationInput{
			SubscribeVolume: decimal.NewFromInt(10000),
			RedeemVolume:   decimal.NewFromInt(1),
		},
		Expected: ValidationExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_VAL_009",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "边界测试",
			Title:    "活期部分赎回边界值测试-接近全额",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"用户有活期订单",
			},
		},
		Input: ValidationInput{
			SubscribeVolume: decimal.NewFromInt(10000),
			RedeemVolume:   decimal.NewFromInt(9999),
		},
		Expected: ValidationExpect{
			Success:    true,
			StatusCode: 200,
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
