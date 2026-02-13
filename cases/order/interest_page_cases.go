package order

// InterestPageCases 收益明细测试用例表
var InterestPageCases = []InterestPageCase{
	{
		TestCase: TestCase{
			CaseID:   "ORDER_INT_001",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询收益明细成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{"已登录系统", "用户有收益明细"},
		},
		Input: InterestPageInput{
			HasOrder: true,
		},
		Expected: InterestPageExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_INT_002",
			Module:   "用户订单查询",
			Priority: "低",
			Type:     "功能测试",
			Title:    "查询结果为空",
			Tags:     []string{"P2"},
			PreCondition: []string{"已登录系统", "用户没有收益明细"},
		},
		Input: InterestPageInput{},
		Expected: InterestPageExpect{
			Success:    true,
			StatusCode: 200,
			EmptyList:  true,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_INT_003",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "未登录应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"未登录"},
		},
		Input: InterestPageInput{},
		Expected: InterestPageExpect{
			Success:        false,
			StatusCode:     401,
			ErrMsgContains: "登录",
		},
	},
}

// GetInterestPageCaseByID 根据ID获取收益明细用例
func GetInterestPageCaseByID(caseID string) *InterestPageCase {
	for i := range InterestPageCases {
		if InterestPageCases[i].CaseID == caseID {
			return &InterestPageCases[i]
		}
	}
	return nil
}
