# 双Token配置说明

## 概述

本测试框架支持两种类型的Token，用于区分管理端和用户端接口的权限验证。

## Token类型

### 1. ADMIN_TOKEN（管理端Token）
- **用途**：所有管理端接口的认证
- **接口特征**：路径包含 `/admin/`
- **示例接口**：
  - `/api/v1/admin/wth/coin/add` - 添加币种配置
  - `/api/v1/admin/wth/product/add` - 添加产品
  - `/api/v1/admin/wth/coin/page` - 查询币种列表
  - `/api/v1/admin/wth/coin/updateShelves` - 币种上下架

### 2. USER_TOKEN（用户端Token）
- **用途**：所有用户端接口的认证
- **接口特征**：路径不包含 `/admin/`
- **示例接口**：
  - `/api/v1/wth/coin/selectCoin` - 查询可用币种列表
  - `/api/v1/wth/user/order/create` - 创建用户订单
  - `/api/v1/wth/user/order/page` - 查询用户订单

## 配置方式

在 `.env` 文件中配置两种Token：

```env
# API基础配置
API_BASE_URL=http://localhost:8080

# 管理端 Token
ADMIN_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

# 用户端 Token
USER_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=your_password
DB_NAME=my_stonks
```

## 自动Token选择机制

测试框架会根据接口路径自动选择正确的Token：

```go
// 框架内部实现（无需手动设置）
func (c *TestClient) doRequest(method, endpoint string, body interface{}) {
    // 根据端点自动选择Token
    tokenType := "user"
    if strings.Contains(endpoint, "/admin/") {
        tokenType = "admin"
    }

    var token string
    if tokenType == "admin" {
        token = config.Cfg.AdminToken
    } else {
        token = config.Cfg.UserToken
    }

    req.Header.Set("Authorization", "Bearer "+token)
    // ...
}
```

## 测试用例示例

### 管理端接口测试

```go
func TestWthCoinConfig_Add(t *testing.T) {
    client := framework.NewTestClient()

    // 自动使用 ADMIN_TOKEN（因为路径包含 /admin/）
    resp, err := client.Post("/api/v1/admin/wth/coin/add", request)
}
```

### 用户端接口测试

```go
func TestWthCoin_SelectCoin(t *testing.T) {
    client := framework.NewTestClient()

    // 自动使用 USER_TOKEN（因为路径不包含 /admin/）
    resp, err := client.Post("/api/v1/wth/coin/selectCoin", request)
}
```

## 日志输出

测试运行时会显示使用的Token类型：

```
【请求】POST http://localhost:8080/api/v1/admin/wth/coin/add (使用管理端Token)
【请求Body】{"coinKey":"BTC","coinIcon":"..."}
【响应】Status: 200, Body: {"code":0,"message":"success"}
```

或

```
【请求】POST http://localhost:8080/api/v1/wth/coin/selectCoin (使用用户端Token)
【请求Body】{}
【响应】Status: 200, Body: {"code":0,"message":"success"}
```

## 注意事项

1. **必须配置两个Token**：`ADMIN_TOKEN` 和 `USER_TOKEN` 都必须配置，否则会报错
2. **自动选择**：无需手动指定Token类型，框架会自动选择
3. **Token权限**：确保配置的Token具有对应接口的访问权限
4. **Token过期**：如果Token过期，需要更新`.env`文件中的Token

## 获取Token的方式

### 方式1：通过登录接口获取

使用Postman或curl调用登录接口：

```bash
# 管理员登录
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'

# 用户登录
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"user","password":"user123"}'
```

响应中的 `data.token` 就是Token，复制到 `.env` 文件中。

### 方式2：从已有项目中获取

如果已经有运行的系统，可以从浏览器开发者工具或抓包工具中获取Token：

1. 打开浏览器开发者工具（F12）
2. 切换到 Network 标签
3. 发送一个请求
4. 查看请求头中的 `Authorization` 字段
5. 复制 `Bearer ` 后面的Token

## 常见问题

### Q: 可以只用一个Token吗？
A: 不可以。管理端和用户端使用不同的权限体系，必须配置两个Token。

### Q: Token有效期多久？
A: 取决于后端的Token配置，通常是几小时到几天。

### Q: Token过期了会怎么样？
A: 如果Token过期，测试会**立即停止**，并显示错误信息：

```
========================================
❌ Token已过期！
========================================
接口路径: /api/v1/admin/wth/coin/add
Token类型: 管理端
响应Code: 401
错误信息: token expired

解决方法：
1. 重新登录获取新的Token
2. 更新 .env 文件中的 ADMIN_TOKEN
3. 重新运行测试
========================================
```

**解决步骤：**
1. 重新登录获取新Token
2. 编辑`.env`文件，更新对应的`ADMIN_TOKEN`或`USER_TOKEN`
3. 重新运行`make test`

### Q: 如何查看使用的Token类型？
A: 查看测试日志，会显示"使用管理端Token"或"使用用户端Token"。

### Q: 不同环境需要不同的Token吗？
A: 是的。测试环境、预发布环境、生产环境需要分别获取对应的Token。
