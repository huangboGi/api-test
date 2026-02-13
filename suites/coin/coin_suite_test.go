package coin_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/cases/coin"
	"my_stonks_api_tests/models"
	"my_stonks_api_tests/testdata"
)

func TestCoin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "币种配置测试套件")
}

// CoinFuncCase 币种功能测试用例类型
type CoinFuncCase struct {
	coin.TestCase
	Input  coin.CoinInput
	Expect coin.CoinExpect
}

// RunFuncTest 执行币种功能测试
func RunFuncTest(tc CoinFuncCase) {
	It(tc.Title, Label(tc.CaseID), func() {
		By("【用例编号】" + tc.CaseID)
		By("【优先级】" + tc.Priority)

		env := testdata.NewTestEnv(tc.CaseID)

		switch tc.Input.Action {
		case "add":
			runAddCoinTest(env, tc)
		case "update":
			runUpdateCoinTest(env, tc)
		case "shelves":
			runShelvesCoinTest(env, tc)
		case "page":
			runPageCoinTest(env, tc)
		case "detail":
			runDetailCoinTest(env, tc)
		case "selectCoin":
			runSelectCoinTest(env, tc)
		default:
			runAddCoinTest(env, tc)
		}
	})
}

// runAddCoinTest 执行添加币种测试
func runAddCoinTest(env *testdata.TestEnv, tc CoinFuncCase) {
	// 生成唯一的币种代码（成功场景始终生成唯一值，避免重复）
	var coinCode string
	if tc.Expect.Success {
		coinCode = testdata.GenerateUniqueCoin()
	} else {
		coinCode = tc.Input.Coin
		if coinCode == "" {
			coinCode = testdata.GenerateUniqueCoin()
		}
	}

	// CoinKey处理：只有用例明确指定时才传入，否则不传（让后端自动使用Coin值）
	var coinKey string
	if tc.Input.CoinKey != "" {
		coinKey = tc.Input.CoinKey
	}

	// 构建多语言列表
	langList := make([]map[string]string, len(tc.Input.LangNameList))
	for i, lang := range tc.Input.LangNameList {
		langList[i] = map[string]string{
			"langKey": lang.LangKey,
			"content": lang.Content,
		}
	}

	// 执行添加币种
	req := testdata.NewCoinConfig(coinCode, coinKey, tc.Input.Tag, langList)
	resp, err := env.Client.Post(api.AdminCoinAdd, req)
	Expect(err).ToNot(HaveOccurred())

	// 验证结果
	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		Expect(resp.Code).To(Equal(0))

		// 验证数据库
		if tc.Expect.DBCheck.RecordCreated {
			var coinConfig models.WthCoinConfig
			err := env.DB.Query(&coinConfig, "coin = ?", coinCode)
			Expect(err).ToNot(HaveOccurred())
			env.CoinID = coinConfig.ID

			// 验证主表字段值与传入参数一致
			Expect(coinConfig.Coin).To(Equal(coinCode), "Coin字段值应与传入参数一致")
			Expect(coinConfig.Tag).To(Equal(tc.Input.Tag), "Tag字段值应与传入参数一致")

			// 验证CoinKey
			// - 若用例明确指定了CoinKey，则数据库应使用指定的值
			// - 若用例未指定CoinKey（空字符串），则后端默认使用Coin值
			if tc.Input.CoinKey != "" {
				Expect(coinConfig.CoinKey).To(Equal(coinKey), "CoinKey应与传入的自定义值一致")
			} else {
				Expect(coinConfig.CoinKey).To(Equal(coinCode), "未传入CoinKey时后端应默认使用Coin值")
			}

			// 验证多语言数据是否正确入库
			if tc.Expect.DBCheck.LangDataCheck.ShouldExist && len(tc.Input.LangNameList) > 0 {
				// 确定用于关联的configKey（使用数据库中实际存储的CoinKey）
				configKey := coinConfig.CoinKey
				if configKey == "" {
					configKey = coinConfig.Coin
				}

				// 查询多语言数据
				var langData []models.ConfigLanguage
				err := env.DB.Where("config_key = ? AND type = ?", configKey, models.LanguageTypeCoin).Find(&langData).Error
				Expect(err).ToNot(HaveOccurred())

				// 验证多语言数据数量
				if tc.Expect.DBCheck.LangDataCheck.ExpectedCount > 0 {
					Expect(len(langData)).To(Equal(tc.Expect.DBCheck.LangDataCheck.ExpectedCount), "多语言数据数量应与预期一致")
				} else {
					Expect(len(langData)).To(Equal(len(tc.Input.LangNameList)), "多语言数据数量应与传入参数一致")
				}

				// 验证多语言内容
				langMap := make(map[string]string)
				for _, item := range langData {
					langMap[item.LangKey] = item.Content
				}
				for _, expected := range tc.Input.LangNameList {
					actual, exists := langMap[expected.LangKey]
					Expect(exists).To(BeTrue(), "应存在语言 %s 的多语言数据", expected.LangKey)
					Expect(actual).To(Equal(expected.Content), "语言 %s 的内容应与传入参数一致", expected.LangKey)
				}
			}
		}
		env.LogResult(true, "添加币种成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "添加币种失败（预期）: "+resp.Message)
	}
}

