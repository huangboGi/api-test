package product_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/product"
	"my_stonks_api_tests/testdata"
)

func TestProduct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "产品管理测试套件")
}

// AddFuncCase 添加产品功能测试用例类型
type AddFuncCase struct {
	product.TestCase
	Input  product.AddInput
	Expect product.AddExpect
}

// RunAddFuncTest 执行添加产品功能测试
func RunAddFuncTest(tc AddFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(getSpecValue(tc.Input), getDeadlineType(tc.Input.DeadlineType))

		if !tc.Input.SpecOff && !tc.Input.CoinOff {
			env = env.WithProduct(nil)
		}

		env = env.Build()

		// 如果需要规格下架
		if tc.Input.SpecOff {
			env.Client.Post(api.AdminSpecShelves, map[string]interface{}{
				"id":            env.SpecID,
				"shelvesStatus": 0,
			})
		}

		// 如果需要币种下架
		if tc.Input.CoinOff {
			env.Client.Post(api.AdminCoinShelves, map[string]interface{}{
				"id":      env.CoinID,
				"shelves": 0,
			})
		}

		// 执行添加产品
		resp, err := env.Client.Post(api.AdminProductAdd, buildProductRequest(tc.Input))
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expect.Success {
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "添加产品成功")
		} else {
			Expect(resp.Code).ToNot(Equal(0))
			if tc.Expect.ErrMsgContains != "" {
				Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
			}
			env.LogResult(false, "添加产品失败（预期）: "+resp.Message)
		}
	})
}

// getSpecValue 获取规格值
func getSpecValue(input product.AddInput) int {
	switch v := input.SpecValue.(type) {
	case int:
		return v
	default:
		return 30
	}
}

// getDeadlineType 获取期限类型
func getDeadlineType(v interface{}) int {
	switch t := v.(type) {
	case int:
		return t
	default:
		return 1
	}
}

// buildProductRequest 构建产品请求
func buildProductRequest(input product.AddInput) map[string]interface{} {
	req := map[string]interface{}{
		"coin":         input.Coin,
		"specValue":    input.SpecValue,
		"deadlineType": input.DeadlineType,
	}
	if input.AnnualAte.IsPositive() || input.AnnualAte.IsNegative() {
		req["annualAte"] = input.AnnualAte
	}
	if input.MinVol.IsPositive() || input.MinVol.IsNegative() {
		req["minVol"] = input.MinVol
	}
	if input.UseQuotaTotal.IsPositive() || input.UseQuotaTotal.IsNegative() {
		req["useQuotaTotal"] = input.UseQuotaTotal
	}
	if input.PersonQuota.IsPositive() || input.PersonQuota.IsNegative() {
		req["personQuota"] = input.PersonQuota
	}
	if input.Tag != "" {
		req["tag"] = input.Tag
	}
	return req
}

