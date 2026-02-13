package interest

// FuncCases 利息功能测试用例
var FuncCases = []struct {
	TestCase
	Input  InterestInput
	Expect InterestExpect
}{
	{
		TestCase: TestCase{
			CaseID:   "WTH_INTEREST_FUNC_001",
			Module:   "利息管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "利息发放-成功场景",
			PreCondition: []string{
				"已登录管理端系统",
				"有利息发放权限",
			},
		},
		Input: InterestInput{},
		Expect: InterestExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
}
