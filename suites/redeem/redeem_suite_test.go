package redeem_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/redeem"
	"my_stonks_api_tests/testdata"
)

func TestRedeem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "赎回功能测试套件")
}

// RunFuncTest 执行单个赎回功能测试用例
func RunFuncTest(tc redeem.FuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)
		By("【前置条件】" + joinPreConditions(tc.PreCondition))

		// 特殊情况：无效订单号
		if tc.Input.InvalidOrderNo {
			resp, err := testdata.NewTestEnv(tc.CaseID).Client.Post(api.UserRedeem, map[string]interface{}{
				"orderNo": "ORDER_NOT_EXIST_123456",
				"volume":  tc.Input.RedeemVolume,
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			if resp.Code != 0 {
				Expect(resp.Message).To(ContainSubstring(tc.Expected.ErrMsgContains))
			}
			return
		}

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

		// 检查余额并申购
		env.WithBalance(tc.Input.SubscribeVolume.Mul(testdata.DecimalFromInt(2)).String())

		// 执行申购
		subscribeResp, err := env.Subscribe(tc.Input.SubscribeVolume)
		Expect(err).ToNot(HaveOccurred())
		Expect(subscribeResp.StatusCode).To(Equal(200))
		Expect(subscribeResp.Code).To(Equal(0))

		// 获取订单
		order := env.GetOrder()
		Expect(order).ToNot(BeNil())

		// 如果是已完成订单场景，先全额赎回
		if tc.Input.OrderCompleted {
			env.Redeem(order.OrderNo, tc.Input.SubscribeVolume)
		}

		// 获取初始余额
		balanceBefore := env.GetBalance(env.Coin)

		// 执行赎回
		resp, err := env.Redeem(order.OrderNo, tc.Input.RedeemVolume)
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expected.Success {
			Expect(resp.StatusCode).To(Equal(tc.Expected.StatusCode))
			Expect(resp.Code).To(Equal(0))

			// 数据库验证
			if tc.Expected.DBCheck.BalanceChanged {
				expectedBalance := balanceBefore.Add(tc.Input.RedeemVolume)
				env.BalanceShouldBe(env.Coin, expectedBalance.String())
			}

			if tc.Expected.DBCheck.OrderCompleted {
				updatedOrder := env.GetOrderByNo(order.OrderNo)
				Expect(updatedOrder).ToNot(BeNil())
				Expect(updatedOrder.Status).To(Equal(int8(1))) // 1-已完成
			}

			if tc.Expected.DBCheck.VolumeRemain.IsPositive() {
				updatedOrder := env.GetOrderByNo(order.OrderNo)
				Expect(updatedOrder).ToNot(BeNil())
				Expect(updatedOrder.Volume.String()).To(Equal(tc.Expected.DBCheck.VolumeRemain.String()))
			}

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

var _ = Describe("赎回功能测试", func() {
	for _, tc := range redeem.FuncCases {
		tc := tc
		if tc.Type == "功能测试" {
			RunFuncTest(tc)
		}
	}
})

var _ = Describe("赎回P0用例", Label("P0"), func() {
	for _, tc := range redeem.FuncCases {
		tc := tc
		for _, tag := range tc.Tags {
			if tag == "P0" {
				RunFuncTest(tc)
				break
			}
		}
	}
})

var _ = Describe("赎回冒烟测试", Label("smoke"), func() {
	for _, tc := range redeem.FuncCases {
		tc := tc
		for _, tag := range tc.Tags {
			if tag == "smoke" {
				RunFuncTest(tc)
				break
			}
		}
	}
})
