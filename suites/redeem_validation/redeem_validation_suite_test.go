package redeem_validation_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/redeem"
	"my_stonks_api_tests/testdata"
)

func TestRedeemValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "赎回验证测试套件")
}

// RunValidationTest 执行单个赎回验证测试用例
func RunValidationTest(tc redeem.ValidationCase) {
	It(tc.Title, func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)
		By("【前置条件】" + joinPreConditions(tc.PreCondition))

		var env *testdata.TestEnv
		var orderNo string

		// 根据输入条件准备环境
		if tc.Input.OrderNo == "" {
			// 空订单号场景 - 不需要创建环境
			resp, err := testdata.NewTestEnv(tc.CaseID).Client.Post(api.UserRedeem, map[string]interface{}{
				"orderNo": "",
				"volume":  tc.Input.RedeemVolume,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Or(Equal(tc.Expected.StatusCode), Equal(200)))
			if resp.Code != 0 {
				Expect(resp.Message).To(ContainSubstring(tc.Expected.ErrMsgContains))
			}
			return
		} else if tc.Input.OrderNo == "INVALID" {
			// 无效订单号场景
			resp, err := testdata.NewTestEnv(tc.CaseID).Client.Post(api.UserRedeem, map[string]interface{}{
				"orderNo": "ORDER_NOT_EXIST_123456",
				"volume":  tc.Input.RedeemVolume,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Or(Equal(tc.Expected.StatusCode), Equal(200)))
			if resp.Code != 0 {
				Expect(resp.Message).To(ContainSubstring(tc.Expected.ErrMsgContains))
			}
			return
		} else if tc.Input.OrderCompleted {
			// 已完成订单场景
			env = testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				WithSpec(-1, 0).
				WithProduct().
				Build()
			env.WithBalance(tc.Input.SubscribeVolume.Mul(testdata.DecimalFromInt(2)).String())

			// 申购
			subscribeResp, err := env.Subscribe(tc.Input.SubscribeVolume)
			Expect(err).ToNot(HaveOccurred())
			Expect(subscribeResp.Code).To(Equal(0))

			// 获取订单
			order := env.GetOrder()
			Expect(order).ToNot(BeNil())
			orderNo = order.OrderNo

			// 全额赎回使订单完成
			env.Redeem(orderNo, tc.Input.SubscribeVolume)
		} else {
			// 正常环境
			env = testdata.NewTestEnv(tc.CaseID).
				WithCoin().
				WithSpec(-1, 0).
				WithProduct(func(p map[string]interface{}) {
					p["minVol"] = tc.Input.SubscribeVolume
					if tc.Input.DailyLimit.IsPositive() {
						p["dailyLimit"] = tc.Input.DailyLimit
					}
				}).
				Build()
			env.WithBalance(tc.Input.SubscribeVolume.Mul(testdata.DecimalFromInt(2)).String())

			// 申购
			subscribeResp, err := env.Subscribe(tc.Input.SubscribeVolume)
			Expect(err).ToNot(HaveOccurred())
			Expect(subscribeResp.Code).To(Equal(0))

			// 获取订单
			order := env.GetOrder()
			Expect(order).ToNot(BeNil())
			orderNo = order.OrderNo
		}

		// 执行赎回
		resp, err := env.Client.Post(api.UserRedeem, map[string]interface{}{
			"orderNo": orderNo,
			"volume":  tc.Input.RedeemVolume,
		})
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "赎回成功")
		} else {
			Expect(resp.StatusCode).To(Or(Equal(tc.Expected.StatusCode), Equal(200)))
			if resp.Code != 0 {
				Expect(resp.Message).To(ContainSubstring(tc.Expected.ErrMsgContains))
			}
			env.LogResult(false, "赎回失败（预期）: "+resp.Message)
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

var _ = Describe("赎回验证测试", func() {
	for _, tc := range redeem.ValidationCases {
		tc := tc
		RunValidationTest(tc)
	}
})

var _ = Describe("赎回验证P0用例", Label("P0"), func() {
	for _, tc := range redeem.ValidationCases {
		tc := tc
		for _, tag := range tc.Tags {
			if tag == "P0" {
				RunValidationTest(tc)
				break
			}
		}
	}
})
