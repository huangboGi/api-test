package coin

// UpdateCases 更新币种测试用例
var UpdateCases = []struct {
	TestCase
	Input  CoinInput
	Expect CoinExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_UPD_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "功能测试",
			Title:    "更新币种配置-修改标签成功",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在币种配置",
			},
		},
		Input: CoinInput{
			Action: "update",
			Coin:   "USDT",
			Tag:    "更新后的标签",
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordUpdated: true,
				TagMatch:      "更新后的标签",
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_UPD_002",
			Module:   "币种配置",
			Priority: "中",
			Type:     "功能测试",
			Title:    "更新币种配置-更新多语言名称",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在币种配置",
			},
		},
		Input: CoinInput{
			Action: "update",
			Coin:   "USDT",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "更新后的泰达币"},
				{LangKey: "en", Content: "Updated Tether"},
			},
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_UPD_003",
			Module:   "币种配置",
			Priority: "中",
			Type:     "功能测试",
			Title:    "更新币种配置-更新CoinKey",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在币种配置",
			},
		},
		Input: CoinInput{
			Action:  "update",
			Coin:    "USDT",
			CoinKey: "USDT_NEW_KEY",
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				CoinKeyMatch: "USDT_NEW_KEY",
			},
		},
	},

	// ==================== 验证测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_UPD_NEG_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "验证测试",
			Title:    "更新币种-ID为空应失败",
			Tags:     []string{"P0"},
		},
		Input: CoinInput{
			Action: "update",
			ID:     0, // ID为空
			Coin:   "USDT",
		},
		Expect: CoinExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_UPD_NEG_002",
			Module:   "币种配置",
			Priority: "高",
			Type:     "验证测试",
			Title:    "更新币种-币种不存在应失败",
			Tags:     []string{"P0"},
		},
		Input: CoinInput{
			Action:   "update",
			ID:       999999, // 不存在的ID
			Coin:     "USDT",
			NotExist: true,
		},
		Expect: CoinExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_UPD_NEG_003",
			Module:   "币种配置",
			Priority: "高",
			Type:     "验证测试",
			Title:    "更新币种-币种代码为空应失败",
			Tags:     []string{"P0"},
		},
		Input: CoinInput{
			Action: "update",
			ID:     1,
			Coin:   "", // 空币种代码
		},
		Expect: CoinExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},

	// ==================== 安全测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_UPD_SEC_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "安全测试",
			Title:    "更新币种-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: CoinInput{
			Action: "update",
			ID:     1,
			Coin:   "USDT",
			NoAuth: true,
		},
		Expect: CoinExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
}
