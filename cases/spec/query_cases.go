package spec

// QueryCases 规格查询测试用例
var QueryCases = []struct {
	TestCase
	Input  SpecInput
	Expect SpecExpect
}{
	// ==================== Page 分页查询 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_PAGE_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "分页查询-成功返回列表",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在规格数据",
			},
		},
		Input: SpecInput{
			Action:    "page",
			PageIndex: 1,
			PageSize:  10,
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				HasID:       true,
				HasSpecName: true,
				HasShelves:  true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_PAGE_002",
			Module:   "规格管理",
			Priority: "中",
			Type:     "功能测试",
			Title:    "分页查询-按规格值筛选",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在规格数据",
			},
		},
		Input: SpecInput{
			Action:    "page",
			PageIndex: 1,
			PageSize:  10,
			SpecValue: 30,
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_PAGE_003",
			Module:   "规格管理",
			Priority: "中",
			Type:     "功能测试",
			Title:    "分页查询-按期限类型筛选",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在规格数据",
			},
		},
		Input: SpecInput{
			Action:       "page",
			PageIndex:    1,
			PageSize:     10,
			DeadlineType: 1,
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_PAGE_004",
			Module:   "规格管理",
			Priority: "低",
			Type:     "功能测试",
			Title:    "分页查询-空数据返回空列表",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: SpecInput{
			Action:      "page",
			PageIndex:   1,
			PageSize:    10,
			SpecValue:   999999,
			EmptyResult: true,
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				IsEmptyResult: true,
			},
		},
	},

	// ==================== List 列表查询 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_LIST_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "列表查询-成功返回所有规格",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在规格数据",
			},
		},
		Input: SpecInput{
			Action: "list",
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				HasID:       true,
				HasSpecName: true,
			},
		},
	},

	// ==================== Detail 详情查询 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_DETAIL_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询详情-成功返回",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在规格数据",
			},
		},
		Input: SpecInput{
			Action: "detail",
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				HasID:       true,
				HasSpecName: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_DETAIL_NEG_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "查询详情-ID为空应失败",
			Tags:     []string{"P0"},
		},
		Input: SpecInput{
			Action: "detail",
			ID:     0,
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_DETAIL_NEG_002",
			Module:   "规格管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "查询详情-规格不存在应失败",
			Tags:     []string{"P0"},
		},
		Input: SpecInput{
			Action:   "detail",
			ID:       999999,
			NotExist: true,
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},

	// ==================== 安全测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_QUERY_SEC_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "安全测试",
			Title:    "分页查询-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: SpecInput{
			Action: "page",
			NoAuth: true,
		},
		Expect: SpecExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_QUERY_SEC_002",
			Module:   "规格管理",
			Priority: "高",
			Type:     "安全测试",
			Title:    "列表查询-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: SpecInput{
			Action: "list",
			NoAuth: true,
		},
		Expect: SpecExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_QUERY_SEC_003",
			Module:   "规格管理",
			Priority: "高",
			Type:     "安全测试",
			Title:    "查询详情-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: SpecInput{
			Action: "detail",
			ID:     1,
			NoAuth: true,
		},
		Expect: SpecExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
}