var _ = Describe("产品添加功能测试", func() {
	for _, tc := range product.AddCases {
		tc := tc
		RunAddFuncTest(AddFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

// UpdateFuncCase 修改产品功能测试用例类型
type UpdateFuncCase struct {
	product.TestCase
	Input  product.UpdateInput
	Expect product.UpdateExpect
}

// RunUpdateFuncTest 执行修改产品功能测试
func RunUpdateFuncTest(tc UpdateFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(-1, 0).
			WithProduct(nil).
			Build()

		// 执行修改产品
		req := map[string]interface{}{
			"id": env.ProductID,
		}
		// SpecValue 和 DeadlineType 是 interface{}，需要检查是否有实际值
		if tc.Input.SpecValue != nil {
			req["specValue"] = tc.Input.SpecValue
		}
		if tc.Input.DeadlineType != nil {
			req["deadlineType"] = tc.Input.DeadlineType
		}
		if tc.Input.ChangeCoin {
			req["coin"] = "CHANGEDCOIN"
		}
		if tc.Input.Coin != "" {
			req["coin"] = tc.Input.Coin
		}
		if tc.Input.AnnualAte.IsPositive() || tc.Input.AnnualAte.IsNegative() {
			req["annualAte"] = tc.Input.AnnualAte
		}
		if tc.Input.MinVol.IsPositive() || tc.Input.MinVol.IsNegative() {
			req["minVol"] = tc.Input.MinVol
		}

		resp, err := env.Client.Post(api.AdminProductUpdate, req)
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expect.Success {
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "修改产品成功")
		} else {
			Expect(resp.Code).ToNot(Equal(0))
			if tc.Expect.ErrMsgContains != "" {
				Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
			}
			env.LogResult(false, "修改产品失败（预期）: "+resp.Message)
		}
	})
}

var _ = Describe("产品修改功能测试", func() {
	for _, tc := range product.UpdateCases {
		tc := tc
		RunUpdateFuncTest(UpdateFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

// ShelvesFuncCase 上下架功能测试用例类型
type ShelvesFuncCase struct {
	product.TestCase
	Input  product.ShelvesInput
	Expect product.ShelvesExpect
}

// RunShelvesFuncTest 执行上下架功能测试
func RunShelvesFuncTest(tc ShelvesFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.TestCase.CaseID).
			WithCoin().
			WithSpec(-1, 0).
			WithProduct(nil).
			Build()

		// 执行上下架操作
		resp, err := env.Client.Post(api.AdminProductShelves, map[string]interface{}{
			"id":            env.ProductID,
			"shelvesStatus": tc.Input.ShelvesStatus,
		})
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expect.Success {
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "上下架操作成功")
		} else {
			Expect(resp.Code).ToNot(Equal(0))
			env.LogResult(false, "上下架操作失败（预期）: "+resp.Message)
		}
	})
}

var _ = Describe("产品上下架功能测试", func() {
	for _, tc := range product.ShelvesCases {
		tc := tc
		RunShelvesFuncTest(ShelvesFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

// PageFuncCase 分页查询功能测试用例类型
type PageFuncCase struct {
	product.TestCase
	Input  product.QueryInput
	Expect product.QueryExpect
}

// RunPageFuncTest 执行分页查询功能测试
func RunPageFuncTest(tc PageFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(-1, 0).
			WithProduct(nil).
			Build()

		// 执行分页查询
		req := map[string]interface{}{
			"pageIndex": tc.Input.PageIndex,
			"pageSize":  tc.Input.PageSize,
		}
		if tc.Input.Name != "" {
			req["name"] = tc.Input.Name
		}
		if tc.Input.SpecValue != 0 {
			req["specValue"] = tc.Input.SpecValue
		}
		if tc.Input.DeadlineType != 0 {
			req["deadlineType"] = tc.Input.DeadlineType
		}

		resp, err := env.Client.Post(api.AdminProductPage, req)
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expect.Success {
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "分页查询成功")
		} else {
			Expect(resp.Code).ToNot(Equal(0))
			if tc.Expect.ErrMsgContains != "" {
				Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
			}
			env.LogResult(false, "分页查询失败（预期）: "+resp.Message)
		}
	})
}

var _ = Describe("产品分页查询功能测试", func() {
	for _, tc := range product.PageCases {
		tc := tc
		RunPageFuncTest(PageFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

// DetailFuncCase 详情查询功能测试用例类型
type DetailFuncCase struct {
	product.TestCase
	Input  product.DetailInput
	Expect product.DetailExpect
}

// RunAdminDetailFuncTest 执行管理端详情查询功能测试
func RunAdminDetailFuncTest(tc DetailFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(-1, 0).
			WithProduct(nil).
			Build()

		// 确定产品ID
		productID := env.ProductID
		if tc.Input.NotExist {
			productID = int64(tc.Input.ID)
		}

		// 执行详情查询
		resp, err := env.Client.Post(api.AdminProductDetail, map[string]interface{}{
			"id": productID,
		})
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expect.Success {
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "获取详情成功")
		} else {
			Expect(resp.Code).ToNot(Equal(0))
			if tc.Expect.ErrMsgContains != "" {
				Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
			}
			env.LogResult(false, "获取详情失败（预期）: "+resp.Message)
		}
	})
}

var _ = Describe("产品详情查询功能测试（管理端）", func() {
	for _, tc := range product.AdminDetailCases {
		tc := tc
		RunAdminDetailFuncTest(DetailFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

// ListFuncCase 列表查询功能测试用例类型
type ListFuncCase struct {
	product.TestCase
	Input  product.QueryInput
	Expect product.QueryExpect
}

// RunListFuncTest 执行用户端列表查询功能测试
func RunListFuncTest(tc ListFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(-1, 0).
			WithProduct(nil).
			Build()

		// 上架产品
		env.Client.Post(api.AdminProductShelves, map[string]interface{}{
			"id":            env.ProductID,
			"shelvesStatus": 1,
		})

		// 执行列表查询
		req := map[string]interface{}{}
		if tc.Input.Name != "" {
			req["name"] = tc.Input.Name
		}
		if tc.Input.SpecValue != 0 {
			req["specValue"] = tc.Input.SpecValue
		}
		if tc.Input.DeadlineType != 0 {
			req["deadlineType"] = tc.Input.DeadlineType
		}

		resp, err := env.Client.Post(api.AppProductList, req)
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expect.Success {
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "列表查询成功")
		} else {
			Expect(resp.Code).ToNot(Equal(0))
			if tc.Expect.ErrMsgContains != "" {
				Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
			}
			env.LogResult(false, "列表查询失败（预期）: "+resp.Message)
		}
	})
}

var _ = Describe("产品列表查询功能测试（用户端）", func() {
	for _, tc := range product.ListCases {
		tc := tc
		RunListFuncTest(ListFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

// RunAppDetailFuncTest 执行用户端详情查询功能测试
func RunAppDetailFuncTest(tc DetailFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		// 准备环境
		env := testdata.NewTestEnv(tc.CaseID).
			WithCoin().
			WithSpec(-1, 0).
			WithProduct(nil).
			Build()

		// 确定产品ID和上架状态
		productID := env.ProductID
		if tc.Input.NotExist {
			productID = int64(tc.Input.ID)
		}
		if !tc.Input.NotOnShelf {
			// 上架产品
			env.Client.Post(api.AdminProductShelves, map[string]interface{}{
				"id":            env.ProductID,
				"shelvesStatus": 1,
			})
		}

		// 执行详情查询
		resp, err := env.Client.Post(api.AppProductDetail, map[string]interface{}{
			"id": productID,
		})
		Expect(err).ToNot(HaveOccurred())

		// 验证结果
		if tc.Expect.Success {
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.Code).To(Equal(0))
			env.LogResult(true, "获取详情成功")
		} else {
			Expect(resp.Code).ToNot(Equal(0))
			if tc.Expect.ErrMsgContains != "" {
				Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
			}
			env.LogResult(false, "获取详情失败（预期）: "+resp.Message)
		}
	})
}

var _ = Describe("产品详情查询功能测试（用户端）", func() {
	for _, tc := range product.AppDetailCases {
		tc := tc
		RunAppDetailFuncTest(DetailFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})
