package subscribe_security_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/subscribe"
	"my_stonks_api_tests/framework"
	"my_stonks_api_tests/testdata"
)

func TestSubscribeSecurity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "申购安全测试套件")
}

var _ = Describe("申购安全测试", func() {
	for _, tc := range subscribe.SecurityCases {
		tc := tc
		It(tc.Title, Label(tc.Tags...), func() {
			By("【用例编号】" + tc.CaseID)
			By("【优先级】" + tc.Priority)
			By("【前置条件】" + joinPreConditions(tc.PreCondition))

			// 准备环境
			client := framework.NewTestClient()

			// 未登录场景
			if tc.TestData.NotLogin {
				client.SetAuthToken("")
				resp, err := client.Post(api.UserSubscribe, map[string]interface{}{
					"coin":         "TEST",
					"specValue":    -1,
					"volume":       "1000",
					"deadlineType": 0,
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
				for _, msg := range tc.Expected.ErrMsgContains {
					Expect(resp.Message).To(Or(ContainSubstring(msg), ContainSubstring("登录")))
				}
				return
			}

			// 正常登录场景
			env := testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				WithSpec(-1, 0).
				WithProduct(func(p map[string]interface{}) {
					p["minVol"] = testdata.DecimalFromInt(100)
				}).
				WithBalance("20000").
				Build()

			// 构建请求
			req := map[string]interface{}{
				"coin":         env.Coin,
				"specValue":    env.SpecValue,
				"deadlineType": env.DeadlineType,
			}

			// 根据测试数据设置金额
			if tc.TestData.LargeAmount != "" {
				req["volume"] = tc.TestData.LargeAmount
			} else if tc.TestData.SmallAmount != "" {
				req["volume"] = tc.TestData.SmallAmount
			} else {
				req["volume"] = "1000"
			}

			// 2FA相关
			if tc.TestData.Provide2FA || tc.TestData.TwoFACode != "" {
				req["twoFACode"] = tc.TestData.TwoFACode
			}

			// 手动配置提醒
			if tc.ManualConfig != "" {
				By("【手动配置】" + tc.ManualConfig)
			}

			// 执行申购
			resp, err := client.Post(api.UserSubscribe, req)
			Expect(err).ToNot(HaveOccurred())

			// 验证结果
			if tc.Expected.Success {
				Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
				Expect(resp.Code).To(Equal(0))
				env.LogResult(true, "安全测试通过")
			} else {
				Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
				if len(tc.Expected.ErrMsgContains) > 0 {
					for _, msg := range tc.Expected.ErrMsgContains {
						Expect(resp.Message).To(ContainSubstring(msg))
					}
				}
				env.LogResult(false, "安全测试失败（预期）: "+resp.Message)
			}
		})
	}
})

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
