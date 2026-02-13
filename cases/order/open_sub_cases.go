package order

// OpenSubCases 自动申购设置测试用例表
var OpenSubCases = []OpenSubCase{
	{
		TestCase: TestCase{
			CaseID:   "ORDER_OPEN_001",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "开启自动申购成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{"已登录系统", "用户有理财订单"},
		},
		Input: OpenSubInput{
			HasOrder:       true,
			OpenSub:        1,
			InitialOpenSub: 0,
		},
		Expected: OpenSubExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_OPEN_002",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "关闭自动申购成功",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统", "用户有理财订单"},
		},
		Input: OpenSubInput{
			HasOrder:       true,
			OpenSub:        0,
			InitialOpenSub: 1,
		},
		Expected: OpenSubExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_OPEN_003",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "订单不属于当前用户应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统", "存在其他用户的订单"},
		},
		Input: OpenSubInput{
			OtherUserOrder: true,
			OpenSub:        1,
		},
		Expected: OpenSubExpect{
			Success:        false,
			StatusCode:     404,
			ErrMsgContains: "订单",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_OPEN_004",
			Module:   "用户订单查询",
			Priority: "中",
			Type:     "逆向测试",
			Title:    "订单不存在应失败",
			Tags:     []string{"P1"},
			PreCondition: []string{"已登录系统"},
		},
		Input: OpenSubInput{
			NotExistOrder: true,
			OpenSub:       1,
		},
		Expected: OpenSubExpect{
			Success:        false,
			StatusCode:     404,
			ErrMsgContains: "订单",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_OPEN_005",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "orderId为空应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统"},
		},
		Input: OpenSubInput{
			EmptyOrderID: true,
			OpenSub:      1,
		},
		Expected: OpenSubExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "id",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_OPEN_006",
			Module:   "用户订单查询",
			Priority: "低",
			Type:     "逆向测试",
			Title:    "openSub值非法应失败",
			Tags:     []string{"P2"},
			PreCondition: []string{"已登录系统", "用户有理财订单"},
		},
		Input: OpenSubInput{
			HasOrder:       true,
			OpenSub:        2,
			InvalidOpenSub: true,
		},
		Expected: OpenSubExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "openSub",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_OPEN_007",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "逆向测试",
			Title:    "openSub为空应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"已登录系统"},
		},
		Input: OpenSubInput{
			HasOrder:     true,
			EmptyOpenSub: true,
		},
		Expected: OpenSubExpect{
			Success:        false,
			StatusCode:     400,
			ErrMsgContains: "openSub",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_OPEN_008",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "未登录应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"未登录"},
		},
		Input: OpenSubInput{
			OrderID: 1,
			OpenSub: 1,
		},
		Expected: OpenSubExpect{
			Success:        false,
			StatusCode:     401,
			ErrMsgContains: "登录",
		},
	},
}

// GetOpenSubCaseByID 根据ID获取自动申购设置用例
func GetOpenSubCaseByID(caseID string) *OpenSubCase {
	for i := range OpenSubCases {
		if OpenSubCases[i].CaseID == caseID {
			return &OpenSubCases[i]
		}
	}
	return nil
}
