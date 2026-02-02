package cases

import (
	"testing"

	"my_stonks_api_tests/framework"
	"my_stonks_api_tests/models"
	"my_stonks_api_tests/testdata"

	_ "my_stonks_api_tests/config"
)

// TestWthUserOrder_Create 测试创建用户订单接口
func TestWthUserOrder_Create(t *testing.T) {
	client := framework.NewTestClient()
	db := framework.NewDBClient(t)

	// 准备测试数据：币种和产品
	coinKey := "ORDER_TEST_COIN"
	coin := "ORDER_TEST_COIN"
	productKey := "ORDER_TEST_PRODUCT"
	client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "测试"))
	client.Post("/api/v1/admin/wth/product/add", testdata.NewProduct(productKey, "测试产品", coinKey))

	tests := []struct {
		name           string
		request        interface{}
		setup          func()
		wantSuccess    bool
		wantMsgContain string
		verify         func(orderNo string)
	}{
		{
			name: "正常创建订单",
			request: map[string]interface{}{
				"userId":       1001,
				"productId":    1,
				"coinKey":      coinKey,
				"investAmount": "1000.00",
			},
			wantSuccess: true,
			verify: func(orderNo string) {
				var order models.WthUserOrder
				err := db.GetByCondition(&order, "order_no = ?", orderNo)
				framework.AssertDBRecordExists(t, err == nil, err)
				framework.AssertDBFieldEqual(t, "UserID", uint(1001), order.UserID)
				framework.AssertDBFieldEqual(t, "OrderStatus", int8(0), order.OrderStatus)
			},
		},
		{
			name: "用户ID为空应失败",
			request: map[string]interface{}{
				"productId":    1,
				"coinKey":      coinKey,
				"investAmount": "1000.00",
			},
			wantSuccess:    false,
			wantMsgContain: "用户ID不能为空",
		},
		{
			name: "投资金额为空应失败",
			request: map[string]interface{}{
				"userId":    1001,
				"productId": 1,
				"coinKey":   coinKey,
			},
			wantSuccess:    false,
			wantMsgContain: "投资金额不能为空",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			framework.LogTestStep(t, 1, "准备测试环境")
			if tt.setup != nil {
				tt.setup()
			}

			framework.LogTestStep(t, 2, "调用API接口")
			resp, err := client.Post("/api/v1/wth/user/order/create", tt.request)

			framework.LogTestStep(t, 3, "验证测试结果")
			if err != nil {
				t.Errorf("请求失败: %v", err)
				return
			}

			if tt.wantSuccess {
				framework.AssertStatusCode(t, resp, 200)
				if tt.verify != nil {
					// 从响应中获取订单号
					if data, ok := resp.Data.(map[string]interface{}); ok {
						if orderNo, ok := data["orderNo"].(string); ok {
							tt.verify(orderNo)
						}
					}
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

// TestWthUserOrder_Page 测试用户订单分页查询接口
func TestWthUserOrder_Page(t *testing.T) {
	client := framework.NewTestClient()

	// 准备测试数据
	coinKey := "ORDER_PAGE_COIN"
	coin := "ORDER_PAGE_COIN"
	productKey := "ORDER_PAGE_PRODUCT"
	client.Post("/api/v1/admin/wth/coin/add", testdata.NewCoinConfig(coin, coinKey, "测试"))
	client.Post("/api/v1/admin/wth/product/add", testdata.NewProduct(productKey, "测试产品", coinKey))

	tests := []struct {
		name    string
		request interface{}
		verify  func()
	}{
		{
			name: "查询所有订单",
			request: map[string]interface{}{
				"pageIndex": 1,
				"pageSize":  10,
			},
		},
		{
			name: "按用户ID查询",
			request: map[string]interface{}{
				"pageIndex": 1,
				"pageSize":  10,
				"userId":    1001,
			},
		},
		{
			name: "按币种查询",
			request: map[string]interface{}{
				"pageIndex": 1,
				"pageSize":  10,
				"coinKey":   coinKey,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.Post("/api/v1/wth/user/order/page", tt.request)
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
