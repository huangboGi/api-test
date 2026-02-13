package spec

// UpdateCases 更新规格测试用例
var UpdateCases = []struct {
	TestCase
	Input  SpecInput
	Expect SpecExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_UPD_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "修改规格-更新备注成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在规格数据",
			},
		},
		Input: SpecInput{
			Action: "update",
			Remark: "更新后的备注",
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordUpdated: true,
				RemarkMatch:   "更新后的备注",
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_UPD_002",
			Module:   "规格管理",
			Priority: "中",
			Type:     "功能测试",
			Title:    "修改规格-更新多语言名称",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在规格数据",
			},
		},
		Input: SpecInput{
			Action: "update",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "更新后的中文名"},
				{LangKey: "en", Content: "Updated English Name"},
			},
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_UPD_003",
			Module:   "规格管理",
			Priority: "中",
			Type:     "功能测试",
			Title:    "修改规格-更新期限类型",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在规格数据",
			},
		},
		Input: SpecInput{
			Action:       "update",
			DeadlineType: 0, // 改为活期
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				DeadlineTypeMatch: 0,
			},
		},
	},

	// ==================== 验证测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_UPD_NEG_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改规格-规格不存在应失败",
			Tags:     []string{"P0"},
		},
		Input: SpecInput{
			Action:   "update",
			ID:       999999,
			NotExist: true,
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "规格不存在",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_UPD_NEG_002",
			Module:   "规格管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改规格-ID为空应失败",
			Tags:     []string{"P0"},
		},
		Input: SpecInput{
			Action: "update",
			ID:     0, // 空ID
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_UPD_NEG_003",
			Module:   "规格管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "修改规格-DeadlineType非法值应失败",
			Tags:     []string{"P1"},
		},
		Input: SpecInput{
			Action:       "update",
			ID:           1,
			DeadlineType: 2, // 非法值
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "期限类型错误",
		},
	},

	// ==================== 安全测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_UPD_SEC_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "安全测试",
			Title:    "修改规格-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: SpecInput{
			Action: "update",
			ID:     1,
			NoAuth: true,
		},
		Expect: SpecExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
}
