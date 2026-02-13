package spec_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/spec"
	"my_stonks_api_tests/testdata"
)

func TestSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "规格管理测试套件")
}

// SpecFuncCase 规格测试用例类型
type SpecFuncCase struct {
	spec.TestCase
	Input  spec.SpecInput
	Expect spec.SpecExpect
}

// RunSpecFuncTest 执行规格功能测试
func RunSpecFuncTest(tc SpecFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		env := testdata.NewTestEnv(tc.CaseID)

		switch tc.Input.Action {
		case "add":
			runAddSpecTest(env, tc)
		case "update":
			runUpdateSpecTest(env, tc)
		case "shelves":
			runShelvesSpecTest(env, tc)
		case "page":
			runPageSpecTest(env, tc)
		case "list":
			runListSpecTest(env, tc)
		case "detail":
			runDetailSpecTest(env, tc)
		default:
			runAddSpecTest(env, tc)
		}
	})
}

// runAddSpecTest 执行添加规格测试
func runAddSpecTest(env *testdata.TestEnv, tc SpecFuncCase) {
	env = env.WithCoin().Build()

	req := map[string]interface{}{
		"specValue":    tc.Input.SpecValue,
		"deadlineType": tc.Input.DeadlineType,
	}
	if tc.Input.SpecKey != "" {
		req["specKey"] = tc.Input.SpecKey
	}
	if tc.Input.Remark != "" {
		req["remark"] = tc.Input.Remark
	}
	if tc.Input.ShelvesStatus != 0 {
		req["shelvesStatus"] = tc.Input.ShelvesStatus
	}
	if len(tc.Input.LangNameList) > 0 {
		langList := make([]map[string]string, len(tc.Input.LangNameList))
		for i, lang := range tc.Input.LangNameList {
			langList[i] = map[string]string{
				"langKey": lang.LangKey,
				"content": lang.Content,
			}
		}
		req["langNameList"] = langList
	}

	resp, err := env.Client.Post(api.AdminSpecAdd, req)
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		Expect(resp.Code).To(Equal(0))
		env.LogResult(true, "添加规格成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "添加规格失败（预期）: "+resp.Message)
	}
}

// runUpdateSpecTest 执行更新规格测试
func runUpdateSpecTest(env *testdata.TestEnv, tc SpecFuncCase) {
	env = env.WithCoin().WithSpec(30, 1).Build()

	req := map[string]interface{}{
		"id": env.SpecID,
	}
	if tc.Input.SpecKey != "" {
		req["specKey"] = tc.Input.SpecKey
	}
	if tc.Input.Remark != "" {
		req["remark"] = tc.Input.Remark
	}
	if tc.Input.DeadlineType != nil {
		req["deadlineType"] = tc.Input.DeadlineType
	}
	if len(tc.Input.LangNameList) > 0 {
		langList := make([]map[string]string, len(tc.Input.LangNameList))
		for i, lang := range tc.Input.LangNameList {
			langList[i] = map[string]string{
				"langKey": lang.LangKey,
				"content": lang.Content,
			}
		}
		req["langNameList"] = langList
	}

	resp, err := env.Client.Post(api.AdminSpecUpdate, req)
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		Expect(resp.Code).To(Equal(0))
		env.LogResult(true, "更新规格成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "更新规格失败（预期）: "+resp.Message)
	}
}

// runShelvesSpecTest 执行上下架规格测试
func runShelvesSpecTest(env *testdata.TestEnv, tc SpecFuncCase) {
	env = env.WithCoin().WithSpec(30, 1).Build()

	req := map[string]interface{}{
		"id":            env.SpecID,
		"shelvesStatus": tc.Input.ShelvesStatus,
	}

	resp, err := env.Client.Post(api.AdminSpecShelves, req)
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		Expect(resp.Code).To(Equal(0))
		env.LogResult(true, "上下架成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "上下架失败（预期）: "+resp.Message)
	}
}

// runPageSpecTest 执行分页查询测试
func runPageSpecTest(env *testdata.TestEnv, tc SpecFuncCase) {
	env = env.WithCoin().WithSpec(30, 1).Build()

	req := map[string]interface{}{
		"pageIndex": tc.Input.PageIndex,
		"pageSize":  tc.Input.PageSize,
	}
	if tc.Input.SpecValue != nil {
		req["specValue"] = tc.Input.SpecValue
	}
	if tc.Input.DeadlineType != nil {
		req["deadlineType"] = tc.Input.DeadlineType
	}

	resp, err := env.Client.Post(api.AdminSpecPage, req)
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		env.LogResult(true, "分页查询成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "分页查询失败（预期）: "+resp.Message)
	}
}

// runListSpecTest 执行列表查询测试
func runListSpecTest(env *testdata.TestEnv, tc SpecFuncCase) {
	env = env.WithCoin().WithSpec(30, 1).Build()

	resp, err := env.Client.Post(api.AdminSpecList, map[string]interface{}{})
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		env.LogResult(true, "列表查询成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "列表查询失败（预期）: "+resp.Message)
	}
}

// runDetailSpecTest 执行详情查询测试
func runDetailSpecTest(env *testdata.TestEnv, tc SpecFuncCase) {
	env = env.WithCoin().WithSpec(30, 1).Build()

	req := map[string]interface{}{
		"id": env.SpecID,
	}
	if tc.Input.ID > 0 {
		req["id"] = tc.Input.ID
	}

	resp, err := env.Client.Post(api.AdminSpecDetail, req)
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		env.LogResult(true, "详情查询成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "详情查询失败（预期）: "+resp.Message)
	}
}

var _ = Describe("规格添加测试", func() {
	for _, tc := range spec.AddCases {
		tc := tc
		RunSpecFuncTest(SpecFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

var _ = Describe("规格更新测试", func() {
	for _, tc := range spec.UpdateCases {
		tc := tc
		RunSpecFuncTest(SpecFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

var _ = Describe("规格上下架测试", func() {
	for _, tc := range spec.ShelvesCases {
		tc := tc
		RunSpecFuncTest(SpecFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

var _ = Describe("规格查询测试", func() {
	for _, tc := range spec.QueryCases {
		tc := tc
		RunSpecFuncTest(SpecFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})
