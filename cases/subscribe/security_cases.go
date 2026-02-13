package subscribe

// SecurityCase 安全测试用例
type SecurityCase struct {
	CaseID        string
	Module        string
	Priority      string
	Type          string
	Title         string
	Tags          []string
	PreCondition  []string
	ManualConfig  string   // 手动配置说明
	TestData      SecurityTestData
	Expected      SecurityExpect
}

// SecurityTestData 安全测试数据
type SecurityTestData struct {
	// 未登录场景
	NotLogin bool

	// IP黑名单场景
	IPBlacklisted bool

	// 2FA相关
	LargeAmount     string // 大额申购金额
	Need2FA         bool   // 是否需要2FA
	TwoFACode       string // 2FA验证码
	User2FAEnabled  bool   // 用户是否启用2FA
	User2FADisabled bool   // 用户是否未启用2FA
	SmallAmount     string // 小额申购金额
	Provide2FA      bool   // 是否提供2FA
}

// SecurityExpect 安全测试预期结果
type SecurityExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains []string
}

// SecurityCases 申购安全测试用例表
var SecurityCases = []SecurityCase{
	{
		CaseID:   "WTH_SUB_SEC_001",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "未登录应失败",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"未登录",
		},
		TestData: SecurityTestData{
			NotLogin: true,
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{"登录", "授权"},
		},
	},
	{
		CaseID:   "WTH_SUB_SEC_002",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "IP在黑名单应失败",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户IP在黑名单中",
		},
		ManualConfig: "在 sys_config 表中配置 wth_security_config，设置 ip_blacklist_enable=true，ip_blacklist=[\"127.0.0.1\"]",
		TestData: SecurityTestData{
			IPBlacklisted: true,
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{},
		},
	},
	{
		CaseID:   "WTH_SUB_SEC_003",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "大额申购未提供2FA应失败",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户已开启2FA",
			"申购金额>=10000",
		},
		ManualConfig: "配置 large_amount_2fa_enable=true, large_amount_threshold=10000, large_amount_2fa_for_subscribe=true",
		TestData: SecurityTestData{
			LargeAmount:    "10001",
			Need2FA:        true,
			User2FAEnabled: true,
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{"2FA", "验证"},
		},
	},
	{
		CaseID:   "WTH_SUB_SEC_004",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "大额申购2FA错误应失败",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户已开启2FA",
			"申购金额>=10000",
		},
		ManualConfig: "同 WTH_SUB_SEC_003",
		TestData: SecurityTestData{
			LargeAmount:    "10001",
			Need2FA:        true,
			TwoFACode:      "000000", // 错误的2FA
			User2FAEnabled: true,
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{"2FA", "验证"},
		},
	},
	{
		CaseID:   "WTH_SUB_SEC_005",
		Module:   "用户理财订单",
		Priority: "中",
		Type:     "安全测试",
		Title:    "小额申购不应要求2FA",
		Tags:     []string{"P1", "security"},
		PreCondition: []string{
			"已登录系统",
			"申购金额<10000",
		},
		ManualConfig: "配置 large_amount_2fa_enable=true, large_amount_threshold=10000",
		TestData: SecurityTestData{
			SmallAmount: "9999",
			Provide2FA:  true,
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{"2FA", "参数"},
		},
	},
	{
		CaseID:   "WTH_SUB_SEC_006",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "大额申购正确2FA应成功",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户已开启2FA",
			"申购金额>=10000",
		},
		ManualConfig: "同 WTH_SUB_SEC_003，需获取正确的2FA验证码",
		TestData: SecurityTestData{
			LargeAmount:     "10000",
			Need2FA:         true,
			TwoFACode:       "VALID_CODE", // 需替换为实际验证码
			User2FAEnabled:  true,
		},
		Expected: SecurityExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		CaseID:   "WTH_SUB_SEC_007",
		Module:   "用户理财订单",
		Priority: "中",
		Type:     "安全测试",
		Title:    "大额申购用户未启用2FA应失败",
		Tags:     []string{"P1", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户未开启2FA",
			"申购金额>=10000",
		},
		ManualConfig: "配置 large_amount_2fa_enable=true, 用户 two_fa_enabled=0",
		TestData: SecurityTestData{
			LargeAmount:      "10000",
			User2FADisabled:  true,
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{"2FA", "启用"},
		},
	},
	{
		CaseID:   "WTH_SUB_SEC_008",
		Module:   "用户理财订单",
		Priority: "中",
		Type:     "安全测试",
		Title:    "小额申购提供2FA应失败",
		Tags:     []string{"P1", "security"},
		PreCondition: []string{
			"已登录系统",
			"申购金额<10000",
		},
		ManualConfig: "配置 large_amount_2fa_enable=true, large_amount_threshold=10000",
		TestData: SecurityTestData{
			SmallAmount: "9999",
			Provide2FA:  true,
			TwoFACode:   "123456",
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{"2FA", "参数"},
		},
	},
}

// GetSecurityCaseByID 根据ID获取安全测试用例
func GetSecurityCaseByID(caseID string) *SecurityCase {
	for i := range SecurityCases {
		if SecurityCases[i].CaseID == caseID {
			return &SecurityCases[i]
		}
	}
	return nil
}
