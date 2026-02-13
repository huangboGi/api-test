package coin

// AddCases 添加币种测试用例
var AddCases = []struct {
	TestCase
	Input  CoinInput
	Expect CoinExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_ADD_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "功能测试",
			Title:    "添加币种-成功场景",
			Tags:     []string{"P0", "smoke"},
			PreCondition: []string{
				"已登录管理端系统",
				"有添加币种权限",
			},
		},
		Input: CoinInput{
			Action: "add",
			Coin:   "USDT",
			Tag:    "热门",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "泰达币"},
				{LangKey: "en", Content: "Tether"},
			},
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
				TagMatch:      "热门",
				LangDataCheck: LangDataCheck{
					ShouldExist: true,
				},
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_ADD_002",
			Module:   "币种配置",
			Priority: "中",
			Type:     "功能测试",
			Title:    "添加币种-CoinKey自动使用Coin值",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: CoinInput{
			Action:  "add",
			Coin:    "BTC",
			CoinKey: "", // 不指定CoinKey，应自动使用Coin值
			Tag:     "主流",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "比特币"},
			},
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
				CoinKeyMatch:  "BTC",
				TagMatch:      "主流",
				LangDataCheck: LangDataCheck{
					ShouldExist: true,
				},
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_ADD_003",
			Module:   "币种配置",
			Priority: "中",
			Type:     "功能测试",
			Title:    "添加币种-指定自定义CoinKey",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: CoinInput{
			Action:  "add",
			Coin:    "ETH",
			CoinKey: "ETH_CUSTOM_KEY",
			Tag:     "主流",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "以太坊"},
			},
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
				CoinKeyMatch:  "ETH_CUSTOM_KEY",
				TagMatch:      "主流",
				LangDataCheck: LangDataCheck{
					ShouldExist: true,
				},
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_ADD_004",
			Module:   "币种配置",
			Priority: "中",
			Type:     "功能测试",
			Title:    "添加币种-多语言完整测试",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: CoinInput{
			Action: "add",
			Coin:   "DOGE",
			Tag:    "热门",
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "狗狗币"},
				{LangKey: "zh-Hant", Content: "狗狗幣"},
				{LangKey: "en", Content: "Dogecoin"},
				{LangKey: "ja", Content: "ドージコイン"},
				{LangKey: "ko", Content: "도지코인"},
			},
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
				TagMatch:      "热门",
				LangDataCheck: LangDataCheck{
					ShouldExist:   true,
					ExpectedCount: 5,
				},
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_ADD_005",
			Module:   "币种配置",
			Priority: "低",
			Type:     "功能测试",
			Title:    "添加币种-无标签",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: CoinInput{
			Action: "add",
			Coin:   "XRP",
			Tag:    "", // 无标签
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "瑞波币"},
			},
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DBCheck: DBCheck{
				RecordCreated: true,
				TagMatch:      "",
				LangDataCheck: LangDataCheck{
					ShouldExist: true,
				},
			},
		},
	},

	// ==================== 验证测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_ADD_NEG_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加币种-币种代码为空应失败",
			Tags:     []string{"P0"},
		},
		Input: CoinInput{
			Action: "add",
			Coin:   "", // 空币种代码
			LangNameList: []LangItem{
				{LangKey: "zh-Hans", Content: "测试币"},
			},
		},
		Expect: CoinExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_ADD_NEG_002",
			Module:   "币种配置",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加币种-多语言名称列表为空应失败",
			Tags:     []string{"P0"},
		},
		Input: CoinInput{
			Action:       "add",
			Coin:         "TEST",
			LangNameList: []LangItem{}, // 空多语言列表
		},
		Expect: CoinExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "请至少填写一种语言的币种名称",
		},
	},

	// ==================== 安全测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_ADD_SEC_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "安全测试",
			Title:    "添加币种-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: CoinInput{
			Action: "add",
			Coin:   "USDT",
			NoAuth: true,
		},
		Expect: CoinExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
}
