package testdata

import (
	"fmt"
	"sync"
	"time"

	"github.com/shopspring/decimal"

	"my_stonks_api_tests/api"
	"my_stonks_api_tests/config"
	"my_stonks_api_tests/framework"
	"my_stonks_api_tests/models"
)

var (
	configOnce sync.Once
)

// TestEnv 测试环境构建器
type TestEnv struct {
	Client *framework.TestClient
	DB     *framework.DBClient
	UserID int

	// 环境资源
	Coin         string
	CoinKey      string
	CoinID       int64
	SpecValue    int
	DeadlineType int
	SpecID       int64
	ProductID    int64
	OrderNo      string
}

// NewTestEnv 创建测试环境
func NewTestEnv(caseID string) *TestEnv {
	// 确保配置只加载一次
	configOnce.Do(func() {
		config.Load()
	})

	client := framework.NewTestClient()

	env := &TestEnv{
		Client: client,
		DB:     framework.NewDBClient(),
		UserID: framework.GetTestUserId(),
	}

	return env
}

// WithCoin 创建并上架币种
func (e *TestEnv) WithCoin() *TestEnv {
	e.Coin = GenerateUniqueCoin()
	e.CoinKey = GenerateUniqueCoinKey()

	resp, err := e.Client.Post(api.AdminCoinAdd, NewCoinConfig(e.Coin, e.CoinKey, "测试", nil))
	if err != nil {
		panic(fmt.Sprintf("创建币种失败: %v", err))
	}
	if resp.StatusCode != 200 || resp.Code != 0 {
		panic(fmt.Sprintf("创建币种失败，状态码: %d, 业务码: %d, 响应: %s", resp.StatusCode, resp.Code, resp.RawBody))
	}

	var coin models.WthCoinConfig
	if err := e.DB.Query(&coin, "coin = ?", e.Coin); err != nil {
		panic(fmt.Sprintf("查询币种失败: %v", err))
	}
	e.CoinID = coin.ID

	e.Client.Post(api.AdminCoinShelves, map[string]interface{}{
		"id":      e.CoinID,
		"shelves": 1,
	})
	return e
}

// WithSpec 创建并上架规格
func (e *TestEnv) WithSpec(specValue, deadlineType int) *TestEnv {
	e.SpecValue = specValue
	e.DeadlineType = deadlineType

	// 活期规格（specValue = -1）是系统预置的，直接查询获取
	if specValue == -1 && deadlineType == 0 {
		var spec models.WthSpec
		if err := e.DB.Query(&spec, "spec_value = ? AND deadline_type = ?", -1, 0); err != nil {
			panic(fmt.Sprintf("查询活期规格失败: %v，请确保数据库中存在活期规格记录", err))
		}
		e.SpecID = spec.ID
		fmt.Printf("【规格】使用预置活期规格，ID: %d\n", e.SpecID)
		return e
	}

	// 定期规格需要创建
	specDesc := "定期规格"
	actualSpecValue := GenerateUniqueSpecValue()
	e.SpecValue = actualSpecValue

	resp, err := e.Client.Post(api.AdminSpecAdd, NewSpec(actualSpecValue, "", deadlineType, specDesc))
	if err != nil {
		panic(fmt.Sprintf("创建规格失败: %v", err))
	}
	if resp.StatusCode != 200 || resp.Code != 0 {
		panic(fmt.Sprintf("创建规格失败，状态码: %d, 业务码: %d, 响应: %s", resp.StatusCode, resp.Code, resp.RawBody))
	}

	var spec models.WthSpec
	if err := e.DB.Query(&spec, "spec_value = ? AND deadline_type = ?", actualSpecValue, deadlineType); err != nil {
		panic(fmt.Sprintf("查询规格失败: %v", err))
	}
	e.SpecID = spec.ID

	e.Client.Post(api.AdminSpecShelves, map[string]interface{}{
		"id":            e.SpecID,
		"shelvesStatus": 1,
	})
	return e
}

// WithProduct 创建并上架产品
func (e *TestEnv) WithProduct(modifiers ...func(map[string]interface{})) *TestEnv {
	req := NewProduct(e.Coin, e.CoinKey, e.SpecValue, e.DeadlineType)
	for _, m := range modifiers {
		if m != nil {
			m(req)
		}
	}

	resp, err := e.Client.Post(api.AdminProductAdd, req)
	if err != nil {
		panic(fmt.Sprintf("创建产品失败: %v", err))
	}
	if resp.StatusCode != 200 || resp.Code != 0 {
		panic(fmt.Sprintf("创建产品失败，状态码: %d, 业务码: %d, 响应: %s", resp.StatusCode, resp.Code, resp.RawBody))
	}

	var product models.WthProduct
	if err := e.DB.Query(&product, "coin = ?", e.Coin); err != nil {
		panic(fmt.Sprintf("查询产品失败: %v", err))
	}
	e.ProductID = product.ID

	e.Client.Post(api.AdminProductShelves, map[string]interface{}{
		"id":            e.ProductID,
		"shelvesStatus": 1,
	})
	return e
}

// WithBalance 检查余额是否足够
func (e *TestEnv) WithBalance(minBalance string) *TestEnv {
	balance, err := framework.GetUserBalance(e.DB, e.UserID, e.Coin)
	if err != nil {
		panic(fmt.Sprintf("查询用户余额失败: %v", err))
	}

	current, _ := decimal.NewFromString(balance)
	required, _ := decimal.NewFromString(minBalance)

	if current.LessThan(required) {
		panic(fmt.Sprintf("⚠️ 余额不足！请手动充值\n"+
			"当前余额: %s\n需要余额: %s\n用户ID: %d\n币种: %s\n\n"+
			"执行SQL: UPDATE account SET balance = '%s' WHERE user_id = %d AND symbol = '%s';",
			balance, minBalance, e.UserID, e.Coin, minBalance, e.UserID, e.Coin))
	}
	return e
}

