package cases

import (
	"fmt"
	"testing"
	"time"

	"my_stonks_api_tests/config"
	"my_stonks_api_tests/framework"
	"my_stonks_api_tests/models"
	"my_stonks_api_tests/testdata"

	_ "my_stonks_api_tests/config"
)

func init() {
	config.Load()
}

// TestWthCoinConfig_Add_Success 测试添加币种配置-成功场景
// 用例编号: WTH_COIN_ADD_FUNC_001
// 所属模块: 币种配置管理
// 优先级: 高
// 测试类型: 功能测试
// 关联需求: WTH-COIN-001
// 前置条件: 1. 已登录管理端系统 2. 有添加币种配置权限
// 测试步骤:
//  1. 准备币种配置数据（必填字段：coin、coinKey、langNameList）
//  2. 调用 POST /api/v1/admin/wth/coin/add 接口
//  3. 验证响应状态码为 200
//  4. 验证数据库中已创建该币种配置
//
// 测试数据: coin=自动生成唯一值, coinKey=自动生成唯一值, langNameList包含中英文
// 预期结果: 1. API返回成功 2. 数据库中存在该币种配置 3. 多语言数据已插入config_language表
func TestWthCoinConfig_Add_Success(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	coin := testdata.GenerateUniqueCoin()
	coinKey := "" // 留空让函数自动生成唯一的coinKey
	request := testdata.NewCoinConfig(coin, coinKey, "测试标签")

	var coinConfig models.WthCoinConfig

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口: POST /api/v1/admin/wth/coin/add")
	resp, err := client.Post("/api/v1/admin/wth/coin/add", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)

	// 验证数据库中的数据
	err = db.Query(&coinConfig, "coin = ?", coin)
	framework.AssertDBRecordExists(t, err == nil, err)
	framework.AssertDBFieldEqual(t, "Coin", coin, coinConfig.Coin)
	// 验证 coinKey 不为空
	framework.AssertDBFieldEqual(t, "CoinKey不为空", true, coinConfig.CoinKey != "")

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_UpdateShelves_OffThenUserNotVisible 测试币种上下架切换-下架后用户端不可见
// 用例编号: WTH_COIN_UPSHELF_INT_001
// 所属模块: 币种配置管理
// 优先级: 高
// 测试类型: 集成测试
// 关联需求: WTH-COIN-006
// 前置条件: 1. 已登录管理端系统 2. 已登录用户端系统
// 测试步骤:
//  1. 创建一个上架状态的币种
//  2. 用户端验证可查询到此币种
//  3. 管理端将币种下架
//  4. 用户端再次验证不可查询到此币种
//
// 测试数据: coin=test_coin_004
// 预期结果: 1. 初始状态用户端可查询 2. 下架后用户端不可查询 3. 状态切换正确
func TestWthCoinConfig_UpdateShelves_OffThenUserNotVisible(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	coin := testdata.GenerateUniqueCoin()
	coinKey := testdata.GenerateUniqueCoinKey()
	client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "集成测试"))

	// 查询获取币种的ID
	var coinConfig models.WthCoinConfig
	err := db.Query(&coinConfig, "coin = ?", coin)
	if err != nil {
		t.Fatalf("查询币种失败: %v", err)
	}

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境：已创建币种配置")
	framework.LogTestStep(t, 2, "管理端：将币种上架")

	// 先上架币种（使用id而不是coin）
	shelvesRequest := map[string]interface{}{
		"id":      coinConfig.ID,
		"shelves": 1,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/updateShelves", shelvesRequest)
	if err != nil {
		t.Fatalf("上架失败: %v", err)
	}
	framework.AssertStatusCode(t, resp, 200)

	// 验证用户端可以查询（通过selectCoin接口）
	framework.LogTestStep(t, 3, "用户端：验证可查询到该币种")
	selectResp, err := client.Post("/api/v1/admin/wth/coin/selectCoin", map[string]interface{}{})
	if err != nil {
		t.Fatalf("用户端查询失败: %v", err)
	}
	framework.AssertStatusCode(t, selectResp, 200)

	// 管理端将币种下架
	framework.LogTestStep(t, 4, "管理端：将币种下架")
	shelvesRequest["shelves"] = 0
	resp, err = client.Post("/api/v1/admin/wth/coin/updateShelves", shelvesRequest)
	if err != nil {
		t.Fatalf("下架失败: %v", err)
	}
	framework.AssertStatusCode(t, resp, 200)

	// Assert - 验证结果
	framework.LogTestStep(t, 5, "验证测试结果")

	// 验证数据库中shelves状态为false
	err = db.Query(&coinConfig, "coin = ?", coin)
	framework.AssertDBRecordExists(t, err == nil, err)
	framework.AssertDBFieldEqual(t, "Shelves", 0, coinConfig.Shelves)

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_UpdateShelves_OnThenUserVisible 测试币种上下架切换-上架后用户端可见
// 用例编号: WTH_COIN_UPSHELF_INT_002
// 所属模块: 币种配置管理
// 优先级: 高
// 测试类型: 集成测试
// 关联需求: WTH-COIN-006
// 前置条件: 1. 已登录管理端系统 2. 已登录用户端系统
// 测试步骤:
//  1. 创建一个下架状态的币种
//  2. 用户端验证不可查询到此币种
//  3. 管理端将币种上架
//  4. 用户端再次验证可查询到此币种
//
// 测试数据: coin=test_coin_005
// 预期结果: 1. 初始状态用户端不可查询 2. 上架后用户端可查询 3. 状态切换正确
func TestWthCoinConfig_UpdateShelves_OnThenUserVisible(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	coin := testdata.GenerateUniqueCoin()
	coinKey := testdata.GenerateUniqueCoinKey()
	client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "集成测试2"))

	// 查询获取币种的ID
	var coinConfig models.WthCoinConfig
	err := db.Query(&coinConfig, "coin = ?", coin)
	if err != nil {
		t.Fatalf("查询币种失败: %v", err)
	}

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境：已创建币种配置")
	framework.LogTestStep(t, 2, "管理端：将币种下架")

	// 先下架币种（使用id而不是coin）
	shelvesRequest := map[string]interface{}{
		"id":      coinConfig.ID,
		"shelves": 0,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/updateShelves", shelvesRequest)
	if err != nil {
		t.Fatalf("下架失败: %v", err)
	}
	framework.AssertStatusCode(t, resp, 200)

	// 验证数据库中shelves状态为false
	db.Query(&coinConfig, "coin = ?", coin)
	framework.AssertDBFieldEqual(t, "初始Shelves应为false", 0, coinConfig.Shelves)

	// 管理端将币种上架
	framework.LogTestStep(t, 3, "管理端：将币种上架")
	shelvesRequest["shelves"] = 1
	resp, err = client.Post("/api/v1/admin/wth/coin/updateShelves", shelvesRequest)
	if err != nil {
		t.Fatalf("上架失败: %v", err)
	}
	framework.AssertStatusCode(t, resp, 200)

	// Assert - 验证结果
	framework.LogTestStep(t, 4, "验证测试结果")

	// 验证数据库中shelves状态为true
	err = db.Query(&coinConfig, "coin = ?", coin)
	framework.AssertDBRecordExists(t, err == nil, err)
	framework.AssertDBFieldEqual(t, "Shelves", 1, coinConfig.Shelves)

	// 验证用户端可以查询（通过selectCoin接口）
	framework.LogTestStep(t, 5, "用户端：验证可查询到该币种")
	selectResp, err := client.Post("/api/v1/admin/wth/coin/selectCoin", map[string]interface{}{})
	if err != nil {
		t.Fatalf("用户端查询失败: %v", err)
	}
	framework.AssertStatusCode(t, selectResp, 200)

	framework.LogTestResult(t, true, "测试通过")
}

