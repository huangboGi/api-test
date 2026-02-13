package product

// UpdateCases 修改产品测试用例（功能+验证）
var UpdateCases = []struct {
	TestCase
	Input  UpdateInput
	Expect UpdateExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_FUNC_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "修改产品-成功场景",
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在产品数据",
			},
		},
		Input: UpdateInput{
			SpecValue:    -1,
			DeadlineType: 0,
		},
		Expect: UpdateExpect{
			Success:    true,
			StatusCode: 200,
		},
	},

	// ==================== 验证测试 - 活期改定期 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改产品-活期改定期-specValue>0应成功",
		},
		Input: UpdateInput{
			SpecValue:    30,
			DeadlineType: 1,
		},
		Expect: UpdateExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_002",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改产品-活期改定期-specValue<=0应失败",
		},
		Input: UpdateInput{
			SpecValue:    0,
			DeadlineType: 1,
		},
		Expect: UpdateExpect{
			Success:        false,
			ErrMsgContains: "定期产品的规则值必须大于0",
		},
	},

	// ==================== 验证测试 - 定期改活期 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_003",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改产品-定期改活期-specValue=-1应成功",
		},
		Input: UpdateInput{
			SpecValue:    -1,
			DeadlineType: 0,
		},
		Expect: UpdateExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_004",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改产品-定期改活期-specValue!=-1应失败",
		},
		Input: UpdateInput{
			SpecValue:    7,
			DeadlineType: 0,
		},
		Expect: UpdateExpect{
			Success:        false,
			ErrMsgContains: "活期产品的规则值必须为-1",
		},
	},

	// ==================== 验证测试 - 修改币种 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_005",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改产品-尝试修改coin字段应失败",
		},
		Input: UpdateInput{
			ChangeCoin: true,
		},
		Expect: UpdateExpect{
			Success:        false,
			ErrMsgContains: "不允许修改",
		},
	},

	// ==================== 验证测试 - 不存在验证 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_006",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改产品-规格不存在应失败",
		},
		Input: UpdateInput{
			SpecNotExist: true,
			SpecValue:    9999,
		},
		Expect: UpdateExpect{
			Success:        false,
			ErrMsgContains: "规格",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_007",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "修改产品-币种不存在应失败",
		},
		Input: UpdateInput{
			CoinNotExist: true,
			Coin:         "NONEXISTCOIN",
		},
		Expect: UpdateExpect{
			Success:        false,
			ErrMsgContains: "币种",
		},
	},

	// ==================== 验证测试 - 数值验证 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_008",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "修改产品-年化收益率为负数应失败",
		},
		Input: UpdateInput{
			AnnualAte: decimalPtr(-5),
		},
		Expect: UpdateExpect{
			Success:        false,
			ErrMsgContains: "年化收益率",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_UPD_VAL_009",
			Module:   "产品管理",
			Priority: "中",
			Type:     "验证测试",
			Title:    "修改产品-最小金额为负数应失败",
		},
		Input: UpdateInput{
			MinVol: decimalPtr(-100),
		},
		Expect: UpdateExpect{
			Success:        false,
			ErrMsgContains: "最小金额",
		},
	},
}
