package debug_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/framework"
	"my_stonks_api_tests/testdata"
)

func TestDebug(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "调试测试套件")
}

// APICaller API调用器
type APICaller struct {
	client *framework.TestClient
}

// NewCaller 创建API调用器
func NewCaller() *APICaller {
	return &APICaller{
		client: framework.NewTestClient(),
	}
}

// Call 调用API
func (c *APICaller) Call(method, path string, body interface{}) *framework.TestResponse {
	var resp *framework.TestResponse
	var err error

	switch method {
	case "GET":
		resp, err = c.client.Get(path)
	case "POST":
		resp, err = c.client.Post(path, body)
	case "PUT":
		resp, err = c.client.Put(path, body)
	case "DELETE":
		resp, err = c.client.Delete(path)
	}

	if err != nil {
		panic(err)
	}
	return resp
}

// PrintResponse 打印响应
func PrintResponse(resp *framework.TestResponse) {
	GinkgoWriter.Printf("========== API响应 ==========\n")
	GinkgoWriter.Printf("状态码: %d\n", resp.StatusCode)
	GinkgoWriter.Printf("Code: %d\n", resp.Code)
	GinkgoWriter.Printf("消息: %s\n", resp.Message)
	GinkgoWriter.Printf("响应体: %s\n", resp.RawBody)
	GinkgoWriter.Printf("==============================\n")
}

var _ = Describe("快速调试", func() {
	It("快速创建环境并测试申购", func() {
		// 创建完整测试环境
		env := testdata.NewTestEnv("DEBUG").
			WithCoin().
			WithSpec(-1, 0).
			WithProduct(func(p map[string]interface{}) {
				p["minVol"] = testdata.DecimalFromInt(100)
			}).
			Build()

		env.WithBalance("5000")
		env.PrintEnvInfo()

		// 执行申购
		resp, err := env.Subscribe(testdata.DecimalFromInt(1000))
		Expect(err).ToNot(HaveOccurred())
		PrintResponse(resp)

		// 检查订单
		order := env.GetOrder()
		if order != nil {
			GinkgoWriter.Printf("订单创建成功: %+v\n", order)
		}
	})

	It("快速创建环境并测试赎回", func() {
		env := testdata.NewTestEnv("DEBUG").
			WithCoin().
			WithSpec(-1, 0).
			WithProduct().
			Build()

		env.WithBalance("5000")

		// 先申购
		subscribeResp, err := env.Subscribe(testdata.DecimalFromInt(1000))
		Expect(err).ToNot(HaveOccurred())
		PrintResponse(subscribeResp)

		order := env.GetOrder()
		Expect(order).ToNot(BeNil())

		// 再赎回
		redeemResp, err := env.Redeem(order.OrderNo, testdata.DecimalFromInt(500))
		Expect(err).ToNot(HaveOccurred())
		PrintResponse(redeemResp)
	})

	It("直接调用API测试", func() {
		caller := NewCaller()

		// 测试产品列表
		resp := caller.Call("GET", api.UserProductPage, nil)
		PrintResponse(resp)

		// 测试订单列表
		resp = caller.Call("POST", api.UserOrderList, map[string]interface{}{
			"page": 1,
			"size": 10,
		})
		PrintResponse(resp)
	})
})
