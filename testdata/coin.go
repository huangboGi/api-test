package testdata

import (
	cr "crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// NewCoinConfig 创建币种配置测试数据
// 参数符合 dto.WthCoinConfigAdd 结构
// 自动生成唯一的 coinKey，避免重复
func NewCoinConfig(coin string, coinKey string, tag string) map[string]interface{} {
	// 如果 coinKey 为空，自动生成唯一的
	if coinKey == "" {
		coinKey = GenerateUniqueCoinKey()
	}
	return map[string]interface{}{
		"coin":    coin,
		"coinKey": coinKey,
		"tag":     tag,
		"langNameList": []map[string]string{
			{"langKey": "zh-Hans", "content": "中文币种名称"},
			{"langKey": "en", "content": "English Coin Name"},
		},
	}
}

// NewCoinConfigMinimal 创建最小币种配置（仅必填字段）
func NewCoinConfigMinimal(coin string, coinKey string) map[string]interface{} {
	// 如果 coinKey 为空，自动生成唯一的
	if coinKey == "" {
		coinKey = GenerateUniqueCoinKey()
	}
	return map[string]interface{}{
		"coin":    coin,
		"coinKey": coinKey,
		"langNameList": []map[string]string{
			{"langKey": "zh-Hans", "content": "中文币种名称"},
		},
	}
}

// NewCoinConfigUpdate 创建更新币种配置测试数据
func NewCoinConfigUpdate(id int64, coin string, coinKey string) map[string]interface{} {
	// 如果 coinKey 为空，自动生成唯一的
	if coinKey == "" {
		coinKey = GenerateUniqueCoinKey()
	}
	return map[string]interface{}{
		"id":      id,
		"coin":    coin,
		"coinKey": coinKey,
		"tag":     "更新标签",
		"langNameList": []map[string]string{
			{"langKey": "zh-Hans", "content": "更新的中文币种名称"},
			{"langKey": "en", "content": "Updated English Coin Name"},
		},
	}
}

// GenerateUniqueCoinKey 生成唯一的币种标识
// 格式: TEST_8位随机数_纳秒级时间戳
// 使用 crypto/rand 确保随机性，避免并发冲突
func GenerateUniqueCoinKey() string {
	// 生成4字节随机数（8位十六进制）
	randomBytes := make([]byte, 4)
	cr.Read(randomBytes)
	randomHex := hex.EncodeToString(randomBytes)

	// 纳秒级时间戳
	nanos := time.Now().UnixNano()

	return fmt.Sprintf("TEST_%s_%d", randomHex, nanos)
}

// GenerateUniqueCoin 生成唯一的币种代码
// 格式: COIN_8位随机数_纳秒级时间戳
// 使用 crypto/rand 确保随机性，避免并发冲突
func GenerateUniqueCoin() string {
	// 生成4字节随机数（8位十六进制）
	randomBytes := make([]byte, 4)
	cr.Read(randomBytes)
	randomHex := hex.EncodeToString(randomBytes)

	// 纳秒级时间戳
	nanos := time.Now().UnixNano()

	return fmt.Sprintf("COIN_%s_%d", randomHex, nanos)
}
