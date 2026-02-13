package testdata

import (
	"github.com/shopspring/decimal"
)

// NewCoinConfig 创建币种配置（完整参数）
// coinKey 为空时不传入，让后端自动使用 coin 值
func NewCoinConfig(coin, coinKey, tag string, langNameList []map[string]string) map[string]interface{} {
	if len(langNameList) == 0 {
		langNameList = []map[string]string{
			{"langKey": "zh-Hans", "content": "中文币种名称"},
			{"langKey": "en", "content": "English Coin Name"},
		}
	}
	req := map[string]interface{}{
		"coin":         coin,
		"tag":          tag,
		"langNameList": langNameList,
	}
	// 只有明确指定 coinKey 时才传入
	if coinKey != "" {
		req["coinKey"] = coinKey
	}
	return req
}

// NewCoinConfigMinimal 创建最小币种配置（仅必填字段）
// coinKey 为空时不传入，让后端自动使用 coin 值
func NewCoinConfigMinimal(coin, coinKey string) map[string]interface{} {
	req := map[string]interface{}{
		"coin": coin,
		"langNameList": []map[string]string{
			{"langKey": "zh-Hans", "content": "中文币种名称"},
		},
	}
	if coinKey != "" {
		req["coinKey"] = coinKey
	}
	return req
}

// NewCoinShelvesRequest 创建币种上下架请求
func NewCoinShelvesRequest(id int64, shelves int) map[string]interface{} {
	return map[string]interface{}{
		"id":      id,
		"shelves": shelves,
	}
}

// NewSpec 创建规格配置
func NewSpec(specValue int, specKey string, deadlineType int, desc string) map[string]interface{} {
	return map[string]interface{}{
		"specValue":    specValue,
		"specKey":      specKey,
		"deadlineType": deadlineType,
		"desc":         desc,
		"langNameList": []map[string]string{
			{"langKey": "zh-Hans", "content": "活期理财"},
			{"langKey": "en", "content": "Flexible"},
		},
	}
}

// NewSpecWithLang 创建规格配置（自定义语言）
func NewSpecWithLang(specValue int, specKey string, deadlineType int, desc string, langNameList []map[string]string) map[string]interface{} {
	if len(langNameList) == 0 {
		langNameList = []map[string]string{
			{"langKey": "zh-Hans", "content": "活期理财"},
			{"langKey": "en", "content": "Flexible"},
		}
	}
	return map[string]interface{}{
		"specValue":    specValue,
		"specKey":      specKey,
		"deadlineType": deadlineType,
		"desc":         desc,
		"langNameList": langNameList,
	}
}

// NewProduct 创建产品测试数据
func NewProduct(coin, coinKey string, specValue int, deadlineType int) map[string]interface{} {
	if coinKey == "" {
		coinKey = GenerateUniqueCoinKey()
	}
	return map[string]interface{}{
		"coin":           coin,
		"coinKey":        coinKey,
		"classifyKey":    "wealth",
		"specValue":      specValue,
		"deadlineType":   deadlineType,
		"annualAte":      decimal.NewFromInt(5),
		"tag":            "测试标签",
		"minVol":         decimal.NewFromInt(100),
		"useQuotaTotal":  decimal.NewFromInt(100000),
		"personQuota":    decimal.NewFromInt(10000),
		"openSub":        1,
		"shelvesStatus":  0,
		"extraAnnualAte": decimal.NewFromInt(2),
		"dailyMaximum":   decimal.NewFromInt(1000),
		"sort":           1,
	}
}

// NewProductMinimal 创建最小产品配置
func NewProductMinimal(coin, coinKey string, specValue int) map[string]interface{} {
	if coinKey == "" {
		coinKey = GenerateUniqueCoinKey()
	}
	return map[string]interface{}{
		"coin":         coin,
		"coinKey":      coinKey,
		"specValue":    specValue,
		"deadlineType": 0,
	}
}

// NewProductUpdate 创建更新产品测试数据
func NewProductUpdate(id int64, coin string, coinKey string, specValue int) map[string]interface{} {
	if coinKey == "" {
		coinKey = GenerateUniqueCoinKey()
	}
	return map[string]interface{}{
		"id":             id,
		"coin":           coin,
		"coinKey":        coinKey,
		"classifyKey":    "wealth",
		"specValue":      specValue,
		"deadlineType":   0,
		"annualAte":      decimal.NewFromInt(6),
		"tag":            "更新标签",
		"minVol":         decimal.NewFromInt(200),
		"useQuotaTotal":  decimal.NewFromInt(200000),
		"personQuota":    decimal.NewFromInt(20000),
		"openSub":        0,
		"shelvesStatus":  1,
		"extraAnnualAte": decimal.NewFromInt(3),
		"dailyMaximum":   decimal.NewFromInt(2000),
		"sort":           2,
	}
}
