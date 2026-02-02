package testdata

// NewProduct 创建产品测试数据
func NewProduct(productKey, productName, coinKey string) map[string]interface{} {
	return map[string]interface{}{
		"productKey":  productKey,
		"productName": productName,
		"coinKey":     coinKey,
		"status":      1,
		"sort":        1,
	}
}

// NewProductWithSort 创建自定义排序的产品
func NewProductWithSort(productKey, productName, coinKey string, status, sort int) map[string]interface{} {
	return map[string]interface{}{
		"productKey":  productKey,
		"productName": productName,
		"coinKey":     coinKey,
		"status":      status,
		"sort":        sort,
	}
}
