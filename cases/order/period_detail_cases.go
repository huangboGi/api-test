package order

// PeriodDetailCases 期间详情测试用例表
var PeriodDetailCases = []PeriodDetailCase{
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PERIOD_001",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询定期产品期间详情成功",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统", "用户有定期理财订单"},
		},
		Input: PeriodDetailInput{
			HasOrder: true,
			IsFixed:  true,
		},
		Expected: PeriodDetailExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PERIOD_002",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询活期产品期间详情成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{"已登录系统", "用户有活期理财订单"},
		},
		Input: PeriodDetailInput{
			HasOrder: true,
			IsFixed:  false,
		},
		Expected: PeriodDetailExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PERIOD_003",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "订单不属于当前用户应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统", "存在其他用户的订单"},
		},
		Input: PeriodDetailInput{
			OtherUserOrder: true,
		},
		Expected: PeriodDetailExpect{
			Success:        false,
			StatusCode:     404,
			ErrMsgContains: "订单",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PERIOD_004",
			Module:   "用户订单查询",
			Priority: "中",
			Type:     "逆向测试",
			Title:    "订单不存在应失败",
			Tags:     []string{"P1"},
			PreCondition: []string{"已登录系统"},
		},
		Input: PeriodDetailInput{
			NotExistOrder: true,
		},
		Expected: PeriodDetailExpect{
			Success:        false,
			StatusCode:     404,
			ErrMsgContains: "订单",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PERIOD_005",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "orderId为空应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统"},
		},
		Input: PeriodDetailInput{
			EmptyOrderID: true,
		},
		Expected: PeriodDetailExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "id",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PERIOD_006",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "未登录应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"未登录"},
		},
		Input: PeriodDetailInput{
			OrderID: 1,
		},
		Expected: PeriodDetailExpect{
			Success:        false,
			StatusCode:     401,
			ErrMsgContains: "登录",
		},
	},
}

// GetPeriodDetailCaseByID 根据ID获取期间详情用例
func GetPeriodDetailCaseByID(caseID string) *PeriodDetailCase {
	for i := range PeriodDetailCases {
		if PeriodDetailCases[i].CaseID == caseID {
			return &PeriodDetailCases[i]
		}
	}
	return nil
}
