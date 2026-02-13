package framework

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"my_stonks_api_tests/config"
)

// ParseResponseBody 解析响应体
func ParseResponseBody(resp *TestResponse, data interface{}) error {
	if resp.Data == nil {
		respBytes := []byte(resp.RawBody)
		return json.Unmarshal(respBytes, data)
	}

	jsonData, err := json.Marshal(resp.Data)
	if err != nil {
		return fmt.Errorf("marshal response data failed: %w", err)
	}

	return json.Unmarshal(jsonData, data)
}

// GetTestUserId 从配置中获取测试用户ID
func GetTestUserId() int {
	token := config.Cfg.UserToken
	userId, err := extractUserIdFromToken(token)
	if err != nil {
		panic(fmt.Sprintf("从Token中提取用户ID失败: %v", err))
	}
	return userId
}

// GetTestUserIdWithIndex 根据索引生成测试用户ID
func GetTestUserIdWithIndex(index int) int {
	// 基于基础用户ID生成其他用户ID
	baseUserId := GetTestUserId()
	return baseUserId + index
}

// extractUserIdFromToken 从JWT Token中提取用户ID
func extractUserIdFromToken(token string) (int, error) {
	if token == "" {
		return 0, fmt.Errorf("token为空")
	}

	// JWT token格式: header.payload.signature
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("token格式错误")
	}

	// 解析payload
	payload := parts[1]
	if len(payload) < 4 {
		return 0, fmt.Errorf("payload太短")
	}

	// Base64 URL 解码
	decoded, err := decodeBase64(payload)
	if err != nil {
		return 0, fmt.Errorf("base64解码失败: %w", err)
	}

	// JSON解析
	jsonData := make(map[string]interface{})
	if err := json.Unmarshal(decoded, &jsonData); err != nil {
		return 0, fmt.Errorf("JSON解析失败: %w", err)
	}

	// 提取userId
	userId, ok := jsonData["userId"]
	if !ok {
		return 0, fmt.Errorf("token中未找到userId字段")
	}

	// 转换为int
	switch v := userId.(type) {
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("userId类型不支持: %T", userId)
	}
}

// decodeBase64 Base64 URL 解码
func decodeBase64(payload string) ([]byte, error) {
	// 添加填充
	switch len(payload) % 4 {
	case 2:
		payload += "=="
	case 3:
		payload += "="
	}

	// 替换 URL 安全字符
	payload = strings.ReplaceAll(payload, "-", "+")
	payload = strings.ReplaceAll(payload, "_", "/")

	// 解码
	result := make([]byte, len(payload)*6/8)
	n, err := decodeBase64Impl(payload, result)
	if err != nil {
		return nil, err
	}
	return result[:n], nil
}

// decodeBase64Impl 简化的 Base64 解码实现
func decodeBase64Impl(s string, dst []byte) (int, error) {
	const std = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	decodeMap := make(map[byte]int)
	for i := 0; i < 64; i++ {
		decodeMap[std[i]] = i
	}

	var n int
	for i := 0; i < len(s); i += 4 {
		v0, ok0 := decodeMap[s[i]]
		v1, ok1 := decodeMap[s[i+1]]
		if !ok0 || !ok1 {
			return 0, fmt.Errorf("invalid base64 character")
		}

		dst[n] = byte(v0<<2 | v1>>4)
		n++

		if s[i+2] == '=' {
			break
		}
		v2, ok2 := decodeMap[s[i+2]]
		if !ok2 {
			return 0, fmt.Errorf("invalid base64 character")
		}

		dst[n] = byte(v1<<4 | v2>>2)
		n++

		if s[i+3] == '=' {
			break
		}
		v3, ok3 := decodeMap[s[i+3]]
		if !ok3 {
			return 0, fmt.Errorf("invalid base64 character")
		}

		dst[n] = byte(v2<<6 | v3)
		n++
	}

	return n, nil
}

// GetUserBalance 获取用户余额（从account表）
func GetUserBalance(db *DBClient, userId int, symbol string) (string, error) {
	type Account struct {
		Balance string `gorm:"column:balance"`
	}

	var account Account
	err := db.Table("account").Where("user_id = ? AND symbol = ?", userId, symbol).First(&account).Error
	if err != nil {
		return "", fmt.Errorf("查询用户余额失败: %w", err)
	}

	return account.Balance, nil
}

// CheckUserBalance 检查用户余额是否足够
// 如果余额不足，打印提示信息并 panic
// 返回当前余额（如果余额足够）
func CheckUserBalance(db *DBClient, userId int, symbol string, requiredAmount string) string {
	currentBalance, err := GetUserBalance(db, userId, symbol)
	if err != nil {
		panic(fmt.Sprintf("查询用户余额失败: %v", err))
	}

	fmt.Printf("【余额检查】用户ID: %d, 币种: %s, 当前余额: %s, 需要余额: %s\n",
		userId, symbol, currentBalance, requiredAmount)

	// 比较余额（字符串比较）
	if currentBalance < requiredAmount {
		panic(fmt.Sprintf("⚠️  余额不足！请手动充值后再运行测试\n"+
			"当前余额: %s\n"+
			"需要余额: %s\n"+
			"用户ID: %d\n"+
			"币种: %s\n\n"+
			"请执行以下SQL充值余额:\n"+
			"UPDATE account SET balance = '%s' WHERE user_id = %d AND symbol = '%s';",
			currentBalance, requiredAmount, userId, symbol, requiredAmount, userId, symbol))
	}

	fmt.Printf("✅ 余额充足，继续执行测试\n")
	return currentBalance
}
