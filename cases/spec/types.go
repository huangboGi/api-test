package spec

// TestCase 规格测试用例基类
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

// SpecInput 规格输入
type SpecInput struct {
	// 操作类型
	Action string // "add" | "update" | "shelves" | "page" | "list" | "detail"

	// 基础字段
	SpecValue     interface{} // 规格值
	SpecKey       string      // 规格Key
	DeadlineType  interface{} // 期限类型 0-活期 1-定期
	ShelvesStatus int         // 上下架状态 0-下架 1-上架
	Remark        string      // 备注
	LangNameList  []LangItem  // 多语言名称列表

	// 更新/查询参数
	ID int64 // 更新/上下架/详情时使用

	// 分页查询参数
	PageIndex int    // 页码
	PageSize  int    // 每页条数
	SpecName  string // 规格名称（模糊查询）
	Lang      string // 语言

	// 特殊场景
	NotExist    bool // 规格不存在
	NoAuth      bool // 无权限
	EmptyResult bool // 预期空结果
}

// SpecExpect 规格预期结果
type SpecExpect struct {
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
	RecordCreated     bool
	RecordUpdated     bool
	ShelvesMatch      int    // 预期的上下架状态
	RemarkMatch       string // 预期的备注
	SpecKeyMatch      string // 预期的SpecKey
	DeadlineTypeMatch int    // 预期的期限类型
}

// DataCheck 响应数据验证点
type DataCheck struct {
	TotalCount    int64  // 预期总数
	HasID         bool   // 返回数据包含ID
	HasSpecName   bool   // 返回数据包含规格名称
	HasShelves    bool   // 返回数据包含上下架状态
	IsEmptyResult bool   // 是否为空结果
}
