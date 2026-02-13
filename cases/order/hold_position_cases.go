package order

// HoldPositionCases 持仓查询测试用例表
var HoldPositionCases = []HoldPositionCase{
	{
		TestCase: TestCase{
			CaseID:   "ORDER_POS_001",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询所有持仓成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{"已登录系统", "用户有多个理财订单"},
		},
		Input: HoldPositionInput{
			HasOrder: true,
		},
		Expected: HoldPositionExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_POS_002",
			Module:   "用户订单查询",
			Priority: "低",
			Type:     "功能测试",
			Title:    "查询结果为空",
			Tags:     []string{"P2"},
			PreCondition: []string{"已登录系统", "用户没有理财订单"},
		},
		Input: HoldPositionInput{},
		Expected: HoldPositionExpect{
			Success:    true,
			StatusCode: 200,
			EmptyList:  true,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_POS_003",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "未登录应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"未登录"},
		},
		Input: HoldPositionInput{},
		Expected: HoldPositionExpect{
			Success:        false,
			StatusCode:     401,
			ErrMsgContains: "登录",
		},
	},
}

// GetHoldPositionCaseByID 根据ID获取持仓查询用例
func GetHoldPositionCaseByID(caseID string) *HoldPositionCase {
	for i := range HoldPositionCases {
		if HoldPositionCases[i].CaseID == caseID {
			return &HoldPositionCases[i]
		}
	}
	return nil
}