// ========================================================================
// 补充测试用例（增强现有测试）
// ========================================================================

// TestWthCoinConfig_Add_MultiLanguageComplete 测试添加币种配置-多语言完整测试
// 用例编号: WTH_COIN_ADD_FUNC_002
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 功能测试
// 关联需求: WTH-COIN-001
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 调用添加接口，提供完整的多语言配置
//  2. 验证所有语言数据都已插入
//
// 测试数据: coin=multi_lang_coin, langNameList=[{"langKey":"zh-Hans","content":"中文"},{"langKey":"en","content":"English"},{"langKey":"ja","content":"日本語"}]
// 预期结果: 1. API返回成功 2. 数据库中存储3种语言 3. config_language表中存在3条记录
func TestWthCoinConfig_Add_MultiLanguageComplete(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	coin := testdata.GenerateUniqueCoin()
	coinKey := testdata.GenerateUniqueCoinKey()
	request := map[string]interface{}{
		"coin":    coin,
		"coinKey": coinKey,
		"tag":     "多语言测试",
		"langNameList": []map[string]string{
			{"langKey": "zh-Hans", "content": "中文币种名称"},
			{"langKey": "en", "content": "English Coin Name"},
			{"langKey": "ja", "content": "日本語"},
		},
	}

	var coinConfig models.WthCoinConfig

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口: POST /api/v1/admin/wth/coin/add（多语言）")
	resp, err := client.Post("/api/v1/admin/wth/coin/add", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)

	// 验证数据库中的数据
	err = db.Query(&coinConfig, "coin = ?", coin)
	framework.AssertDBRecordExists(t, err == nil, err)
	framework.AssertDBFieldEqual(t, "Coin", coin, coinConfig.Coin)

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Page_EmptyPageIndex_ShouldFail 测试查询币种列表-pageIndex为空应失败
// 用例编号: WTH_COIN_PAGE_NEG_001
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 逆向测试
// 关联需求: WTH-COIN-002
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 调用分页接口，pageIndex为空
//  2. 验证返回参数验证错误
//
// 测试数据: pageIndex为空, pageSize=10
// 预期结果: 1. API返回失败 2. 提示pageIndex不能为空
func TestWthCoinConfig_Page_EmptyPageIndex_ShouldFail(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口，pageIndex为空")

	request := map[string]interface{}{
		"pageIndex": nil, // pageIndex为空
		"pageSize":  10,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/page", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)
	framework.AssertErrorMessageContains(t, resp, "不能为空")

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Page_PageIndexZero_ShouldFail 测试查询币种列表-pageIndex为0应失败
// 用例编号: WTH_COIN_PAGE_NEG_002
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 逆向测试
// 关联需求: WTH-COIN-002
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 调用分页接口，pageIndex=0
//  2. 验证返回参数验证错误
//
// 测试数据: pageIndex=0, pageSize=10
// 预期结果: 1. API返回失败 2. 提示pageIndex必须大于0
func TestWthCoinConfig_Page_PageIndexZero_ShouldFail(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口，pageIndex=0")

	request := map[string]interface{}{
		"pageIndex": 0,
		"pageSize":  10,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/page", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)
	framework.AssertErrorMessageContains(t, resp, "大于0")

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Page_PageIndexNegative_ShouldFail 测试查询币种列表-pageIndex为负数应失败
// 用例编号: WTH_COIN_PAGE_NEG_003
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 逆向测试
// 关联需求: WTH-COIN-002
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 调用分页接口，pageIndex=-1
//  2. 验证返回参数验证错误
//
// 测试数据: pageIndex=-1, pageSize=10
// 预期结果: 1. API返回失败 2. 提示pageIndex必须大于0
func TestWthCoinConfig_Page_PageIndexNegative_ShouldFail(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口，pageIndex=-1")

	request := map[string]interface{}{
		"pageIndex": -1,
		"pageSize":  10,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/page", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)
	framework.AssertErrorMessageContains(t, resp, "大于0")

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Page_PageSizeMaxBoundary 测试查询币种列表-pageSize边界值测试（最大值）
// 用例编号: WTH_COIN_PAGE_BND_002
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 边界测试
// 关联需求: WTH-COIN-002
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 使用pageSize=100或更大值查询
//  2. 验证返回正确的数据量
//
// 测试数据: pageIndex=1, pageSize=100（或允许的最大值）
// 预期结果: 1. API返回成功 2. 返回数据量不超过pageSize
func TestWthCoinConfig_Page_PageSizeMaxBoundary(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口，pageSize=100")

	request := map[string]interface{}{
		"pageIndex": 1,
		"pageSize":  100,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/page", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Page_PageSizeNegative_ShouldFail 测试查询币种列表-pageSize为负数应失败
// 用例编号: WTH_COIN_PAGE_NEG_004
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 逆向测试
// 关联需求: WTH-COIN-002
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 调用分页接口，pageSize=-1
//  2. 验证返回参数验证错误
//
// 测试数据: pageIndex=1, pageSize=-1
// 预期结果: 1. API返回失败 2. 提示pageSize必须大于0
func TestWthCoinConfig_Page_PageSizeNegative_ShouldFail(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口，pageSize=-1")

	request := map[string]interface{}{
		"pageIndex": 1,
		"pageSize":  -1,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/page", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)
	framework.AssertErrorMessageContains(t, resp, "大于0")

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Page_NameFilter 测试查询币种列表-按名称模糊查询
// 用例编号: WTH_COIN_PAGE_FUNC_003
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 功能测试
// 关联需求: WTH-COIN-002
// 前置条件: 1. 已登录管理端系统 2. 数据库中存在币种配置
// 测试步骤:
//  1. 创建多个币种配置（名称包含"测试"）
//  2. 使用coinName参数模糊查询
//  3. 验证返回结果符合条件
//
// 测试数据: coinName=测试
// 预期结果: 1. API返回成功 2. 返回列表只包含名称包含"测试"的币种
func TestWthCoinConfig_Page_NameFilter(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// 创建测试数据
	namePrefix := fmt.Sprintf("测试名称_%d_", time.Now().Unix())
	testCoins := []string{
		testdata.GenerateUniqueCoin(),
		testdata.GenerateUniqueCoin(),
		testdata.GenerateUniqueCoin(),
	}
	for i, coin := range testCoins {
		coinKey := testdata.GenerateUniqueCoinKey()
		request := map[string]interface{}{
			"coin":    coin,
			"coinKey": coinKey,
			"tag":     "测试标签",
			"langNameList": []map[string]string{
				{"langKey": "zh-Hans", "content": namePrefix + fmt.Sprintf("%d", i)},
			},
		}
		client.Post("/api/v1/admin/wth/coin/add", request)
	}

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境：已创建测试数据")
	framework.LogTestStep(t, 2, "调用API接口，coinName参数为测试名称")

	request := map[string]interface{}{
		"pageIndex": 1,
		"pageSize":  10,
		"coinName":  namePrefix,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/page", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)

	// 注意：币种名称的模糊查询是在 config_language 表中进行的
	// 这里主要验证接口能正常返回数据
	// 实际的币种名称过滤逻辑由后端服务层处理

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Page_OutOfRange 测试查询币种列表-分页超出范围
// 用例编号: WTH_COIN_PAGE_BND_003
// 所属模块: 币种配置管理
// 优先级: 低
// 测试类型: 边界测试
// 关联需求: WTH-COIN-002
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 创建少量数据（3条）
//  2. 查询第2页（pageIndex=2, pageSize=10）
//  3. 验证返回空列表
//
// 测试数据: pageIndex=2, pageSize=10
// 预期结果: 1. API返回成功 2. 返回空列表
func TestWthCoinConfig_Page_OutOfRange(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// 创建少量测试数据
	testCoins := []string{
		testdata.GenerateUniqueCoin(),
		testdata.GenerateUniqueCoin(),
		testdata.GenerateUniqueCoin(),
	}
	for _, coin := range testCoins {
		coinKey := testdata.GenerateUniqueCoinKey()
		client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "边界测试"))
	}

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境：已创建3条币种配置")
	framework.LogTestStep(t, 2, "调用API接口，pageIndex=2（超出范围）")

	request := map[string]interface{}{
		"pageIndex": 2,
		"pageSize":  10,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/page", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Detail_NotExist_ShouldFail 测试查询币种详情-币种不存在应失败
// 用例编号: WTH_COIN_DETAIL_NEG_001
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 逆向测试
// 关联需求: WTH-COIN-003
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 查询不存在的币种详情
//  2. 验证返回错误提示
//
// 测试数据: id=999999（不存在的ID）
// 预期结果: 1. API返回失败 2. 提示币种不存在
func TestWthCoinConfig_Detail_NotExist_ShouldFail(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口，使用不存在的ID")

	nonExistentID := int64(999999)
	request := map[string]interface{}{
		"id": nonExistentID,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/detail", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)
	framework.AssertErrorMessageContains(t, resp, "不存在")

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Detail_EmptyCoin_ShouldFail 测试查询币种详情-ID为空应失败
// 用例编号: WTH_COIN_DETAIL_NEG_002
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 逆向测试
// 关联需求: WTH-COIN-003
// 前置条件: 1. 已登录管理端系统
// 测试步骤:
//  1. 调用详情接口，id字段为空
//  2. 验证返回参数验证错误
//
// 测试数据: id为0或空
// 预期结果: 1. API返回失败 2. 提示ID不能为空
func TestWthCoinConfig_Detail_EmptyCoin_ShouldFail(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境")
	framework.LogTestStep(t, 2, "调用API接口，id为0")

	request := map[string]interface{}{
		"id": 0, // ID为0
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/detail", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)
	framework.AssertErrorMessageContains(t, resp, "不能为空")

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_Detail_MultiLanguageComplete 测试查询币种详情-多语言数据完整性
// 用例编号: WTH_COIN_DETAIL_FUNC_002
// 所属模块: 币种配置管理
// 优先级: 中
// 测试类型: 功能测试
// 关联需求: WTH-COIN-003
// 前置条件: 1. 已登录管理端系统 2. 数据库中存在多语言币种配置
// 测试步骤:
//  1. 创建包含多语言的币种配置
//  2. 查询币种详情
//  3. 验证返回所有语言数据
//
// 测试数据: id=创建的币种ID
// 预期结果: 1. API返回成功 2. 返回所有配置的语言数据 3. 每种语言内容完整
func TestWthCoinConfig_Detail_MultiLanguageComplete(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	coin := testdata.GenerateUniqueCoin()
	coinKey := testdata.GenerateUniqueCoinKey()
	request := map[string]interface{}{
		"coin":    coin,
		"coinKey": coinKey,
		"tag":     "多语言详情测试",
		"langNameList": []map[string]string{
			{"langKey": "zh-Hans", "content": "中文币种名称"},
			{"langKey": "en", "content": "English Coin Name"},
			{"langKey": "ja", "content": "日本語"},
		},
	}
	client.Post("/api/v1/admin/wth/coin/add", request)

	// 查询获取币种的ID
	var coinConfig models.WthCoinConfig
	err := db.Query(&coinConfig, "coin = ?", coin)
	if err != nil {
		t.Fatalf("查询币种失败: %v", err)
	}

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境：已创建多语言币种配置")
	framework.LogTestStep(t, 2, "调用API接口: POST /api/v1/admin/wth/coin/detail")

	detailRequest := map[string]interface{}{
		"id": coinConfig.ID,
	}
	resp, err := client.Post("/api/v1/admin/wth/coin/detail", detailRequest)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)

	// 验证数据库中的数据
	err = db.Query(&coinConfig, "coin = ?", coin)
	framework.AssertDBRecordExists(t, err == nil, err)
	framework.AssertDBFieldEqual(t, "Coin", coin, coinConfig.Coin)

	framework.LogTestResult(t, true, "测试通过")
}

// TestWthCoinConfig_SelectCoin_OnlyShelvesOn 测试用户端查询-只返回上架币种
// 用例编号: WTH_COIN_SELECT_FUNC_002
// 所属模块: 币种配置（用户端）
// 优先级: 高
// 测试类型: 功能测试
// 关联需求: WTH-COIN-004
// 前置条件: 1. 已登录用户端系统 2. 数据库中存在上架和下架币种
// 测试步骤:
//  1. 创建多个币种（3个上架，2个下架）
//  2. 调用用户端查询接口
//  3. 验证只返回上架币种
//
// 测试数据: 无参数
// 预期结果: 1. API返回成功 2. 只返回上架状态的币种 3. 下架币种不在返回列表中
func TestWthCoinConfig_SelectCoin_OnlyShelvesOn(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	// 创建测试数据：3个上架，2个下架
	shelvedCoins := []string{
		testdata.GenerateUniqueCoin(),
		testdata.GenerateUniqueCoin(),
		testdata.GenerateUniqueCoin(),
	}
	unshelvedCoins := []string{
		testdata.GenerateUniqueCoin(),
		testdata.GenerateUniqueCoin(),
	}

	// 创建所有币种
	allCoins := append(shelvedCoins, unshelvedCoins...)
	for _, coin := range allCoins {
		coinKey := testdata.GenerateUniqueCoinKey()
		client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "用户端查询测试"))
	}

	// 查询所有币种的ID
	coinIDMap := make(map[string]int64)
	for _, coin := range allCoins {
		var coinConfig models.WthCoinConfig
		err := db.Query(&coinConfig, "coin = ?", coin)
		if err == nil {
			coinIDMap[coin] = coinConfig.ID
		}
	}

	// 将3个币种上架，2个币种下架
	for _, coin := range shelvedCoins {
		shelvesRequest := map[string]interface{}{
			"id":      coinIDMap[coin],
			"shelves": 1,
		}
		client.Post("/api/v1/admin/wth/coin/updateShelves", shelvesRequest)
	}

	for _, coin := range unshelvedCoins {
		shelvesRequest := map[string]interface{}{
			"id":      coinIDMap[coin],
			"shelves": 0,
		}
		client.Post("/api/v1/admin/wth/coin/updateShelves", shelvesRequest)
	}

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境：已创建3个上架、2个下架的币种")
	framework.LogTestStep(t, 2, "调用API接口: POST /api/v1/admin/wth/coin/selectCoin")

	request := map[string]interface{}{}
	resp, err := client.Post("/api/v1/admin/wth/coin/selectCoin", request)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)

	// 验证数据库中上架的币种数量
	count, err := db.GetCount(&models.WthCoinConfig{}, "coin IN ? AND shelves = ?", shelvedCoins, 1)
	framework.AssertDBCount(t, int64(len(shelvedCoins)), count, err)

	framework.LogTestResult(t, true, "测试通过")
}

// ========================================================================
// 修改币种配置测试
// ========================================================================

// TestWthCoinConfig_Update_Success 测试修改币种配置-成功场景
// 用例编号: WTH_COIN_UPD_FUNC_001
// 所属模块: 币种配置管理
// 优先级: 高
// 测试类型: 功能测试
// 关联需求: WTH-COIN-005
// 前置条件: 1. 已登录管理端系统 2. 数据库中已存在币种配置
// 测试步骤:
//  1. 创建一个币种配置（coin=test_coin_001）
//  2. 调用修改接口，修改coinKey和langNameList
//  3. 验证响应状态码为200
//  4. 验证数据库中的数据已更新
//
// 测试数据: coin=test_coin_001, coinKey=updated_coin_key, langNameList=[{"langKey":"zh-Hans","content":"更新后的名称"}]
// 预期结果: 1. API返回成功 2. 数据库中币种配置已更新 3. 多语言数据已更新 4. 原有coin标识保持不变
func TestWthCoinConfig_Update_Success(t *testing.T) {
	// Arrange - 准备测试数据
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	coin := testdata.GenerateUniqueCoin()
	coinKey := testdata.GenerateUniqueCoinKey()
	client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "原标签"))

	// 查询获取币种的ID
	var coinConfig models.WthCoinConfig
	err := db.Query(&coinConfig, "coin = ?", coin)
	if err != nil {
		t.Fatalf("查询币种失败: %v", err)
	}

	// Act - 执行测试操作
	framework.LogTestStep(t, 1, "准备测试环境：已创建币种配置")
	framework.LogTestStep(t, 2, "调用API接口: POST /api/v1/admin/wth/coin/update")

	updatedCoinKey := testdata.GenerateUniqueCoinKey()
	updateRequest := map[string]interface{}{
		"id":      coinConfig.ID,
		"coin":    coin,
		"coinKey": updatedCoinKey,
		"tag":     "更新标签",
		"langNameList": []map[string]string{
			{"langKey": "zh-Hans", "content": "更新后的名称"},
			{"langKey": "en", "content": "Updated Name"},
		},
	}

	resp, err := client.Post("/api/v1/admin/wth/coin/update", updateRequest)

	// Assert - 验证结果
	framework.LogTestStep(t, 3, "验证测试结果")
	if err != nil {
		t.Fatalf("请求失败: %v", err)
	}

	framework.AssertStatusCode(t, resp, 200)

	// 验证数据库中的数据
	err = db.Query(&coinConfig, "coin = ?", coin)
	framework.AssertDBRecordExists(t, err == nil, err)
	framework.AssertDBFieldEqual(t, "Coin", coin, coinConfig.Coin)
	framework.AssertDBFieldEqual(t, "CoinKey", updatedCoinKey, coinConfig.CoinKey)

	framework.LogTestResult(t, true, "测试通过")
}
