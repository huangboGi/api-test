package subscribe_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/subscribe"
	"my_stonks_api_tests/testdata"
)

func TestSubscribe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "申购功能测试套件")
}

// RunFuncTest 执行单个申购功能测试用例
func RunFuncTest(tc subscribe.FuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)
		By("【前置条件】" + joinPreConditions(tc.PreCondition))

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(tc.Input.SpecValue, tc.Input.DeadlineType).
			WithProduct(func(p map[string]interface{}) {
				p["minVol"] = tc.Input.MinVol
				if tc.Input.MaxVol.IsPositive() {
					p["useQuotaTotal"] = tc.Input.MaxVol
				}
			}).
			Build()

		// 如果期望成功，检查余额
		if tc.Expected.Success {
			env.WithBalance(tc.Input.Volume.Mul(testdata.DecimalFromInt(2)).String())
		}

		// 如果需要产品下架
		if tc.Input.ProductOff {
			env.Client.Post(api.AdminProductShelves, map[string]interface{}{
				"id":            env.ProductID,
				"shelvesStatus": 0,
			})
		}

		// 如果需要规格下架
		if tc.Input.SpecOff {
			env.Client.Post(api.AdminSpecShelves, map[string]interface{}{
				"id":            env.SpecID,
				"shelvesStatus": 0,
			})
		}

		// 获取初始余额
		balanceBefore := env.GetBalance(env.Coin)

		// 执行申购
		resp, err := env.Subscribe(tc.Input.Volume)
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			Expect(resp.Code).To(Equal(0))

			// 数据库验证
			if tc.Expected.DBCheck.OrderCreated {
				order := env.GetOrder()
				Expect(order).ToNot(BeNil())
				if tc.Expected.DBCheck.VolumeMatch {
					Expect(order.Volume.String()).To(Equal(tc.Input.Volume.String()))
				}
			}

			if tc.Expected.DBCheck.BalanceChanged {
				expectedBalance := balanceBefore.Sub(tc.Input.Volume)
				env.BalanceShouldBe(env.Coin, expectedBalance.String())
			}

			if tc.Expected.DBCheck.HisCreated {
				order := env.GetOrder()
				if order != nil {
					his := env.GetSubscribeHis(int64(order.ID))
					Expect(his).ToNot(BeNil())
				}
			}

			env.LogResult(true, "申购成功")
		} else {
			Expect(resp.StatusCode).To(Or(Equal(tc.Expected.StatusCode), Equal(200)))
			if resp.Code != 0 {
				Expect(resp.Message).To(ContainSubstring(tc.Expected.ErrMsgContains))
			}
			env.LogResult(false, "申购失败（预期）: "+resp.Message)
		}
	})
}

// RunValidationTest 执行验证测试用例
func RunValidationTest(tc subscribe.ValidationCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)
		By("【前置条件】" + joinPreConditions(tc.PreCondition))

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(getSpecValue(tc.Input.SpecValue), getDeadlineType(tc.Input.DeadlineType)).
			WithProduct(func(p map[string]interface{}) {
				p["minVol"] = tc.Input.MinVol
				if tc.Input.MaxVol.IsPositive() {
					p["useQuotaTotal"] = tc.Input.MaxVol
				}
				if tc.Input.PersonQuota.IsPositive() {
					p["useQuotaPersonal"] = tc.Input.PersonQuota
				}
				if tc.Input.TotalQuota.IsPositive() {
					p["useQuotaTotal"] = tc.Input.TotalQuota
				}
				if tc.Input.DailyLimit.IsPositive() {
					p["dailyLimit"] = tc.Input.DailyLimit
				}
			}).
			Build()

		// 如果期望成功，设置余额
		if tc.Expected.Success {
			env.WithBalance(tc.Input.Volume.Mul(testdata.DecimalFromInt(2)).String())
		}

		// 处理特殊输入
		coin := env.Coin
		if tc.Input.Coin == "" {
			coin = "" // 空字符串
		} else if tc.Input.Coin == "NONEXISTCOIN" {
			coin = "NONEXISTCOIN"
		}

		// 如果需要币种下架
		if tc.Input.CoinOff {
			env.Client.Post(api.AdminCoinShelves, map[string]interface{}{
				"id":            env.CoinID,
				"shelvesStatus": 0,
			})
		}

		// 如果需要产品下架
		if tc.Input.ProductOff {
			env.Client.Post(api.AdminProductShelves, map[string]interface{}{
				"id":            env.ProductID,
				"shelvesStatus": 0,
			})
		}

		// 如果需要规格下架
		if tc.Input.SpecOff {
			env.Client.Post(api.AdminSpecShelves, map[string]interface{}{
				"id":            env.SpecID,
				"shelvesStatus": 0,
			})
		}

		// 构建请求
		req := map[string]interface{}{
			"coin":         coin,
			"specValue":    getSpecValue(tc.Input.SpecValue),
			"deadlineType": getDeadlineType(tc.Input.DeadlineType),
			"volume":       tc.Input.Volume.String(),
		}

		// 执行申购
		resp, err := env.Client.Post(api.UserSubscribe, req)
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "申购成功")
		} else {
			Expect(resp.StatusCode).To(Or(Equal(tc.Expected.StatusCode), Equal(200)))
			if resp.Code != 0 && tc.Expected.ErrMsgContains != "" {
				Expect(resp.Message).To(ContainSubstring(tc.Expected.ErrMsgContains))
			}
			env.LogResult(false, "申购失败（预期）: "+resp.Message)
		}
	})
}