// runUpdateCoinTest 执行更新币种测试
func runUpdateCoinTest(env *testdata.TestEnv, tc CoinFuncCase) {
	// 如果是指定不存在的ID
	if tc.Input.NotExist {
		req := map[string]interface{}{
			"id":   tc.Input.ID,
			"coin": tc.Input.Coin,
		}
		resp, err := env.Client.Post(api.AdminCoinUpdate, req)
		Expect(err).ToNot(HaveOccurred())

		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "更新失败（预期）: "+resp.Message)
		return
	}

	// 先创建币种
	coinCode := tc.Input.Coin
	if coinCode == "" || coinCode == "USDT" {
		coinCode = testdata.GenerateUniqueCoin()
	}
	coinKey := testdata.GenerateUniqueCoinKey()
	req := testdata.NewCoinConfig(coinCode, coinKey, "更新测试", nil)
	resp, err := env.Client.Post(api.AdminCoinAdd, req)
	Expect(err).ToNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))

	// 查询币种ID
	var coinConfig models.WthCoinConfig
	err = env.DB.Query(&coinConfig, "coin = ?", coinCode)
	Expect(err).ToNot(HaveOccurred())

	// 构建更新请求
	updateReq := map[string]interface{}{
		"id":   coinConfig.ID,
		"coin": coinCode,
	}
	if tc.Input.Tag != "" {
		updateReq["tag"] = tc.Input.Tag
	}
	if tc.Input.CoinKey != "" {
		updateReq["coinKey"] = tc.Input.CoinKey
	}
	if len(tc.Input.LangNameList) > 0 {
		langList := make([]map[string]string, len(tc.Input.LangNameList))
		for i, lang := range tc.Input.LangNameList {
			langList[i] = map[string]string{
				"langKey": lang.LangKey,
				"content": lang.Content,
			}
		}
		updateReq["langNameList"] = langList
	}

	// 执行更新
	resp, err = env.Client.Post(api.AdminCoinUpdate, updateReq)
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		Expect(resp.Code).To(Equal(0))
		env.LogResult(true, "更新成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "更新失败（预期）: "+resp.Message)
	}
}

