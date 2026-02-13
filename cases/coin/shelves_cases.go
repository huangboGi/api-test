package coin

// ShelvesCases 币种上下架测试用例
var ShelvesCases = []struct {
	TestCase
	Input  CoinInput
	Expect CoinExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_SHELF_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "功能测试",
			Title:    "更新币种上下架-上架成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在下架状态的币种配置",
			},
		},
		Input: CoinInput{
			Action:  "shelves",
			Shelves: 1,
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				ShelvesMatch: 1,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_SHELF_002",
			Module:   "币种配置",
			Priority: "高",
			Type:     "功能测试",
			Title:    "更新币种上下架-下架成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在上架状态的币种配置",
			},
		},
		Input: CoinInput{
			Action:  "shelves",
			Shelves: 0,
		},
		Expect: CoinExpect{
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
			CaseID:   "WTH_COIN_SHELF_BND_001",
			Module:   "币种配置",
			Priority: "中",
			Type:     "边界测试",
			Title:    "上下架切换-状态值下边界(0)",
			Tags:     []string{"P1"},
		},
		Input: CoinInput{
			Action:  "shelves",
			Shelves: 0,
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				ShelvesMatch: 0,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_SHELF_BND_002",
			Module:   "币种配置",
			Priority: "中",
			Type:     "边界测试",
			Title:    "上下架切换-状态值上边界(1)",
			Tags:     []string{"P1"},
		},
		Input: CoinInput{
			Action:  "shelves",
			Shelves: 1,
		},
		Expect: CoinExpect{
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
			CaseID:   "WTH_COIN_SHELF_NEG_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "验证测试",
			Title:    "更新上下架-币种不存在应失败",
			Tags:     []string{"P0"},
		},
		Input: CoinInput{
			Action:   "shelves",
			ID:       999999,
			NotExist: true,
			Shelves:  1,
		},
		Expect: CoinExpect{
			Success:    false,
			StatusCode: 200,
		},
	},

	// ==================== 安全测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_SHELF_SEC_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "安全测试",
			Title:    "更新上下架-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: CoinInput{
			Action: "shelves",
			ID:     1,
			NoAuth: true,
		},
		Expect: CoinExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
}
