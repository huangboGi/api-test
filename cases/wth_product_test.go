package cases

import (
	"testing"

	"my_stonks_api_tests/framework"
	"my_stonks_api_tests/models"
	"my_stonks_api_tests/testdata"

	_ "my_stonks_api_tests/config"
)

// TestWthProduct_Add 测试产品添加接口
func TestWthProduct_Add(t *testing.T) {
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	// 准备币种数据
	coinKey := "PRODUCT_TEST_COIN"
	coin := "PRODUCT_TEST_COIN"
	client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "测试"))

	tests := []struct {
		name           string
		request        interface{}
		wantSuccess    bool
		wantMsgContain string
		verify         func(productKey string)
	}{
		{
			name:        "正常添加产品",
			request:     testdata.NewProduct("PROD001", "测试产品1", coinKey),
			wantSuccess: true,
			verify: func(productKey string) {
				var product models.WthProduct
				err := db.GetByCondition(&product, "product_key = ?", productKey)
				framework.AssertDBRecordExists(t, err == nil, err)
				framework.AssertDBFieldEqual(t, "ProductKey", productKey, product.ProductKey)
				framework.AssertDBFieldEqual(t, "Status", int8(1), product.Status)
			},
		},
		{
			name:           "产品标识为空应失败",
			request:        testdata.NewProduct("", "测试产品", coinKey),
			wantSuccess:    false,
			wantMsgContain: "不能为空",
		},
		{
			name:           "币种不存在应失败",
			request:        testdata.NewProduct("PROD002", "测试产品2", "NON_EXIST_COIN"),
			wantSuccess:    false,
			wantMsgContain: "币种不存在",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			framework.LogTestStep(t, 1, "准备测试环境")
			framework.LogTestStep(t, 2, "调用API接口")
			resp, err := client.Post("/api/v1/admin/wth/product/add", tt.request)

			framework.LogTestStep(t, 3, "验证测试结果")
			if err != nil {
				t.Errorf("请求失败: %v", err)
				return
			}

			if tt.wantSuccess {
				framework.AssertStatusCode(t, resp, 200)
				if tt.verify != nil {
					tt.verify(tt.request.(map[string]interface{})["productKey"].(string))
				}
				framework.LogTestResult(t, true, "测试通过")
			} else {
				framework.AssertStatusCode(t, resp, 200)
				if tt.wantMsgContain != "" {
					framework.AssertErrorMessageContains(t, resp, tt.wantMsgContain)
				}
				framework.LogTestResult(t, true, "测试通过")
			}
		})
	}
}

// TestWthProduct_Page 测试产品分页查询接口
func TestWthProduct_Page(t *testing.T) {
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	// 准备测试数据
	coinKey := "PAGE_TEST_COIN"
	coin := "PAGE_TEST_COIN"
	client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "测试"))

	testProductKeys := []string{"PAGE_PROD001", "PAGE_PROD002", "PAGE_PROD003"}
	for _, pk := range testProductKeys {
		client.Post("/api/v1/admin/wth/product/add", testdata.NewProduct(pk, "测试产品"+pk, coinKey))
	}

	tests := []struct {
		name    string
		request interface{}
		verify  func()
	}{
		{
			name: "查询所有产品",
			request: map[string]interface{}{
				"pageIndex": 1,
				"pageSize":  10,
			},
			verify: func() {
				count, err := db.GetCount(&models.WthProduct{}, "product_key IN ?", testProductKeys)
				framework.AssertDBCount(t, int64(len(testProductKeys)), count, err)
			},
		},
		{
			name: "按币种查询",
			request: map[string]interface{}{
				"pageIndex": 1,
				"pageSize":  10,
				"coinKey":   coinKey,
			},
			verify: func() {
				count, err := db.GetCount(&models.WthProduct{}, "coin_key = ?", coinKey)
				framework.AssertDBRecordExists(t, count > 0, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.Post("/api/v1/admin/wth/product/page", tt.request)
			if err != nil {
				t.Errorf("请求失败: %v", err)
				return
			}

			framework.AssertStatusCode(t, resp, 200)
			if tt.verify != nil {
				tt.verify()
			}
		})
	}
}
