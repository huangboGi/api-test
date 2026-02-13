package order

// DetailCases 订单详情测试用例表
var DetailCases = []DetailCase{
	{
		TestCase: TestCase{
			CaseID:   "ORDER_DETAIL_001",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询订单详情成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{"已登录系统", "用户有理财订单"},
		},
		Input: DetailInput{
			HasOrder: true,
		},
		Expected: DetailExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_DETAIL_002",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "订单不属于当前用户应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统", "存在其他用户的订单"},
		},
		Input: DetailInput{
			OtherUserOrder: true,
		},
		Expected: DetailExpect{
			Success:        false,
			StatusCode:     404,
			ErrMsgContains: "订单",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_DETAIL_003",
			Module:   "用户订单查询",
			Priority: "中",
			Type:     "逆向测试",
			Title:    "订单不存在应失败",
			Tags:     []string{"P1"},
			PreCondition: []string{"已登录系统"},
		},
		Input: DetailInput{
			NotExistOrder: true,
		},
		Expected: DetailExpect{
			Success:        false,
			StatusCode:     404,
			ErrMsgContains: "订单",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_DETAIL_004",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "orderId为空应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统"},
		},
		Input:    DetailInput{},
		Expected: DetailExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "id",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_DETAIL_005",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "未登录应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"未登录"},
		},
		Input: DetailInput{
			OrderID: 1,
		},
		Expected: DetailExpect{
			Success:        false,
			StatusCode:     401,
			ErrMsgContains: "登录",
		},
	},
}

// GetDetailCaseByID 根据ID获取订单详情用例
func GetDetailCaseByID(caseID string) *DetailCase {
	for i := range DetailCases {
		if DetailCases[i].CaseID == caseID {
			return &DetailCases[i]
		}
	}
	return nil
}
