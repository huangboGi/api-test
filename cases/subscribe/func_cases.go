package subscribe

import (
	"github.com/shopspring/decimal"
)

// FuncCase 申购功能测试用例
type FuncCase struct {
	TestCase
	Input    SubscribeInput
	Expected SubscribeExpect
}

// FuncCases 申购功能测试用例表
var FuncCases = []FuncCase{
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_001",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "功能测试",
			Title:    "活期产品申购成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{
				"已登录系统",
				"已创建活期产品并上架",
				"用户余额充足",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
			MinVol:       decimal.NewFromInt(100),
		},
		Expected: SubscribeExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				OrderCreated:   true,
				VolumeMatch:    true,
				BalanceChanged: true,
				HisCreated:     true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_002",
			Module:   "用户理财订单",
			Priority: "高",
			Type:     "功能测试",
			Title:    "定期产品申购成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录系统",
				"已创建定期产品并上架",
				"用户余额充足",
			},
		},
		Input: SubscribeInput{
			DeadlineType: 1,
			Volume:       decimal.NewFromInt(5000),
			MinVol:       decimal.NewFromInt(100),
		},
		Expected: SubscribeExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				OrderCreated: true,
				VolumeMatch:  true,
				HisCreated:   true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_003",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "最小申购金额申购成功",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并设置最小申购金额",
				"用户余额充足",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(100),
			MinVol:       decimal.NewFromInt(100),
		},
		Expected: SubscribeExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				OrderCreated: true,
				VolumeMatch:  true,
				HisCreated:   true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_004",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "最大申购金额申购成功",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并设置最大申购金额",
				"用户余额充足",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(10000),
			MinVol:       decimal.NewFromInt(100),
			MaxVol:       decimal.NewFromInt(10000),
		},
		Expected: SubscribeExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				OrderCreated: true,
				VolumeMatch:  true,
				HisCreated:   true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_005",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "多次申购同一产品成功",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户余额充足",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
			MinVol:       decimal.NewFromInt(100),
		},
		Expected: SubscribeExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheckPoints{
				OrderCreated: true,
				HisCreated:   true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_006",
			Module:   "用户理财订单",
			Priority: "低",
			Type:     "功能测试",
			Title:    "申购金额为零失败",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(0),
			MinVol:       decimal.NewFromInt(100),
		},
		Expected: SubscribeExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "申购金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_007",
			Module:   "用户理财订单",
			Priority: "低",
			Type:     "功能测试",
			Title:    "申购金额为负数失败",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(-100),
			MinVol:       decimal.NewFromInt(100),
		},
		Expected: SubscribeExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "申购金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_008",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "余额不足申购失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品并上架",
				"用户余额不足",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000000),
			MinVol:       decimal.NewFromInt(100),
		},
		Expected: SubscribeExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "余额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_009",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "产品下架后申购失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品",
				"产品已下架",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
			MinVol:       decimal.NewFromInt(100),
			ProductOff:   true,
		},
		Expected: SubscribeExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "下架",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SUB_FUNC_010",
			Module:   "用户理财订单",
			Priority: "中",
			Type:     "功能测试",
			Title:    "规格下架后申购失败",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录系统",
				"已创建产品",
				"规格已下架",
			},
		},
		Input: SubscribeInput{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
			MinVol:       decimal.NewFromInt(100),
			SpecOff:      true,
		},
		Expected: SubscribeExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "下架",
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
