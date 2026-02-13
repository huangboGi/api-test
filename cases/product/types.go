package product

import (
	"github.com/shopspring/decimal"
)

// TestCase 产品测试用例基类
type TestCase struct {
	CaseID       string
	Module       string
	Priority     string
	Type         string
	Title        string
	Tags         []string
	PreCondition []string
}

// AddInput 添加产品输入
type AddInput struct {
	Coin            string          // 币种，空表示自动创建
	CoinKey         string          // 币种Key
	SpecValue       interface{}     // 可以是 int 或其他类型
	DeadlineType    interface{}     // 可以是 int 或无效值
	AnnualAte       decimal.Decimal // 年化收益率
	MinVol          decimal.Decimal // 最小金额
	UseQuotaTotal   decimal.Decimal // 总使用额度
	PersonQuota     decimal.Decimal // 个人额度
	ExtraAnnualAte  decimal.Decimal // 附加年化收益率
	DailyMaximum    decimal.Decimal // 每日最大限额
	Tag             string          // 产品名称
	SpecOff         bool            // 规格是否下架
	CoinOff         bool            // 币种是否下架
	SpecNotExist    bool            // 规格不存在
	CoinNotExist    bool            // 币种不存在
}

// AddExpect 添加产品预期结果
type AddExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// UpdateInput 修改产品输入
type UpdateInput struct {
	ID              uint            // 产品ID，0表示自动创建
	Coin            string          // 币种
	CoinKey         string          // 币种Key
	SpecValue       interface{}     // 规格值
	DeadlineType    interface{}     // 期限类型
	AnnualAte       decimal.Decimal // 年化收益率
	MinVol          decimal.Decimal // 最小金额
	UseQuotaTotal   decimal.Decimal // 总使用额度
	PersonQuota     decimal.Decimal // 个人额度
	ExtraAnnualAte  decimal.Decimal // 附加年化收益率
	DailyMaximum    decimal.Decimal // 每日最大限额
	NotExist        bool            // 产品不存在
	SpecNotExist    bool            // 规格不存在
	CoinNotExist    bool            // 币种不存在
	ChangeCoin      bool            // 尝试修改币种
}

// UpdateExpect 修改产品预期结果
type UpdateExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// ShelvesInput 上下架输入
type ShelvesInput struct {
	ID            uint // 产品ID，0表示自动创建
	ShelvesStatus int  // 上下架状态
	NotExist      bool // 产品不存在
	SpecOff       bool // 规格已下架
	CoinOff       bool // 币种已下架
	SpecNotExist  bool // 规格已删除
	CoinNotExist  bool // 币种已删除
}

// ShelvesExpect 上下架预期结果
type ShelvesExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}

// decimalPtr 辅助函数 - 返回 decimal.Decimal
func decimalPtr(v float64) decimal.Decimal {
	return decimal.NewFromFloat(v)
}

// QueryInput 产品查询输入
type QueryInput struct {
	PageIndex    int    // 页码
	PageSize     int    // 每页数量
	Name         string // 产品名称过滤
	SpecValue    int    // 规格值过滤
	DeadlineType int    // 期限类型过滤
	EmptyData    bool   // 是否测试空数据
}

// QueryExpect 产品查询预期结果
type QueryExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string
}
