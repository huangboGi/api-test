package product

// PageCases 产品分页查询测试用例
var PageCases = []struct {
	TestCase
	Input  QueryInput
	Expect QueryExpect
}{
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_PAGE_FUNC_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询产品列表-成功场景",
		},
		Input: QueryInput{
			PageIndex: 1,
			PageSize:  10,
		},
		Expect: QueryExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_PAGE_NEG_001",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "查询产品列表-pageIndex为空应失败",
		},
		Input: QueryInput{
			PageIndex: 0,
			PageSize:  10,
		},
		Expect: QueryExpect{
			Success:        false,
			ErrMsgContains: "不能为空",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_PAGE_NEG_002",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "查询产品列表-pageIndex为0应失败",
		},
		Input: QueryInput{
			PageIndex: 0,
			PageSize:  10,
		},
		Expect: QueryExpect{
			Success:        false,
			ErrMsgContains: "大于0",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_PAGE_NEG_003",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "查询产品列表-pageIndex为负数应失败",
		},
		Input: QueryInput{
			PageIndex: -1,
			PageSize:  10,
		},
		Expect: QueryExpect{
			Success:        false,
			ErrMsgContains: "大于0",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_PAGE_BND_001",
			Module:   "产品管理",
			Priority: "中",
			Type:     "边界测试",
			Title:    "查询产品列表-pageSize边界值测试（最大值）",
		},
		Input: QueryInput{
			PageIndex: 1,
			PageSize:  100,
		},
		Expect: QueryExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_PAGE_NEG_004",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "查询产品列表-pageSize为负数应失败",
		},
		Input: QueryInput{
			PageIndex: 1,
			PageSize:  -1,
		},
		Expect: QueryExpect{
			Success:        false,
			ErrMsgContains: "大于0",
		},
	},
}

// ListCases 产品列表查询测试用例（用户端）
var ListCases = []struct {
	TestCase
	Input  QueryInput
	Expect QueryExpect
}{
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_LIST_FUNC_001",
			Module:   "产品管理（用户端）",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询理财产品列表-成功场景",
		},
		Input: QueryInput{},
		Expect: QueryExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_LIST_FUNC_002",
			Module:   "产品管理（用户端）",
			Priority: "中",
			Type:     "功能测试",
			Title:    "查询理财产品列表-按名称过滤",
		},
		Input: QueryInput{
			Name: "USDT",
		},
		Expect: QueryExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_LIST_FUNC_004",
			Module:   "产品管理（用户端）",
			Priority: "中",
			Type:     "功能测试",
			Title:    "查询理财产品列表-按规格值过滤",
		},
		Input: QueryInput{
			SpecValue:    30,
			DeadlineType: 1,
		},
		Expect: QueryExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_LIST_FUNC_006",
			Module:   "产品管理（用户端）",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询理财产品列表-按活期类型过滤",
		},
		Input: QueryInput{
			DeadlineType: 0,
		},
		Expect: QueryExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_LIST_FUNC_003",
			Module:   "产品管理（用户端）",
			Priority: "中",
			Type:     "功能测试",
			Title:    "查询理财产品列表-数据为空场景",
		},
		Input: QueryInput{
			EmptyData: true,
		},
		Expect: QueryExpect{Success: true, StatusCode: 200},
	},
}