// runShelvesCoinTest 执行上下架测试
func runShelvesCoinTest(env *testdata.TestEnv, tc CoinFuncCase) {
	// 如果是指定不存在的ID
	if tc.Input.NotExist {
		req := testdata.NewCoinShelvesRequest(tc.Input.ID, tc.Input.Shelves)
		resp, err := env.Client.Post(api.AdminCoinShelves, req)
		Expect(err).ToNot(HaveOccurred())

		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "上下架失败（预期）: "+resp.Message)
		return
	}

	// 先创建币种
	coinCode := testdata.GenerateUniqueCoin()
	coinKey := testdata.GenerateUniqueCoinKey()
	req := testdata.NewCoinConfig(coinCode, coinKey, "上下架测试", nil)
	resp, err := env.Client.Post(api.AdminCoinAdd, req)
	Expect(err).ToNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))

	// 查询币种ID
	var coinConfig models.WthCoinConfig
	err = env.DB.Query(&coinConfig, "coin = ?", coinCode)
	Expect(err).ToNot(HaveOccurred())

	// 执行上下架
	shelvesReq := testdata.NewCoinShelvesRequest(coinConfig.ID, tc.Input.Shelves)
	resp, err = env.Client.Post(api.AdminCoinShelves, shelvesReq)
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		Expect(resp.Code).To(Equal(0))

		// 验证数据库状态
		if tc.Expect.DBCheck.ShelvesMatch != 0 || tc.Input.Shelves == 0 {
			env.DB.Query(&coinConfig, "coin = ?", coinCode)
			Expect(int(coinConfig.Shelves)).To(Equal(tc.Input.Shelves))
		}
		env.LogResult(true, "上下架成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "上下架失败（预期）: "+resp.Message)
	}
}

// runPageCoinTest 执行分页查询测试
func runPageCoinTest(env *testdata.TestEnv, tc CoinFuncCase) {
	env = env.WithCoin().Build()

	req := map[string]interface{}{
		"pageIndex": tc.Input.PageIndex,
		"pageSize":  tc.Input.PageSize,
	}
	if tc.Input.Coin != "" {
		req["coin"] = tc.Input.Coin
	}
	if tc.Input.CoinName != "" {
		req["coinName"] = tc.Input.CoinName
	}
	if tc.Input.Lang != "" {
		req["lang"] = tc.Input.Lang
	}

	resp, err := env.Client.Post(api.AdminCoinPage, req)
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

// runDetailCoinTest 执行详情查询测试
func runDetailCoinTest(env *testdata.TestEnv, tc CoinFuncCase) {
	env = env.WithCoin().Build()

	req := map[string]interface{}{}
	if tc.Input.ID > 0 {
		req["id"] = tc.Input.ID
	} else if !tc.Input.NotExist {
		req["id"] = env.CoinID
	}

	resp, err := env.Client.Post(api.AdminCoinDetail, req)
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

// runSelectCoinTest 执行可用币种查询测试
func runSelectCoinTest(env *testdata.TestEnv, tc CoinFuncCase) {
	env = env.WithCoin().Build()

	resp, err := env.Client.Post(api.AdminCoinSelectCoin, map[string]interface{}{})
	Expect(err).ToNot(HaveOccurred())

	if tc.Expect.Success {
		Expect(resp.StatusCode).To(Equal(tc.Expect.StatusCode))
		env.LogResult(true, "查询可用币种成功")
	} else {
		if tc.Expect.ErrMsgContains != "" {
			Expect(resp.Message).To(ContainSubstring(tc.Expect.ErrMsgContains))
		}
		env.LogResult(false, "查询可用币种失败（预期）: "+resp.Message)
	}
}

var _ = Describe("币种添加测试", func() {
	for _, tc := range coin.AddCases {
		tc := tc
		RunFuncTest(CoinFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

var _ = Describe("币种更新测试", func() {
	for _, tc := range coin.UpdateCases {
		tc := tc
		RunFuncTest(CoinFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

var _ = Describe("币种上下架测试", func() {
	for _, tc := range coin.ShelvesCases {
		tc := tc
		RunFuncTest(CoinFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})

var _ = Describe("币种查询测试", func() {
	for _, tc := range coin.QueryCases {
		tc := tc
		RunFuncTest(CoinFuncCase{
			TestCase: tc.TestCase,
			Input:    tc.Input,
			Expect:   tc.Expect,
		})
	}
})
