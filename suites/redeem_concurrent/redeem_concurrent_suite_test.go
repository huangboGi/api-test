package redeem_concurrent_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/redeem"
	"my_stonks_api_tests/framework"
	"my_stonks_api_tests/testdata"
)

func TestRedeemConcurrent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "赎回并发测试套件")
}

var _ = Describe("赎回并发测试", func() {
	It("并发赎回-同一订单多次赎回", func() {
		By("【用例编号】WTH_REDEEM_CONCURRENT_001")
		By("【优先级】高")

		env := testdata.NewTestEnv("REDEEM_CONCURRENT_001").
			WithCoin().
			WithSpec(30, 1).
			WithProduct(nil).
			WithBalance("100000").
			Build()

		// 先申购
		_, err := env.Subscribe(testdata.DecimalFromString("1000"))
		Expect(err).ToNot(HaveOccurred())

		// 获取订单
		order := env.GetOrder()
		Expect(order).ToNot(BeNil())

		// 并发赎回
		concurrency := 3
		results := make(chan *framework.TestResponse, concurrency)

		for i := 0; i < concurrency; i++ {
			go func() {
				resp, _ := env.Client.Post(api.UserRedeem, map[string]interface{}{
					"orderNo": order.OrderNo,
					"volume":  "1000",
				})
				results <- resp
			}()
		}

		// 收集结果
		successCount := 0
		for i := 0; i < concurrency; i++ {
			select {
			case resp := <-results:
				if resp != nil && resp.Code == 0 {
					successCount++
				}
			case <-time.After(10 * time.Second):
				Fail("并发请求超时")
			}
		}

		// 验证：应该只有一个成功
		Expect(successCount).To(BeNumerically("<=", 1))
		env.LogResult(true, "并发赎回测试通过")
	})
})

var _ = Describe("赎回并发测试-验证用例表", func() {
	for _, tc := range redeem.ConcurrentCases {
		tc := tc
		It(tc.Title, Label(tc.CaseID), func() {
			By("【用例编号】" + tc.CaseID)

			env := testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				WithSpec(tc.TestData.SpecValue, tc.TestData.DeadlineType).
				WithProduct(nil).
				WithBalance("100000").
				Build()

			// 先申购
			_, err := env.Subscribe(tc.TestData.SubscribeVolume)
			Expect(err).ToNot(HaveOccurred())

			order := env.GetOrder()
			Expect(order).ToNot(BeNil())

			// 并发赎回
			results := make(chan *framework.TestResponse, tc.TestData.ConcurrentCount)
			for i := 0; i < tc.TestData.ConcurrentCount; i++ {
				go func() {
					resp, _ := env.Client.Post(api.UserRedeem, map[string]interface{}{
						"orderNo": order.OrderNo,
						"volume":  tc.TestData.RedeemVolume.String(),
					})
					results <- resp
				}()
			}

			// 收集结果
			successCount := 0
			for i := 0; i < tc.TestData.ConcurrentCount; i++ {
				select {
				case resp := <-results:
					if resp != nil && resp.Code == 0 {
						successCount++
					}
				case <-time.After(10 * time.Second):
					Fail("并发请求超时")
				}
			}

			// 验证
			if tc.Expect.MaxSuccessCount > 0 {
				Expect(successCount).To(BeNumerically("<=", tc.Expect.MaxSuccessCount))
			}
			env.LogResult(true, "并发测试通过")
		})
	}
})
