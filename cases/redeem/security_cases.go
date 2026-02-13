package redeem

// SecurityCase 赎回安全测试用例
type SecurityCase struct {
	CaseID        string
	Module        string
	Priority      string
	Type          string
	Title         string
	Tags          []string
	PreCondition  []string
	ManualConfig  string
	TestData      SecurityTestData
	Expected      SecurityExpect
}

// SecurityTestData 赎回安全测试数据
type SecurityTestData struct {
	NotLogin      bool
	IPBlacklisted bool
	LargeAmount   string
	Need2FA       bool
	TwoFACode     string
	User2FAEnabled  bool
	User2FADisabled bool
	SmallAmount   string
	Provide2FA    bool
	OtherUserOrder bool
}

// SecurityExpect 赎回安全测试预期
type SecurityExpect struct {
	Success        bool
	StatusCode     int
	ErrMsgContains []string
}

// SecurityCases 赎回安全测试用例表
var SecurityCases = []SecurityCase{
	{
		CaseID:   "WTH_RED_SEC_001",
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
		CaseID:   "WTH_RED_SEC_002",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "IP在黑名单应失败",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户IP在黑名单中",
		},
		ManualConfig: "在 sys_config 表中配置 wth_security_config",
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
		CaseID:   "WTH_RED_SEC_003",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "赎回不属于当前用户的订单应失败",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"存在其他用户的订单",
		},
		TestData: SecurityTestData{
			OtherUserOrder: true,
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{"订单"},
		},
	},
	{
		CaseID:   "WTH_RED_SEC_004",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "大额赎回未提供2FA应失败",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户已开启2FA",
			"赎回金额>=10000",
		},
		ManualConfig: "配置 large_amount_2fa_for_redeem=true",
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
		CaseID:   "WTH_RED_SEC_005",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "大额赎回2FA错误应失败",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户已开启2FA",
			"赎回金额>=10000",
		},
		TestData: SecurityTestData{
			LargeAmount:    "10001",
			Need2FA:        true,
			TwoFACode:      "000000",
			User2FAEnabled: true,
		},
		Expected: SecurityExpect{
			Success:        false,
			StatusCode:     200,
			ErrMsgContains: []string{"2FA", "验证"},
		},
	},
	{
		CaseID:   "WTH_RED_SEC_006",
		Module:   "用户理财订单",
		Priority: "中",
		Type:     "安全测试",
		Title:    "小额赎回不应要求2FA",
		Tags:     []string{"P1", "security"},
		PreCondition: []string{
			"已登录系统",
			"赎回金额<10000",
		},
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
		CaseID:   "WTH_RED_SEC_007",
		Module:   "用户理财订单",
		Priority: "高",
		Type:     "安全测试",
		Title:    "大额赎回正确2FA应成功",
		Tags:     []string{"P0", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户已开启2FA",
			"赎回金额>=10000",
		},
		ManualConfig: "需获取正确的2FA验证码",
		TestData: SecurityTestData{
			LargeAmount:     "10000",
			Need2FA:         true,
			TwoFACode:       "VALID_CODE",
			User2FAEnabled:  true,
		},
		Expected: SecurityExpect{
			Success:    true,
			StatusCode: 200,
		},
	},
	{
		CaseID:   "WTH_RED_SEC_008",
		Module:   "用户理财订单",
		Priority: "中",
		Type:     "安全测试",
		Title:    "大额赎回用户未启用2FA应失败",
		Tags:     []string{"P1", "security"},
		PreCondition: []string{
			"已登录系统",
			"用户未开启2FA",
			"赎回金额>=10000",
		},
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
