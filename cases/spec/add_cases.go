package spec

// AddCases 添加规格测试用例
var AddCases = []struct {
	TestCase
	Input  SpecInput
	Expect SpecExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "添加规格-定期规格成功",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{
				"已登录管理端系统",
				"有添加规格权限",
			},
		},
		Input: SpecInput{
			Action:        "add",
			SpecValue:     30,
			DeadlineType:  1,
			ShelvesStatus: 0,
			Remark:        "30天定期规格",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "30天"},
				{LangKey: "en", Content: "30 Days"},
			},
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_002",
			Module:   "规格管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "添加规格-活期规格成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: SpecInput{
			Action:        "add",
			SpecValue:     -1,
			DeadlineType:  0,
			ShelvesStatus: 0,
			Remark:        "活期规格",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "活期"},
				{LangKey: "en", Content: "Flexible"},
			},
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_003",
			Module:   "规格管理",
			Priority: "中",
			Type:     "功能测试",
			Title:    "添加规格-SpecKey自动生成",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: SpecInput{
			Action:       "add",
			SpecValue:    60,
			DeadlineType: 1,
			SpecKey:      "", // 不指定SpecKey，应自动生成 "60_1"
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "60天"},
			},
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
				SpecKeyMatch:  "60_1",
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_004",
			Module:   "规格管理",
			Priority: "中",
			Type:     "功能测试",
			Title:    "添加规格-指定自定义SpecKey",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: SpecInput{
			Action:       "add",
			SpecValue:    90,
			DeadlineType: 1,
			SpecKey:      "CUSTOM_SPEC_KEY",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "90天"},
			},
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
				SpecKeyMatch:  "CUSTOM_SPEC_KEY",
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_005",
			Module:   "规格管理",
			Priority: "中",
			Type:     "功能测试",
			Title:    "添加规格-多语言完整测试",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: SpecInput{
			Action:       "add",
			SpecValue:    180,
			DeadlineType: 1,
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "180天"},
				{LangKey: "zh-Hant", Content: "180天"},
				{LangKey: "en", Content: "180 Days"},
				{LangKey: "ja", Content: "180日"},
				{LangKey: "ko", Content: "180일"},
			},
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_006",
			Module:   "规格管理",
			Priority: "低",
			Type:     "功能测试",
			Title:    "添加规格-直接上架",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: SpecInput{
			Action:        "add",
			SpecValue:     365,
			DeadlineType:  1,
			ShelvesStatus: 1, // 直接上架
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "365天"},
			},
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
				ShelvesMatch:  1,
			},
		},
	},

	// ==================== 验证测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_NEG_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加规格-SpecValue重复应失败",
			Tags:     []string{"P0"},
		},
		Input: SpecInput{
			Action:       "add",
			SpecValue:    30, // 已存在的规格值
			DeadlineType: 1,
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "测试"},
			},
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "已存在",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_NEG_002",
			Module:   "规格管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加规格-SpecValue为空应失败",
			Tags:     []string{"P0"},
		},
		Input: SpecInput{
			Action:       "add",
			SpecValue:    nil, // 空规格值
			DeadlineType: 1,
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "测试"},
			},
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_NEG_003",
			Module:   "规格管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加规格-LangNameList为空应失败",
			Tags:     []string{"P0"},
		},
		Input: SpecInput{
			Action:        "add",
			SpecValue:     99,
			DeadlineType:  1,
			LangNameList:  []LangItem{}, // 空多语言列表
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "请至少填写一种语言的规格名称",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_ADD_NEG_004",
			Module:   "规格管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "添加规格-DeadlineType非法值应失败",
			Tags:     []string{"P1"},
		},
		Input: SpecInput{
			Action:       "add",
			SpecValue:    31,
			DeadlineType: 2, // 非法值
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "测试"},
			},
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
			CaseID:   "WTH_SPEC_ADD_SEC_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "安全测试",
			Title:    "添加规格-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: SpecInput{
			Action:    "add",
			SpecValue: 100,
			NoAuth:    true,
		},
		Expect: SpecExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
}
