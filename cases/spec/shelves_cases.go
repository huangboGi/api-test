package spec

// ShelvesCases 规格上下架测试用例
var ShelvesCases = []struct {
	TestCase
	Input  SpecInput
	Expect SpecExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_SHELF_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "更新上下架状态-上架成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在下架状态的规格数据",
			},
		},
		Input: SpecInput{
			Action:        "shelves",
			ShelvesStatus: 1,
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				ShelvesMatch: 1,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_SHELF_002",
			Module:   "规格管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "更新上下架状态-下架成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在上架状态的规格数据",
			},
		},
		Input: SpecInput{
			Action:        "shelves",
			ShelvesStatus: 0,
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				ShelvesMatch: 0,
			},
		},
	},

	// ==================== 边界测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_SHELF_BND_001",
			Module:   "规格管理",
			Priority: "中",
			Type:     "边界测试",
			Title:    "上下架切换-状态值下边界(0)",
			Tags:     []string{"P1"},
		},
		Input: SpecInput{
			Action:        "shelves",
			ShelvesStatus: 0,
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				ShelvesMatch: 0,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_SHELF_BND_002",
			Module:   "规格管理",
			Priority: "中",
			Type:     "边界测试",
			Title:    "上下架切换-状态值上边界(1)",
			Tags:     []string{"P1"},
		},
		Input: SpecInput{
			Action:        "shelves",
			ShelvesStatus: 1,
		},
		Expect: SpecExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				ShelvesMatch: 1,
			},
		},
	},

	// ==================== 验证测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_SHELF_NEG_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "更新上下架状态-规格不存在应失败",
			Tags:     []string{"P0"},
		},
		Input: SpecInput{
			Action:        "shelves",
			ID:            999999,
			NotExist:      true,
			ShelvesStatus: 1,
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "规格不存在",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_SHELF_NEG_002",
			Module:   "规格管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "更新上下架状态-ShelvesStatus非法值应失败",
			Tags:     []string{"P1"},
		},
		Input: SpecInput{
			Action:        "shelves",
			ID:            1,
			ShelvesStatus: 2, // 非法值
		},
		Expect: SpecExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "上下架状态错误",
		},
	},

	// ==================== 安全测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_SPEC_SHELF_SEC_001",
			Module:   "规格管理",
			Priority: "高",
			Type:     "安全测试",
			Title:    "更新上下架-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: SpecInput{
			Action: "shelves",
			ID:     1,
			NoAuth: true,
		},
		Expect: SpecExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
}
