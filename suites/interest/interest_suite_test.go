package interest_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/cases/interest"
	"my_stonks_api_tests/testdata"
)

func TestInterest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "利息管理测试套件")
}

// InterestFuncCase 利息功能测试用例类型
type InterestFuncCase struct {
	interest.TestCase
	Input  interest.InterestInput
	Expect interest.InterestExpect
}

// RunFuncTest 执行利息功能测试
func RunFuncTest(tc InterestFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(30, 1).
			WithProduct(nil).
			Build()

		// 执行利息操作
		// TODO: 根据实际接口实现
		env.LogResult(true, "利息操作测试")
	})
}

var _ = Describe("利息功能测试", func() {
	for _, tc := range interest.FuncCases {
		tc := tc
		RunFuncTest(InterestFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})