// RunPerformanceTest 执行性能测试用例
func RunPerformanceTest(tc subscribe.PerformanceCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)
		By("【测试类型】" + tc.Type)

		// 性能测试需要特殊环境配置
		// 这里只做基础验证，实际性能测试应使用专业工具如 k6、JMeter
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(tc.TestData.SpecValue, tc.TestData.DeadlineType).
			WithProduct(func(p map[string]interface{}) {
				p["minVol"] = tc.TestData.MinVol
			}).
			Build()
		env.WithBalance(tc.TestData.Volume.Mul(testdata.DecimalFromInt(2)).String())

		// 单次请求响应时间测试
		startTime := time.Now()
		resp, err := env.Subscribe(tc.TestData.Volume)
		elapsed := time.Since(startTime)

		Expect(err).ToNot(HaveOccurred())
		Expect(resp.Code).To(Equal(0))
		Expect(elapsed).To(BeNumerically("<", tc.Expect.MaxResponseTime), "响应时间应小于最大响应时间")

		env.LogResult(true, "性能测试通过，响应时间: "+elapsed.String())
	})
}

// getSpecValue 获取规格值
func getSpecValue(v interface{}) int {
	switch val := v.(type) {
	case int:
		return val
	case string:
		if val == "" {
			return 0 // 空字符串表示不传
		}
		return -1
	default:
		return -1
	}
}

// getDeadlineType 获取期限类型
func getDeadlineType(v interface{}) int {
	switch val := v.(type) {
	case int:
		return val
	case string:
		return 0
	default:
		return 0
	}
}

// joinPreConditions 拼接前置条件
func joinPreConditions(conditions []string) string {
	result := ""
	for i, c := range conditions {
		if i > 0 {
			result += "; "
		}
		result += c
	}
	return result
}

var _ = Describe("申购功能测试", func() {
	for _, tc := range subscribe.FuncCases {
		tc := tc // 捕获变量
		if tc.Type == "功能测试" {
			RunFuncTest(tc)
		}
	}
})

var _ = Describe("申购验证测试", func() {
	for _, tc := range subscribe.ValidationCases {
		tc := tc
		if tc.Type == "逆向测试" || tc.Type == "边界测试" {
			RunValidationTest(tc)
		}
	}
})

var _ = Describe("申购性能测试", Label("performance"), func() {
	for _, tc := range subscribe.PerformanceCases {
		tc := tc
		RunPerformanceTest(tc)
	}
})

var _ = Describe("申购P0用例", Label("P0"), func() {
	for _, tc := range subscribe.FuncCases {
		tc := tc
		for _, tag := range tc.Tags {
			if tag == "P0" {
				RunFuncTest(tc)
				break
			}
		}
	}
	for _, tc := range subscribe.ValidationCases {
		tc := tc
		for _, tag := range tc.Tags {
			if tag == "P0" {
				RunValidationTest(tc)
				break
			}
		}
	}
})

var _ = Describe("申购冒烟测试", Label("smoke"), func() {
	for _, tc := range subscribe.FuncCases {
		tc := tc
		for _, tag := range tc.Tags {
			if tag == "smoke" {
				RunFuncTest(tc)
				break
			}
		}
	}
})