// WithDefaults 使用默认配置（活期产品）
func (e *TestEnv) WithDefaults() *TestEnv {
	return e.WithCoin().WithSpec(-1, 0).WithProduct()
}

// Build 完成环境构建
func (e *TestEnv) Build() *TestEnv {
	return e
}

// ========== 查询验证方法 ==========

// GetBalance 获取用户余额
func (e *TestEnv) GetBalance(symbol string) decimal.Decimal {
	balance, err := framework.GetUserBalance(e.DB, e.UserID, symbol)
	if err != nil {
		return decimal.Zero
	}
	result, _ := decimal.NewFromString(balance)
	return result
}

// GetOrder 获取订单
func (e *TestEnv) GetOrder() *models.WthUserOrder {
	var order models.WthUserOrder
	err := e.DB.Where("user_id = ? AND product_id = ?", e.UserID, e.ProductID).First(&order).Error
	if err != nil {
		return nil
	}
	return &order
}

// GetOrderByNo 根据订单号获取订单
func (e *TestEnv) GetOrderByNo(orderNo string) *models.WthUserOrder {
	var order models.WthUserOrder
	err := e.DB.Query(&order, "order_no = ?", orderNo)
	if err != nil {
		return nil
	}
	return &order
}

// GetOrders 获取用户所有订单
func (e *TestEnv) GetOrders() []models.WthUserOrder {
	var orders []models.WthUserOrder
	e.DB.Where("user_id = ? AND product_id = ?", e.UserID, e.ProductID).Find(&orders)
	return orders
}

// GetSubscribeHis 获取申购历史
func (e *TestEnv) GetSubscribeHis(orderID int64) *models.WthUserSubscribeHis {
	var his models.WthUserSubscribeHis
	err := e.DB.Query(&his, "order_id = ?", orderID)
	if err != nil {
		return nil
	}
	return &his
}

// BalanceShouldBe 验证余额
func (e *TestEnv) BalanceShouldBe(symbol string, expected string) {
	actual := e.GetBalance(symbol)
	expectedDec, _ := decimal.NewFromString(expected)
	if !actual.Equal(expectedDec) {
		fmt.Printf("[ERROR] 余额不正确！期望: %s, 实际: %s\n", expected, actual.String())
	}
}

// BalanceShouldChange 验证余额变化
func (e *TestEnv) BalanceShouldChange(symbol string, before decimal.Decimal, delta decimal.Decimal, isAdd bool) {
	after := e.GetBalance(symbol)
	var expected decimal.Decimal
	if isAdd {
		expected = before.Add(delta)
	} else {
		expected = before.Sub(delta)
	}
	if !after.Equal(expected) {
		fmt.Printf("[ERROR] 余额变化不正确！期望: %s, 实际: %s\n", expected.String(), after.String())
	}
}

// Subscribe 执行申购
func (e *TestEnv) Subscribe(volume decimal.Decimal) (*framework.TestResponse, error) {
	return e.Client.Post(api.UserSubscribe, map[string]interface{}{
		"coin":         e.Coin,
		"specValue":    e.SpecValue,
		"volume":       volume,
		"deadlineType": e.DeadlineType,
	})
}

// Redeem 执行赎回
func (e *TestEnv) Redeem(orderNo string, volume decimal.Decimal) (*framework.TestResponse, error) {
	return e.Client.Post(api.UserRedeem, map[string]interface{}{
		"orderNo": orderNo,
		"volume":  volume,
	})
}

// LogStep 打印测试步骤
func (e *TestEnv) LogStep(step int, desc string) {
	fmt.Printf("【步骤%d】%s\n", step, desc)
}

// LogResult 打印测试结果
func (e *TestEnv) LogResult(success bool, msg string) {
	if success {
		fmt.Printf("✅ %s\n", msg)
	} else {
		fmt.Printf("❌ %s\n", msg)
	}
}

// PrintEnvInfo 打印环境信息
func (e *TestEnv) PrintEnvInfo() {
	fmt.Printf("========== 测试环境信息 ==========\n")
	fmt.Printf("用户ID: %d\n", e.UserID)
	fmt.Printf("币种: %s (ID: %d)\n", e.Coin, e.CoinID)
	fmt.Printf("规格: %d (类型: %d, ID: %d)\n", e.SpecValue, e.DeadlineType, e.SpecID)
	fmt.Printf("产品ID: %d\n", e.ProductID)
	fmt.Printf("==================================\n")
}

// DecimalFromInt 从整数创建 Decimal
func DecimalFromInt(n int64) decimal.Decimal {
	return decimal.NewFromInt(n)
}

// DecimalFromString 从字符串创建 Decimal
func DecimalFromString(s string) decimal.Decimal {
	d, _ := decimal.NewFromString(s)
	return d
}

// GenerateUniqueCoin 生成唯一币种名称
func GenerateUniqueCoin() string {
	return fmt.Sprintf("COIN_%x_%d", time.Now().UnixNano(), time.Now().Nanosecond())
}

// GenerateUniqueCoinKey 生成唯一币种Key
func GenerateUniqueCoinKey() string {
	return fmt.Sprintf("TEST_%x_%d", time.Now().UnixNano(), time.Now().Nanosecond())
}

// GenerateUniqueSpecValue 生成唯一规格值
func GenerateUniqueSpecValue() int {
	return int(time.Now().UnixNano()%10000) + 100
}
