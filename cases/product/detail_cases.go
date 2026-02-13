package product

// DetailInput 产品详情输入
type DetailInput struct {
	ID         uint // 产品ID
	NotExist   bool // 产品不存在
	NotOnShelf bool // 产品未上架（用户端测试）
}

// DetailExpect 产品详情预期结果
type DetailExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// AdminDetailCases 管理端产品详情测试用例
var AdminDetailCases = []struct {
	TestCase
	Input  DetailInput
	Expect DetailExpect
}{
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_DETAIL_FUNC_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "功能测试",
			Title:    "获取产品详情-成功场景",
			PreCondition: []string{
				"已登录管理端系统",
				"数据库中已存在产品数据",
			},
		},
		Input:  DetailInput{},
		Expect: DetailExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_DETAIL_VAL_001",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "获取产品详情-ID为空应失败",
		},
		Input: DetailInput{
			ID: 0,
		},
		Expect: DetailExpect{
			Success:        false,
			ErrMsgContains: "ID",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_DETAIL_VAL_002",
			Module:   "产品管理",
			Priority: "高",
			Type:     "验证测试",
			Title:    "获取产品详情-产品不存在应失败",
		},
		Input: DetailInput{
			NotExist: true,
			ID:       999999,
		},
		Expect: DetailExpect{
			Success:        false,
			ErrMsgContains: "不存在",
		},
	},
}

// AppDetailCases 用户端产品详情测试用例
var AppDetailCases = []struct {
	TestCase
	Input  DetailInput
	Expect DetailExpect
}{
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_DETAIL_FUNC_001",
			Module:   "产品管理（用户端）",
			Priority: "高",
			Type:     "功能测试",
			Title:    "获取产品详情-成功场景",
			PreCondition: []string{
				"已登录用户端",
				"产品已上架",
			},
		},
		Input:  DetailInput{},
		Expect: DetailExpect{Success: true, StatusCode: 200},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_DETAIL_VAL_001",
			Module:   "产品管理（用户端）",
			Priority: "高",
			Type:     "验证测试",
			Title:    "获取产品详情-ID为空应失败",
		},
		Input: DetailInput{
			ID: 0,
		},
		Expect: DetailExpect{
			Success:        false,
			ErrMsgContains: "ID",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_DETAIL_VAL_002",
			Module:   "产品管理（用户端）",
			Priority: "高",
			Type:     "验证测试",
			Title:    "获取产品详情-产品不存在应失败",
		},
		Input: DetailInput{
			NotExist: true,
			ID:       999999,
		},
		Expect: DetailExpect{
			Success:        false,
			ErrMsgContains: "不存在",
		},
	},
	{
		TestCase: TestCase{
			CaseID:   "WTH_PRODUCT_APP_DETAIL_VAL_003",
			Module:   "产品管理（用户端）",
			Priority: "高",
			Type:     "验证测试",
			Title:    "获取产品详情-产品未上架应失败",
		},
		Input: DetailInput{
			NotOnShelf: true,
		},
		Expect: DetailExpect{
			Success:        false,
			ErrMsgContains: "不存在",
		},
	},
}
