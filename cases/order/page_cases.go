package order

// PageCases 分页查询测试用例表
var PageCases = []PageCase{
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PAGE_001",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "分页查询成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{"已登录系统", "用户有多个订单"},
		},
		Input: PageInput{
			PageIndex:  1,
			PageSize:   10,
			HasOrder:   true,
			OrderCount: 3,
		},
		Expected: PageExpect{
			Success:       true,
			StatusCode:    200,
			MinOrderCount: 3,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PAGE_002",
			Module:   "用户订单查询",
			Priority: "中",
			Type:     "功能测试",
			Title:    "按币种过滤查询",
			Tags:     []string{"P1"},
			PreCondition: []string{"已登录系统", "用户有不同币种的订单"},
		},
		Input: PageInput{
			PageIndex:  1,
			PageSize:   10,
			HasOrder:   true,
			OrderCount: 1,
		},
		Expected: PageExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PAGE_003",
			Module:   "用户订单查询",
			Priority: "中",
			Type:     "功能测试",
			Title:    "按状态过滤查询",
			Tags:     []string{"P1"},
			PreCondition: []string{"已登录系统", "用户有不同状态的订单"},
		},
		Input: PageInput{
			PageIndex: 1,
			PageSize:  10,
			Status:    intPtr(0),
			HasOrder:  true,
		},
		Expected: PageExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PAGE_004",
			Module:   "用户订单查询",
			Priority: "中",
			Type:     "功能测试",
			Title:    "按规格值过滤查询",
			Tags:     []string{"P1"},
			PreCondition: []string{"已登录系统", "用户有不同规格值的订单"},
		},
		Input: PageInput{
			PageIndex: 1,
			PageSize:  10,
			SpecValue: intPtr(-1),
			HasOrder:  true,
		},
		Expected: PageExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PAGE_005",
			Module:   "用户订单查询",
			Priority: "低",
			Type:     "功能测试",
			Title:    "查询结果为空",
			Tags:     []string{"P2"},
			PreCondition: []string{"已登录系统", "用户没有订单"},
		},
		Input: PageInput{
			PageIndex:       1,
			PageSize:        10,
			UseNotExistCoin: true,
		},
		Expected: PageExpect{
			Success:    true,
			StatusCode: 200,
			EmptyList:  true,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PAGE_006",
			Module:   "用户订单查询",
			Priority: "低",
			Type:     "逆向测试",
			Title:    "页码非法",
			Tags:     []string{"P2"},
			PreCondition: []string{"已登录系统"},
		},
		Input: PageInput{
			PageIndex: 0,
			PageSize:  10,
		},
		Expected: PageExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PAGE_007",
			Module:   "用户订单查询",
			Priority: "低",
			Type:     "逆向测试",
			Title:    "每页大小非法",
			Tags:     []string{"P2"},
			PreCondition: []string{"已登录系统"},
		},
		Input: PageInput{
			PageIndex: 1,
			PageSize:  0,
		},
		Expected: PageExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_PAGE_008",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "未登录应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"未登录"},
		},
		Input: PageInput{
			PageIndex: 1,
			PageSize:  10,
		},
		Expected: PageExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "登录",
		},
	},
}

// GetPageCaseByID 根据ID获取分页查询用例
func GetPageCaseByID(caseID string) *PageCase {
	for i := range PageCases {
		if PageCases[i].CaseID == caseID {
			return &PageCases[i]
		}
	}
	return nil
}
