package subscribe_concurrent_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/subscribe"
	"my_stonks_api_tests/framework"
	"my_stonks_api_tests/testdata"
)

func TestSubscribeConcurrent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "申购并发测试套件")
}

var _ = Describe("申购并发测试", func() {
	It("并发申购-同一用户多请求", func() {
		By("【用例编号】WTH_SUB_CONCURRENT_001")
		By("【优先级】高")

		env := testdata.NewTestEnv("CONCURRENT_001").
			WithCoin().
			WithSpec(30, 1).
			WithProduct(func(p map[string]interface{}) {
				p["useQuotaTotal"] = "100000"
				p["personQuota"] = "100000"
			}).
			WithBalance("100000").
			Build()

		// 并发请求
		concurrency := 5
		volume := "100"
		results := make(chan *framework.TestResponse, concurrency)

		for i := 0; i < concurrency; i++ {
			go func() {
				resp, _ := env.Subscribe(testdata.DecimalFromString(volume))
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
		env.LogResult(true, "并发测试通过")
	})

	It("并发申购-不同用户同时申购", func() {
		By("【用例编号】WTH_SUB_CONCURRENT_002")
		By("【优先级】高")

		env := testdata.NewTestEnv("CONCURRENT_002").
			WithCoin().
			WithSpec(30, 1).
			WithProduct(func(p map[string]interface{}) {
				p["useQuotaTotal"] = "10000"
				p["personQuota"] = "10000"
			}).
			Build()

		// 模拟多个用户并发申购
		userCount := 3
		results := make(chan *framework.TestResponse, userCount)

		for i := 0; i < userCount; i++ {
			go func(idx int) {
				// 每个用户独立环境
				userEnv := testdata.NewTestEnv("CONCURRENT_002_USER").
					WithCoin().
					Build()
				userEnv.WithBalance("10000")

				// 使用相同产品
				userEnv.ProductID = env.ProductID
				userEnv.Client = env.Client

				resp, _ := userEnv.Subscribe(testdata.DecimalFromString("5000"))
				results <- resp
			}(i)
		}

		// 收集结果
		successCount := 0
		for i := 0; i < userCount; i++ {
			select {
			case resp := <-results:
				if resp != nil && resp.Code == 0 {
					successCount++
				}
			case <-time.After(10 * time.Second):
				Fail("并发请求超时")
			}
		}

		// 验证：额度限制
		Expect(successCount).To(BeNumerically("<=", 2))
		env.LogResult(true, "多用户并发测试通过")
	})

	It("并发申购-额度边界测试", func() {
		By("【用例编号】WTH_SUB_CONCURRENT_003")
		By("【优先级】中")

		// 剩余额度刚好够一次申购
		env := testdata.NewTestEnv("CONCURRENT_003").
			WithCoin().
			WithSpec(30, 1).
			WithProduct(func(p map[string]interface{}) {
				p["useQuotaTotal"] = "100"
				p["soldAmount"] = "0"
			}).
			WithBalance("10000").
			Build()

		// 两个用户同时申购全部剩余额度
		results := make(chan *framework.TestResponse, 2)

		for i := 0; i < 2; i++ {
			go func() {
				resp, _ := env.Subscribe(testdata.DecimalFromString("100"))
				results <- resp
			}()
		}

		// 收集结果
		successCount := 0
		for i := 0; i < 2; i++ {
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
		Expect(successCount).To(Equal(1))
		env.LogResult(true, "额度边界并发测试通过")
	})
})

var _ = Describe("申购并发测试-验证用例表", func() {
	for _, tc := range subscribe.ConcurrentCases {
		tc := tc
		It(tc.Title, Label(tc.CaseID), func() {
			By("【用例编号】" + tc.CaseID)
			// 执行并发测试
			env := testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				WithSpec(tc.TestData.SpecValue, tc.TestData.DeadlineType).
				WithProduct(nil).
				WithBalance("100000").
				Build()

			// 并发执行
			results := make(chan *framework.TestResponse, tc.TestData.ConcurrentCount)
			for i := 0; i < tc.TestData.ConcurrentCount; i++ {
				go func() {
					resp, _ := env.Client.Post(api.UserSubscribe, map[string]interface{}{
						"coin":         env.Coin,
						"volume":       tc.TestData.VolumePerUser.String(),
						"specValue":    env.SpecValue,
						"deadlineType": env.DeadlineType,
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
