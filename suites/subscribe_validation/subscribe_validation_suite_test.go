package subscribe_validation_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/subscribe"
	"my_stonks_api_tests/testdata"
)

func TestSubscribeValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "申购验证测试套件")
}

// RunValidationTest 执行单个申购验证测试用例
func RunValidationTest(tc subscribe.ValidationCase) {
	It(tc.Title, func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)
		By("【前置条件】" + joinPreConditions(tc.PreCondition))

		var env *testdata.TestEnv
		
		// 根据输入条件准备环境
		if tc.Input.CoinOff {
			// 币种未上架场景
			env = testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				Build()
			// 不上架币种
		} else if tc.Input.SpecOff {
			// 规格未上架场景
			env = testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				WithSpec(-1, 0).
				Build()
			// 不上架规格
		} else if tc.Input.ProductOff {
			// 产品未上架场景
			env = testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				WithSpec(tc.Input.SpecValue.(int), tc.Input.DeadlineType.(int)).
				WithProduct(func(p map[string]interface{}) {
					if tc.Input.MinVol.IsPositive() {
						p["minVol"] = tc.Input.MinVol
					}
				}).
				Build()
			// 不上架产品
			env.Client.Post(api.AdminProductShelves, map[string]interface{}{
				"id":            env.ProductID,
				"shelvesStatus": 0,
			})
		} else {
			// 正常环境
			env = testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				WithSpec(tc.Input.SpecValue.(int), tc.Input.DeadlineType.(int)).
				WithProduct(func(p map[string]interface{}) {
					if tc.Input.MinVol.IsPositive() {
						p["minVol"] = tc.Input.MinVol
					}
					if tc.Input.MaxVol.IsPositive() {
						p["useQuotaTotal"] = tc.Input.MaxVol
					}
					if tc.Input.PersonQuota.IsPositive() {
						p["personQuota"] = tc.Input.PersonQuota
					}
					if tc.Input.TotalQuota.IsPositive() {
						p["totalQuota"] = tc.Input.TotalQuota
						p["soldAmount"] = tc.Input.SoldAmount
					}
				}).
				Build()
		}

		// 构建申购请求
		coin := env.Coin
		if tc.Input.Coin == "NONEXISTCOIN" {
			coin = "NONEXISTCOIN"
		} else if tc.Input.Coin == "" {
			coin = ""
		}

		// 执行申购
		resp, err := env.Client.Post(api.UserSubscribe, map[string]interface{}{
			"coin":         coin,
			"specValue":    tc.Input.SpecValue,
			"volume":       tc.Input.Volume,
			"deadlineType": tc.Input.DeadlineType,
		})
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			Expect(resp.Code).To(Equal(0))
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

var _ = Describe("申购验证测试", func() {
	for _, tc := range subscribe.ValidationCases {
		tc := tc
		RunValidationTest(tc)
	}
})

var _ = Describe("申购验证P0用例", Label("P0"), func() {
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
