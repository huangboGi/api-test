package product

// ShelvesCases 产品上下架测试用例（功能+验证）
var ShelvesCases = []struct {
	TestCase
	Input  ShelvesInput
	Expect ShelvesExpect
}{
	// ==================== 功能测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_SHELF_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "更新上下架状态-上架成功",
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在产品数据",
			},
		},
		Input: ShelvesInput{
			ShelvesStatus: 1,
		},
		Expect: ShelvesExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_SHELF_002",
			Module:   "产品管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "更新上下架状态-下架成功",
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在产品数据",
			},
		},
		Input: ShelvesInput{
			ShelvesStatus: 0,
		},
		Expect: ShelvesExpect{
			Success:    true,
			StatusCode: 200,
		},
	},

	// ==================== 验证测试 ====================
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_SHELF_VAL_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "上架时规格已下架应失败",
		},
		Input: ShelvesInput{
			SpecOff:       true,
			ShelvesStatus: 1,
		},
		Expect: ShelvesExpect{
			Success:        false,
			ErrMsgContains: "规格",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_SHELF_VAL_002",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "上架时币种已禁用应失败",
		},
		Input: ShelvesInput{
			CoinOff:       true,
			ShelvesStatus: 1,
		},
		Expect: ShelvesExpect{
			Success:        false,
			ErrMsgContains: "币种",
		},
	},
}
