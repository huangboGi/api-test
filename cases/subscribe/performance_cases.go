package subscribe

import (
	"time"

	"github.com/shopspring/decimal"
)

// PerformanceCase 申购性能测试用例
type PerformanceCase struct {
	CaseID       string
	Module       string
	Priority     string
	Type         string
	Title        string
	Tags         []string
	PreCondition []string
	TestData     PerformanceTestData
	Expect       PerformanceExpect
}

// PerformanceTestData 性能测试数据
type PerformanceTestData struct {
	SpecValue       int
	DeadlineType    int
	Volume          decimal.Decimal
	MinVol          decimal.Decimal
	ConcurrentUsers int
	Duration        time.Duration
	RampUp          time.Duration
}

// PerformanceExpect 性能测试预期结果
type PerformanceExpect struct {
	MaxResponseTime time.Duration // 最大响应时间
	AvgResponseTime time.Duration // 平均响应时间
	SuccessRate     float64       // 成功率 (0-1)
	MinTPS          float64       // 最小吞吐量
	MaxCPUPercent   float64       // 最大 CPU 使用率
	MaxMemPercent   float64       // 最大内存使用率
}

// PerformanceCases 申购性能测试用例表
var PerformanceCases = []PerformanceCase{
	{
		CaseID:   "WTH_SUB_PERF_001",
		Module:   "用户理财订单",
		Priority: "中",
		Type:     "性能测试",
		Title:    "申购接口响应时间测试",
		Tags:     []string{"P1", "performance"},
		PreCondition: []string{
			"已登录系统",
			"产品已上架",
			"用户余额充足",
		},
		TestData: PerformanceTestData{
			SpecValue:    -1,
			DeadlineType: 0,
			Volume:       decimal.NewFromInt(1000),
			MinVol:       decimal.NewFromInt(100),
		},
		Expect: PerformanceExpect{
			MaxResponseTime: 2 * time.Second,
			AvgResponseTime: 500 * time.Millisecond,
			SuccessRate:     1.0,
		},
	},
	{
		CaseID:   "WTH_SUB_PERF_002",
		Module:   "用户理财订单",
		Priority: "中",
		Type:     "性能测试",
		Title:    "申购接口并发压力测试",
		Tags:     []string{"P1", "performance"},
		PreCondition: []string{
			"已登录系统",
			"产品已上架，总额度充足",
			"多用户余额充足",
		},
		TestData: PerformanceTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			Volume:          decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentUsers: 100,
			Duration:        5 * time.Minute,
			RampUp:          30 * time.Second,
		},
		Expect: PerformanceExpect{
			MaxResponseTime: 3 * time.Second,
			AvgResponseTime: 1 * time.Second,
			SuccessRate:     0.99,
			MinTPS:          50,
			MaxCPUPercent:   80,
			MaxMemPercent:   80,
		},
	},
	{
		CaseID:   "WTH_SUB_PERF_003",
		Module:   "用户理财订单",
		Priority: "低",
		Type:     "性能测试",
		Title:    "申购接口持续负载测试",
		Tags:     []string{"P2", "performance"},
		PreCondition: []string{
			"已登录系统",
			"产品已上架，总额度充足",
			"多用户余额充足",
		},
		TestData: PerformanceTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			Volume:          decimal.NewFromInt(1000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentUsers: 50,
			Duration:        30 * time.Minute,
			RampUp:          1 * time.Minute,
		},
		Expect: PerformanceExpect{
			MaxResponseTime: 2 * time.Second,
			AvgResponseTime: 800 * time.Millisecond,
			SuccessRate:     0.999,
			MaxCPUPercent:   70,
			MaxMemPercent:   70,
		},
	},
	{
		CaseID:   "WTH_SUB_PERF_004",
		Module:   "用户理财订单",
		Priority: "低",
		Type:     "性能测试",
		Title:    "定期产品申购性能测试",
		Tags:     []string{"P2", "performance"},
		PreCondition: []string{
			"已登录系统",
			"定期产品已上架",
			"用户余额充足",
		},
		TestData: PerformanceTestData{
			DeadlineType:    1,
			Volume:          decimal.NewFromInt(5000),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentUsers: 200,
			Duration:        10 * time.Minute,
			RampUp:          1 * time.Minute,
		},
		Expect: PerformanceExpect{
			MaxResponseTime: 3 * time.Second,
			AvgResponseTime: 1 * time.Second,
			SuccessRate:     0.99,
			MinTPS:          30,
			MaxCPUPercent:   80,
			MaxMemPercent:   80,
		},
	},
	{
		CaseID:   "WTH_SUB_PERF_005",
		Module:   "用户理财订单",
		Priority: "低",
		Type:     "性能测试",
		Title:    "高频小额申购性能测试",
		Tags:     []string{"P2", "performance"},
		PreCondition: []string{
			"已登录系统",
			"活期产品已上架",
			"多用户余额充足",
		},
		TestData: PerformanceTestData{
			SpecValue:       -1,
			DeadlineType:    0,
			Volume:          decimal.NewFromInt(100),
			MinVol:          decimal.NewFromInt(100),
			ConcurrentUsers: 500,
			Duration:        5 * time.Minute,
			RampUp:          30 * time.Second,
		},
		Expect: PerformanceExpect{
			MaxResponseTime: 2 * time.Second,
			AvgResponseTime: 500 * time.Millisecond,
			SuccessRate:     0.99,
			MinTPS:          100,
			MaxCPUPercent:   85,
			MaxMemPercent:   85,
		},
	},
}

// GetPerformanceCaseByID 根据ID获取性能测试用例
func GetPerformanceCaseByID(caseID string) *PerformanceCase {
	for i := range PerformanceCases {
		if PerformanceCases[i].CaseID == caseID {
			return &PerformanceCases[i]
		}
	}
	return nil
}
