package coin

// TestCase 币种配置测试用例基类
type TestCase struct {
	CaseID       string
	Module       string
	Priority     string
	Type         string
	Title        string
	Tags         []string
	PreCondition []string
}

// LangItem 多语言项
type LangItem struct {
	LangKey string `json:"langKey"`
	Content string `json:"content"`
}

// CoinInput 币种配置输入
type CoinInput struct {
	// 操作类型
	Action string // "add" | "update" | "shelves" | "page" | "detail" | "selectCoin"

	// 基础字段
	Coin         string     // 币种代码（必填）
	CoinKey      string     // 币种Key
	Tag          string     // 标签
	LangNameList []LangItem // 多语言名称列表

	// 上下架参数
	Shelves int // 0-下架 1-上架

	// 更新/查询参数
	ID int64 // 更新/上下架/详情时使用

	// 分页查询参数
	PageIndex int    // 页码
	PageSize  int    // 每页条数
	CoinName  string // 币种名称（模糊查询）
	Lang      string // 语言

	// 特殊场景
	NotExist      bool // 币种不存在
	NoAuth        bool // 无权限
	EmptyResult   bool // 预期空结果
	CoinDuplicate bool // Coin重复测试
}

// CoinExpect 币种配置预期结果
type CoinExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains string

	// 数据库验证
	DBCheck DBCheck

	// 响应验证
	DataCheck DataCheck
}

// DBCheck 数据库验证点
type DBCheck struct {
	RecordCreated  bool
	RecordUpdated  bool
	ShelvesMatch   int    // 预期的上下架状态
	TagMatch       string // 预期的标签
	CoinKeyMatch   string // 预期的CoinKey

	// 字段值验证
	CoinMatch string // 预期的Coin值

	// 多语言数据验证
	LangDataCheck LangDataCheck // 多语言数据验证
}

// LangDataCheck 多语言数据验证点
type LangDataCheck struct {
	ShouldExist   bool     // 是否应该存在多语言数据
	ExpectedCount int      // 预期的语言数量
	ExpectedItems []LangItem // 预期的语言项（用于验证内容）
}

// DataCheck 响应数据验证点
type DataCheck struct {
	TotalCount    int64  // 预期总数
	HasID         bool   // 返回数据包含ID
	HasCoinName   bool   // 返回数据包含币种名称
	HasShelves    bool   // 返回数据包含上下架状态
	IsEmptyResult bool   // 是否为空结果
}
