package framework

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// AssertStatusCode 断言HTTP状态码
func AssertStatusCode(t *testing.T, resp *TestResponse, expectedCode int) {
	assert.Equal(t, expectedCode, resp.StatusCode,
		"期望状态码: %d, 实际: %d, 响应: %s", expectedCode, resp.StatusCode, resp.RawBody)
}

// AssertSuccess 断言API返回成功
func AssertSuccess(t *testing.T, resp *TestResponse) {
	AssertStatusCode(t, resp, 200)
	assert.Equal(t, 0, resp.Code, "期望Code: 0, 实际: %d, 消息: %s", resp.Code, resp.Message)
}

// AssertAPIError 断言API返回错误
func AssertAPIError(t *testing.T, resp *TestResponse, expectedCode int) {
	AssertStatusCode(t, resp, 200)
	assert.NotEqual(t, 0, resp.Code, "期望Code不为0，但实际为0")
	if expectedCode > 0 {
		assert.Equal(t, expectedCode, resp.Code, "期望Code: %d, 实际: %d", expectedCode, resp.Code)
	}
}

// AssertErrorMessageContains 断言错误消息包含指定内容
func AssertErrorMessageContains(t *testing.T, resp *TestResponse, msg string) {
	assert.Contains(t, resp.Message, msg,
		"错误消息应包含: %s, 实际: %s", msg, resp.Message)
}

// AssertDBRecordExists 断言数据库记录存在
func AssertDBRecordExists(t *testing.T, exists bool, err error) {
	assert.NoError(t, err, "查询数据库不应出错")
	assert.True(t, exists, "数据库中应该存在该记录")
}

// AssertDBRecordNotExists 断言数据库记录不存在
func AssertDBRecordNotExists(t *testing.T, exists bool, err error) {
	assert.NoError(t, err, "查询数据库不应出错")
	assert.False(t, exists, "数据库中不应该存在该记录")
}

// AssertDBCount 断言记录数量
func AssertDBCount(t *testing.T, expected int64, actual int64, err error) {
	assert.NoError(t, err, "查询数据库不应出错")
	assert.Equal(t, expected, actual,
		"期望记录数: %d, 实际: %d", expected, actual)
}

// AssertDBFieldEqual 断言字段值相等
func AssertDBFieldEqual[T comparable](t *testing.T, field string, expected, actual T) {
	assert.Equal(t, expected, actual,
		"字段 %s 的值不匹配，期望: %v, 实际: %v", field, expected, actual)
}

// AssertDecimalEqual 断言Decimal值相等（处理精度问题）
func AssertDecimalEqual(t *testing.T, field string, expected, actual decimal.Decimal) {
	assert.True(t, expected.Equal(actual),
		"字段 %s 的值不匹配，期望: %s, 实际: %s", field, expected.String(), actual.String())
}

// AssertDBFieldNotEqual 断言字段值不相等
func AssertDBFieldNotEqual[T comparable](t *testing.T, field string, notExpected, actual T) {
	assert.NotEqual(t, notExpected, actual,
		"字段 %s 的值不应为: %v, 实际: %v", field, notExpected, actual)
}

// LogTestStep 打印测试步骤
func LogTestStep(t *testing.T, step int, description string) {
	t.Logf("【步骤%d】%s", step, description)
}

// LogTestResult 打印测试结果
func LogTestResult(t *testing.T, success bool, message string) {
	status := "✅ 通过"
	if !success {
		status = "❌ 失败"
	}
	t.Logf("【测试结果】%s - %s", status, message)
}

// LogDBQuery 打印数据库查询结果
func LogDBQuery(t *testing.T, table, condition string, result interface{}) {
	t.Logf("【数据库查询】表: %s, 条件: %s, 结果: %+v", table, condition, result)
}

// AssertStringNotEmpty 断言字符串不为空
func AssertStringNotEmpty(t *testing.T, field string, value string) {
	if value == "" {
		t.Errorf("字段 %s 不应为空", field)
	}
}

// Logf 打印日志
func Logf(t *testing.T, format string, args ...interface{}) {
	t.Logf(format, args...)
}

// LogResponse 打印API响应
func LogResponse(t *testing.T, resp *TestResponse) {
	t.Logf("【API响应】状态码: %d, Code: %d, 消息: %s, 响应体: %s",
		resp.StatusCode, resp.Code, resp.Message, resp.RawBody)
}
