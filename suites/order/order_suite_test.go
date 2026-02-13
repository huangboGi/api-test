package order_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/order"
	"my_stonks_api_tests/testdata"
)

func TestOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "订单查询测试套件")
}

// RunPageTest 执行分页查询测试
func RunPageTest(tc order.PageCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(-1, 0).
			WithProduct(nil).
			Build()

		// 构建请求参数
		req := map[string]interface{}{
			"pageIndex": tc.Input.PageIndex,
			"pageSize":  tc.Input.PageSize,
		}
		if tc.Input.Coin != "" {
			req["coin"] = tc.Input.Coin
		}
		if tc.Input.Status != nil {
			req["status"] = *tc.Input.Status
		}
		if tc.Input.SpecValue != nil {
			req["specValue"] = *tc.Input.SpecValue
		}

		// 执行查询
		resp, err := env.Client.Post(api.UserOrderList, req)
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			env.LogResult(true, "查询成功")
		} else {
			if tc.Expected.ErrMsgContains != "" {
				Expect(resp.Message).To(ContainSubstring(tc.Expected.ErrMsgContains))
			}
			env.LogResult(false, "查询失败（预期）: "+resp.Message)
		}
	})
}

// RunDetailTest 执行订单详情测试
func RunDetailTest(tc order.DetailCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		env := testdata.NewTestEnv(tc.CaseID).Build()

		// 构建请求参数
		req := map[string]interface{}{}
		if tc.Input.OrderID > 0 {
			req["orderId"] = tc.Input.OrderID
		}

		// 执行查询
		resp, err := env.Client.Post(api.UserOrderDetail, req)
		Expect(err).ToNot(HaveOccurred())

		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			env.LogResult(true, "查询成功")
		} else {
			env.LogResult(false, "查询失败（预期）: "+resp.Message)
		}
	})
}

// RunHisTest 执行历史记录测试
func RunHisTest(tc order.HisCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		env := testdata.NewTestEnv(tc.CaseID).Build()

		// 构建请求参数
		req := map[string]interface{}{
			"pageIndex": 1,
			"pageSize":  10,
		}
		if tc.Input.Coin != "" {
			req["coin"] = tc.Input.Coin
		}

		// 执行查询
		resp, err := env.Client.Post(api.UserHis, req)
		Expect(err).ToNot(HaveOccurred())

		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			env.LogResult(true, "查询成功")
		} else {
			env.LogResult(false, "查询失败（预期）: "+resp.Message)
		}
	})
}

// RunHoldPositionTest 执行持仓查询测试
func RunHoldPositionTest(tc order.HoldPositionCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		env := testdata.NewTestEnv(tc.CaseID).Build()

		// 构建请求参数
		req := map[string]interface{}{}
		if tc.Input.Coin != "" {
			req["coin"] = tc.Input.Coin
		}

		// 执行查询
		resp, err := env.Client.Post(api.UserHoldPosition, req)
		Expect(err).ToNot(HaveOccurred())

		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			env.LogResult(true, "查询成功")
		} else {
			env.LogResult(false, "查询失败（预期）: "+resp.Message)
		}
	})
}

// RunInterestPageTest 执行收益明细测试
func RunInterestPageTest(tc order.InterestPageCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		env := testdata.NewTestEnv(tc.CaseID).Build()

		// 构建请求参数
		req := map[string]interface{}{}
		if tc.Input.OrderID > 0 {
			req["orderId"] = tc.Input.OrderID
		}

		// 执行查询
		resp, err := env.Client.Post(api.UserInterestPage, req)
		Expect(err).ToNot(HaveOccurred())

		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			env.LogResult(true, "查询成功")
		} else {
			env.LogResult(false, "查询失败（预期）: "+resp.Message)
		}
	})
}

var _ = Describe("订单分页查询测试", func() {
	for _, tc := range order.PageCases {
		tc := tc
		RunPageTest(tc)
	}
})

var _ = Describe("订单详情查询测试", func() {
	for _, tc := range order.DetailCases {
		tc := tc
		RunDetailTest(tc)
	}
})

var _ = Describe("历史记录查询测试", func() {
	for _, tc := range order.HisCases {
		tc := tc
		RunHisTest(tc)
	}
})

var _ = Describe("持仓查询测试", func() {
	for _, tc := range order.HoldPositionCases {
		tc := tc
		RunHoldPositionTest(tc)
	}
})

var _ = Describe("收益明细查询测试", func() {
	for _, tc := range order.InterestPageCases {
		tc := tc
		RunInterestPageTest(tc)
	}
})
