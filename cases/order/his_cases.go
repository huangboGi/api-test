package order

// HisCases 历史记录测试用例表
var HisCases = []HisCase{
	{
		TestCase: TestCase{
			CaseID:   "ORDER_HIS_001",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询历史记录成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{"已登录系统", "用户有历史记录"},
		},
		Input: HisInput{
			HasOrder: true,
		},
		Expected: HisExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_HIS_002",
			Module:   "用户订单查询",
			Priority: "中",
			Type:     "功能测试",
			Title:    "查询结果为空",
			Tags:     []string{"P1"},
			PreCondition: []string{"已登录系统", "用户没有历史记录"},
		},
		Input: HisInput{},
		Expected: HisExpect{
			Success:    true,
			StatusCode: 200,
			EmptyList:  true,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "ORDER_HIS_003",
			Module:   "用户订单查询",
			Priority: "高",
			Type:     "安全测试",
			Title:    "未登录应失败",
			Tags:     []string{"P0"},
			PreCondition: []string{"未登录"},
		},
		Input: HisInput{},
		Expected: HisExpect{
			Success:        false,
			StatusCode:     401,
			ErrMsgContains: "登录",
		},
	},
}

// GetHisCaseByID 根据ID获取历史记录用例
func GetHisCaseByID(caseID string) *HisCase {
	for i := range HisCases {
		if HisCases[i].CaseID == caseID {
			return &HisCases[i]
		}
	}
	return nil
}
