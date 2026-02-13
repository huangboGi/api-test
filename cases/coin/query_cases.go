package coin

// QueryCases 币种查询测试用例
var QueryCases = []struct {
	TestCase
	Input  CoinInput
	Expect CoinExpect
}{
	// ==================== Page 分页查询 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_PAGE_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "功能测试",
			Title:    "分页查询-成功返回列表",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在币种配置数据",
			},
		},
		Input: CoinInput{
			Action:    "page",
			PageIndex: 1,
			PageSize:  10,
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				HasID:       true,
				HasCoinName: true,
				HasShelves:  true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_PAGE_002",
			Module:   "币种配置",
			Priority: "中",
			Type:     "功能测试",
			Title:    "分页查询-按币种标识模糊查询",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在包含USDT的币种配置",
			},
		},
		Input: CoinInput{
			Action:    "page",
			PageIndex: 1,
			PageSize:  10,
			Coin:      "USD",
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_PAGE_003",
			Module:   "币种配置",
			Priority: "中",
			Type:     "功能测试",
			Title:    "分页查询-按币种名称模糊查询",
			Tags:     []string{"P1"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在多语言币种配置",
			},
		},
		Input: CoinInput{
			Action:    "page",
			PageIndex: 1,
			PageSize:  10,
			CoinName:  "泰达",
			Lang:      "zh-Hans",
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_PAGE_004",
			Module:   "币种配置",
			Priority: "低",
			Type:     "功能测试",
			Title:    "分页查询-空数据返回空列表",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: CoinInput{
			Action:      "page",
			PageIndex:   1,
			PageSize:    10,
			Coin:        "NOT_EXIST_COIN_XXX",
			EmptyResult: true,
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				IsEmptyResult: true,
			},
		},
	},

	// ==================== Detail 详情查询 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_DETAIL_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询详情-成功返回",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在币种配置",
			},
		},
		Input: CoinInput{
			Action: "detail",
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				HasID:       true,
				HasCoinName: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_DETAIL_NEG_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "验证测试",
			Title:    "查询详情-ID为空应失败",
			Tags:     []string{"P0"},
		},
		Input: CoinInput{
			Action: "detail",
			ID:     0,
		},
		Expect: CoinExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_DETAIL_NEG_002",
			Module:   "币种配置",
			Priority: "高",
			Type:     "验证测试",
			Title:    "查询详情-币种不存在应失败",
			Tags:     []string{"P0"},
		},
		Input: CoinInput{
			Action:   "detail",
			ID:       999999,
			NotExist: true,
		},
		Expect: CoinExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: "参数错误",
		},
	},

	// ==================== SelectCoin 查询可用币种 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_SELECT_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "功能测试",
			Title:    "查询可用币种-返回上架币种列表",
			Tags:     []string{"P0"},
			PreCondition: []string{
				"已登录管理端系统",
				"已存在上架状态的币种配置",
			},
		},
		Input: CoinInput{
			Action: "selectCoin",
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				HasCoinName: true,
			},
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_SELECT_002",
			Module:   "币种配置",
			Priority: "低",
			Type:     "功能测试",
			Title:    "查询可用币种-无上架币种返回空列表",
			Tags:     []string{"P2"},
			PreCondition: []string{
				"已登录管理端系统",
				"无上架状态的币种配置",
			},
		},
		Input: CoinInput{
			Action:      "selectCoin",
			EmptyResult: true,
		},
		Expect: CoinExpect{
			Success:    true,
			StatusCode: 200,
			DataCheck: DataCheck{
				IsEmptyResult: true,
			},
		},
	},

	// ==================== 安全测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_QUERY_SEC_001",
			Module:   "币种配置",
			Priority: "高",
			Type:     "安全测试",
			Title:    "分页查询-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: CoinInput{
			Action: "page",
			NoAuth: true,
		},
		Expect: CoinExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_QUERY_SEC_002",
			Module:   "币种配置",
			Priority: "高",
			Type:     "安全测试",
			Title:    "查询详情-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: CoinInput{
			Action: "detail",
			ID:     1,
			NoAuth: true,
		},
		Expect: CoinExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_COIN_QUERY_SEC_003",
			Module:   "币种配置",
			Priority: "高",
			Type:     "安全测试",
			Title:    "查询可用币种-未授权访问应失败",
			Tags:     []string{"P0", "auth"},
		},
		Input: CoinInput{
			Action: "selectCoin",
			NoAuth: true,
		},
		Expect: CoinExpect{
			Success:    false,
			StatusCode: 401,
		},
	},
}
