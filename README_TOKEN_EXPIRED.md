# Token过期处理指南

## Token过期检测机制

测试框架会自动检测Token是否过期。当API返回`code=401`时，测试会**立即停止**，避免继续执行无效的测试。

## Token过期时的错误信息

当Token过期时，你会看到如下错误信息：

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

## Token过期的原因

1. **Token有效期到**：Token有有效期，通常为几小时到几天
2. **密码修改**：修改密码后，旧Token会失效
3. **强制登出**：管理员强制用户登出
4. **服务器重启**：服务器重启可能导致Token失效（取决于配置）

## 如何重新获取Token

### 方式1：通过登录接口（推荐）

#### 管理端Token
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

响应示例：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTk5MjAwMDB9.XXXX",
    "user": {...}
  }
}
```

复制`data.token`的值到`.env`文件的`ADMIN_TOKEN`。

#### 用户端Token
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"user","password":"user123"}'
```

复制`data.token`的值到`.env`文件的`USER_TOKEN`。

### 方式2：通过浏览器开发者工具

1. 打开浏览器，访问系统
2. 按F12打开开发者工具
3. 切换到**Network**标签
4. 执行登录操作
5. 找到登录请求，查看响应
6. 复制`data.token`的值

### 方式3：通过Postman

1. 创建一个新的POST请求
2. URL: `http://localhost:8080/api/v1/login`
3. Headers: `Content-Type: application/json`
4. Body:
   ```json
   {
     "username": "admin",
     "password": "admin123"
   }
   ```
5. 发送请求
6. 复制响应中`data.token`的值

## 更新Token后的操作

1. 编辑`.env`文件
2. 替换`ADMIN_TOKEN`或`USER_TOKEN`的值
3. 重新运行测试：
   ```bash
   make test
   ```

## 预防Token过期

1. **定期更新Token**：如果Token有效期短，可以定期重新登录获取新Token
2. **使用长有效期Token**：如果有多个环境，选择测试环境的Token（有效期更长）
3. **监控Token状态**：定期运行简单测试，检查Token是否即将过期

## 测试日志示例

### 正常情况
```
【请求】POST http://localhost:8080/api/v1/admin/wth/coin/add (使用管理端Token)
【请求Body】{"coinKey":"BTC","coinIcon":"..."}
【响应】Status: 200, Body: {"code":0,"message":"success"}
```

### Token过期
```
【请求】POST http://localhost:8080/api/v1/admin/wth/coin/add (使用管理端Token)
【请求Body】{"coinKey":"BTC","coinIcon":"..."}
【响应】Status: 200, Body: {"code":401,"message":"token expired"}

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

## 快速更新Token脚本

创建一个批处理文件`update_token.bat`：

```batch
@echo off
chcp 65001 >nul
echo 正在更新Token...

set /p ADMIN_TOKEN="请输入新的管理端Token: "
set /p USER_TOKEN="请输入新的用户端Token: "

REM 读取现有.env文件，替换Token
(for /f "delims=" %%a in (.env) do (
    set "line=%%a"
    setlocal enabledelayedexpansion
    echo !line:ADMIN_TOKEN=*=ADMIN_TOKEN=%ADMIN_TOKEN%!
    echo !line:USER_TOKEN=*=USER_TOKEN=%USER_TOKEN%!
    endlocal
)) > .env.new

move /y .env.new .env >nul
echo Token已更新！
pause
```

使用方法：
```bash
update_token.bat
```

## 总结

- ✅ 框架自动检测Token过期（code=401）
- ✅ Token过期时立即停止测试，给出明确错误
- ✅ 通过重新登录获取新Token
- ✅ 更新`.env`文件中的对应Token
- ✅ 重新运行测试即可
