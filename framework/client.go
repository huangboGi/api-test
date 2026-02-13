package framework

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"my_stonks_api_tests/config"
)

// TestClient HTTP测试客户端
type TestClient struct {
	client    *http.Client
	authToken string
	override  bool
}

// NewTestClient 创建测试客户端
func NewTestClient() *TestClient {
	return &TestClient{
		client: &http.Client{
			Timeout: time.Duration(config.Cfg.TestTimeout) * time.Second,
		},
	}
}

// TestResponse 测试响应结构体
type TestResponse struct {
	StatusCode int         `json:"statusCode"`
	Code       int         `json:"code"`
	Message    string      `json:"msg"`
	Data       interface{} `json:"data"`
	RawBody    string      `json:"-"`
}

// doRequest 发送HTTP请求的内部方法
func (c *TestClient) doRequest(method, endpoint string, body interface{}) (*TestResponse, error) {
	url := config.Cfg.APIBaseURL + endpoint

	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body failed: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// 设置Headers
	req.Header.Set("Content-Type", "application/json")

	// 根据端点自动选择Token
	tokenType := "user"
	if strings.Contains(endpoint, "/admin/") {
		tokenType = "admin"
	}

	var token string
	if c.override {
		token = c.authToken
		fmt.Printf("【请求】%s %s (使用自定义Token)\n", method, url)
	} else if tokenType == "admin" {
		token = config.Cfg.AdminToken
		fmt.Printf("【请求】%s %s (使用管理端Token)\n", method, url)
	} else {
		token = config.Cfg.UserToken
		fmt.Printf("【请求】%s %s (使用用户端Token)\n", method, url)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	// 打印请求Body（便于调试）
	if body != nil {
		jsonData, _ := json.Marshal(body)
		fmt.Printf("【请求Body】%s\n", string(jsonData))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	// 打印响应信息
	fmt.Printf("【响应】Status: %d, Body: %s\n", resp.StatusCode, string(bodyBytes))

	var testResp TestResponse
	if err := json.Unmarshal(bodyBytes, &testResp); err != nil {
		return &TestResponse{
			StatusCode: resp.StatusCode,
			RawBody:    string(bodyBytes),
		}, nil
	}

	testResp.StatusCode = resp.StatusCode
	testResp.RawBody = string(bodyBytes)

	// 检查Token是否过期（code为401）
	if testResp.Code == 401 {
		tokenType := "用户端"
		if strings.Contains(endpoint, "/admin/") {
			tokenType = "管理端"
		}
		panic(fmt.Sprintf(
			"\n========================================\n"+
				"❌ Token已过期！\n"+
				"========================================\n"+
				"接口路径: %s\n"+
				"Token类型: %s\n"+
				"响应Code: %d\n"+
				"错误信息: %s\n"+
				"\n解决方法：\n"+
				"1. 重新登录获取新的Token\n"+
				"2. 更新 .env 文件中的 %s\n"+
				"3. 重新运行测试\n"+
				"========================================",
			endpoint, tokenType, testResp.Code, testResp.Message,
			strings.ToUpper(tokenType)+"_TOKEN",
		))
	}

	return &testResp, nil
}

// Get 发送GET请求
func (c *TestClient) Get(endpoint string) (*TestResponse, error) {
	return c.doRequest("GET", endpoint, nil)
}

// Post 发送POST请求
func (c *TestClient) Post(endpoint string, body interface{}) (*TestResponse, error) {
	return c.doRequest("POST", endpoint, body)
}

// Put 发送PUT请求
func (c *TestClient) Put(endpoint string, body interface{}) (*TestResponse, error) {
	return c.doRequest("PUT", endpoint, body)
}

// Delete 发送DELETE请求
func (c *TestClient) Delete(endpoint string) (*TestResponse, error) {
	return c.doRequest("DELETE", endpoint, nil)
}

// SetAuthToken 设置认证Token
func (c *TestClient) SetAuthToken(token string) {
	c.authToken = token
	c.override = token != ""
}
