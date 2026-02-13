package order

// TestCase 订单查询测试用例基类
type TestCase struct {
	CaseID       string
	Module       string
	Priority     string
	Type         string
	Title        string
	Tags         []string
	PreCondition []string
}

// PageCase 分页查询测试用例
type PageCase struct {
	TestCase
	Input    PageInput
	Expected PageExpect
}

// PageInput 分页查询输入
type PageInput struct {
	PageIndex      int
	PageSize       int
	Coin           string
	Status         *int
	SpecValue      *int
	HasOrder       bool
	OrderCount     int
	UseNotExistCoin bool
}

// PageExpect 分页查询预期结果
type PageExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
	EmptyList      bool
	MinOrderCount  int
}

// DetailCase 订单详情测试用例
type DetailCase struct {
	TestCase
	Input    DetailInput
	Expected DetailExpect
}

// DetailInput 订单详情查询输入
type DetailInput struct {
	OrderID        uint
	HasOrder       bool
	NotExistOrder  bool
	OtherUserOrder bool
}

// DetailExpect 订单详情查询预期
type DetailExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// HisCase 历史记录测试用例
type HisCase struct {
	TestCase
	Input    HisInput
	Expected HisExpect
}

// HisInput 历史记录查询输入
type HisInput struct {
	Coin     string
	HasOrder bool
}

// HisExpect 历史记录查询预期
type HisExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
	EmptyList      bool
}

// HoldPositionCase 持仓查询测试用例
type HoldPositionCase struct {
	TestCase
	Input    HoldPositionInput
	Expected HoldPositionExpect
}

// HoldPositionInput 持仓查询输入
type HoldPositionInput struct {
	Coin     string
	HasOrder bool
}

// HoldPositionExpect 持仓查询预期
type HoldPositionExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
	EmptyList      bool
}

// InterestPageCase 收益明细测试用例
type InterestPageCase struct {
	TestCase
	Input    InterestPageInput
	Expected InterestPageExpect
}

// InterestPageInput 收益明细查询输入
type InterestPageInput struct {
	OrderID  uint
	HasOrder bool
}

// InterestPageExpect 收益明细查询预期
type InterestPageExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
	EmptyList      bool
}

// PeriodDetailCase 期间详情测试用例
type PeriodDetailCase struct {
	TestCase
	Input    PeriodDetailInput
	Expected PeriodDetailExpect
}

// PeriodDetailInput 期间详情查询输入
type PeriodDetailInput struct {
	OrderID        uint
	HasOrder       bool
	IsFixed        bool
	NotExistOrder  bool
	OtherUserOrder bool
	EmptyOrderID   bool
}

// PeriodDetailExpect 期间详情查询预期
type PeriodDetailExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// OpenSubCase 自动申购设置测试用例
type OpenSubCase struct {
	TestCase
	Input    OpenSubInput
	Expected OpenSubExpect
}

// OpenSubInput 自动申购设置输入
type OpenSubInput struct {
	OrderID        uint
	OpenSub        int8
	HasOrder       bool
	InitialOpenSub int8
	NotExistOrder  bool
	OtherUserOrder bool
	EmptyOrderID   bool
	EmptyOpenSub   bool
	InvalidOpenSub bool
}

// OpenSubExpect 自动申购设置预期
type OpenSubExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// intPtr 辅助函数
func intPtr(i int) *int {
	return &i
}
