package redeem

import (
	"github.com/shopspring/decimal"
)

// FuncCase 赎回功能测试用例
type FuncCase struct {
	TestCase
	Input    RedeemInput
	Expected RedeemExpect
}

// FuncCases 赎回功能测试用例表
var FuncCases = []FuncCase{
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_001",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "功能测试",
			Title:    "活期产品全额赎回成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{
				"已登录系统",
				"已创建活期产品并上架",
				"用户已有活期订单",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:    decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
		},
		Expected: RedeemExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				BalanceChanged: true,
				OrderCompleted: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_002",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "功能测试",
			Title:    "活期产品部分赎回成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"已创建活期产品并上架",
				"用户已有活期订单",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:    decimal.NewFromInt(500),
			MinVol:          decimal.NewFromInt(100),
		},
		Expected: RedeemExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				BalanceChanged: true,
				VolumeRemain:   decimal.NewFromInt(500),
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_003",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "功能测试",
			Title:    "定期产品全额赎回成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"已创建定期产品并上架",
				"用户已有定期订单",
				"订单已到期",
			},
		},
		Input: RedeemInput{
			DeadlineType:    1,
			SubscribeVolume: decimal.NewFromInt(5000),
			RedeemVolume:    decimal.NewFromInt(5000),
			MinVol:          decimal.NewFromInt(100),
		},
		Expected: RedeemExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				BalanceChanged: true,
				OrderCompleted: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_004",
			Module:   "用户理财订单",
			Priority: "低",
			Type:     "功能测试",
			Title:    "赎回金额为零失败",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户已有订单",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:    decimal.NewFromInt(0),
			MinVol:          decimal.NewFromInt(100),
		},
		Expected: RedeemExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "赎回金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_005",
			Module:   "用户理财订单",
			Priority: "低",
			Type:     "功能测试",
			Title:    "赎回金额为负数失败",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户已有订单",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:    decimal.NewFromInt(-100),
			MinVol:          decimal.NewFromInt(100),
		},
		Expected: RedeemExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "赎回金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_006",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "赎回金额超过持仓失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户已有订单",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:    decimal.NewFromInt(2000),
			MinVol:          decimal.NewFromInt(100),
		},
		Expected: RedeemExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "超过",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_007",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "订单号不存在赎回失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
			},
		},
		Input: RedeemInput{
			InvalidOrderNo:  true,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:    decimal.NewFromInt(1000),
		},
		Expected: RedeemExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "订单",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_008",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "多次赎回同一订单成功",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户余额充足",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(3000),
			RedeemVolume:    decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
		},
		Expected: RedeemExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				BalanceChanged: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_009",
			Module:   "用户理财订单",
			Priority: "低",
			Type:     "功能测试",
			Title:    "最小赎回金额赎回成功",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户余额充足",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:    decimal.NewFromInt(1),
			MinVol:          decimal.NewFromInt(100),
		},
		Expected: RedeemExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				BalanceChanged: true,
				VolumeRemain:   decimal.NewFromInt(999),
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_010",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "大额赎回成功",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户余额充足",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000000),
			RedeemVolume:    decimal.NewFromInt(1000000),
			MinVol:          decimal.NewFromInt(100),
			MaxVol:          decimal.NewFromInt(2000000),
		},
		Expected: RedeemExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				BalanceChanged: true,
				OrderCompleted: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_RED_FUNC_011",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "已完成订单赎回失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户已完成订单",
			},
		},
		Input: RedeemInput{
			SpecValue:       -1,
			DeadlineType:    0,
			SubscribeVolume: decimal.NewFromInt(1000),
			RedeemVolume:    decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
			OrderCompleted:  true,
		},
		Expected: RedeemExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "完成",
		},
	},
}

// GetFuncCaseByID 根据ID获取用例
func GetFuncCaseByID(caseID string) *FuncCase {
	for i := range FuncCases {
		if FuncCases[i].CaseID == caseID {
			return &FuncCases[i]
		}
	}
	return nil
}
