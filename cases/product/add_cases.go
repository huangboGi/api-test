package product

// AddCases 添加产品测试用例（功能+验证）
var AddCases = []struct {
	TestCase
	Input  AddInput
	Expect AddExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_FUNC_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "添加产品-定期产品成功",
			PreCondition: []string{
				"已登录管理端系统",
				"有添加产品权限",
			},
		},
		Input: AddInput{
			SpecValue:    30, // 定期
			DeadlineType: 1,
		},
		Expect: AddExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_FUNC_002",
			Module:   "产品管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "添加产品-活期产品成功",
			PreCondition: []string{
				"已登录管理端系统",
			},
		},
		Input: AddInput{
			SpecValue:    -1, // 活期
			DeadlineType: 0,
		},
		Expect: AddExpect{
			Success:    true,
			StatusCode: 200,
		},
	},

	// ==================== 验证测试 - 活期产品 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加活期产品-specValue为-1应成功",
		},
		Input: AddInput{
			SpecValue:    -1,
			DeadlineType: 0,
		},
		Expect: AddExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_002",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加活期产品-specValue不为-1应失败",
		},
		Input: AddInput{
			SpecValue:    7,
			DeadlineType: 0,
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "活期产品的规则值必须为-1",
		},
	},

	// ==================== 验证测试 - 定期产品 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_003",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加定期产品-specValue>0应成功",
		},
		Input: AddInput{
			SpecValue:    30,
			DeadlineType: 1,
		},
		Expect: AddExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_004",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加定期产品-specValue=0应失败",
		},
		Input: AddInput{
			SpecValue:    0,
			DeadlineType: 1,
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "定期产品的规则值必须大于0",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_005",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加定期产品-specValue<0应失败",
		},
		Input: AddInput{
			SpecValue:    -1,
			DeadlineType: 1,
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "定期产品的规则值必须大于0",
		},
	},

	// ==================== 验证测试 - 规格验证 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_006",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加产品-规格不存在应失败",
		},
		Input: AddInput{
			SpecNotExist: true,
			SpecValue:    9999,
			DeadlineType: 1,
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "规格",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_007",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加产品-规格已下架应失败",
		},
		Input: AddInput{
			SpecOff:      true,
			DeadlineType: 1,
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "规格",
		},
	},

	// ==================== 验证测试 - 币种验证 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_008",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加产品-币种不存在应失败",
		},
		Input: AddInput{
			CoinNotExist: true,
			Coin:         "NONEXISTCOIN",
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "币种",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_009",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "添加产品-币种已禁用应失败",
		},
		Input: AddInput{
			CoinOff: true,
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "币种",
		},
	},

	// ==================== 验证测试 - 数值验证 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_010",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "添加产品-年化收益率为负数应失败",
		},
		Input: AddInput{
			AnnualAte: decimalPtr(-5),
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "年化收益率",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_011",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "添加产品-最小金额为负数应失败",
		},
		Input: AddInput{
			MinVol: decimalPtr(-100),
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "最小金额",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_012",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "添加产品-总使用额度为负数应失败",
		},
		Input: AddInput{
			UseQuotaTotal: decimalPtr(-100000),
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "总使用额度",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_013",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "添加产品-个人额度为负数应失败",
		},
		Input: AddInput{
			PersonQuota: decimalPtr(-10000),
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "个人额度",
		},
	},

	// ==================== 验证测试 - 产品名称重复 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_ADD_VAL_016",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "添加产品-同一币种下产品名称重复应失败",
		},
		Input: AddInput{
			Tag: "测试产品名称",
		},
		Expect: AddExpect{
			Success:        false,
			ErrMsgContains: "产品名称",
		},
	},
}
