package interest

import "github.com/shopspring/decimal"

// TestCase 利息测试用例基类
type TestCase struct {
	CaseID       string
	Module       string
	Priority     string
	Type         string
	Title        string
	Tags         []string
	PreCondition []string
}

// InterestInput 利息操作输入
type InterestInput struct {
	Coin         string          // 币种
	Volume       decimal.Decimal // 金额
	AnnualAte    decimal.Decimal // 年化收益率
	ProductOff   bool            // 产品是否下架
	SpecOff      bool            // 规格是否下架
}

// InterestExpect 利息操作预期结果
type InterestExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}
