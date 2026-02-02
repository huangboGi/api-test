@echo off
chcp 65001 >nul
cls
echo ========================================
echo MyStonks API 测试项目快速配置
echo ========================================
echo.

REM 检查.env文件
if exist .env (
    echo ✅ .env 文件已存在
    echo.
    echo 当前配置内容：
    type .env | findstr /V "TOKEN" | findstr /V "PASSWORD"
    echo.
    choice /C YN /M "是否要重新配置"
    if errorlevel 2 goto end
    echo.
)

REM 复制.env.example
if not exist .env.example (
    echo ❌ 错误：.env.example 文件不存在
    pause
    exit /b 1
)

copy .env.example .env >nul
echo ✅ .env 文件已创建
echo.
echo ========================================
echo 重要说明
echo ========================================
echo.
echo 📝 请编辑 .env 文件，填写以下必填项：
echo.
echo 1. ADMIN_TOKEN  - 管理端Token ⚠️ 必填
echo    用于：/api/v1/admin/* 接口
echo.
echo 2. USER_TOKEN   - 用户端Token ⚠️ 必填
echo    用于：/api/v1/wth/* 等用户端接口
echo.
echo 3. 数据库配置 - 根据实际情况填写
echo    DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME
echo.
echo ========================================
echo 如何获取Token
echo ========================================
echo.
echo 方式1：登录接口获取
echo   curl -X POST http://localhost:8080/api/v1/login ^
echo     -H "Content-Type: application/json" ^
echo     -d "{\"username\":\"admin\",\"password\":\"admin123\"}"
echo.
echo 方式2：浏览器获取
echo   1. 打开浏览器开发者工具（F12）
echo   2. 切换到 Network 标签
echo   3. 登录系统
echo   4. 查看登录响应中的 data.token
echo.
echo ========================================
echo 接下来的操作
echo ========================================
echo.
echo 1. 使用记事本编辑 .env 文件
echo.
echo 2. 填写 ADMIN_TOKEN 和 USER_TOKEN
echo.
echo 3. 保存文件
echo.
echo 4. 运行测试：make test
echo.
echo ========================================

REM 询问是否现在打开.env文件编辑
choice /C YN /M "是否现在打开 .env 文件进行编辑？"
if errorlevel 1 (
    start notepad .env
    echo.
    echo ⏳ 记事本已打开，请填写Token后保存...
    echo.
    pause
) else (
    echo.
    echo 💡 提示：你可以随时编辑 .env 文件
    echo.
    pause
)

:end
echo.
echo ✅ 配置完成！
echo.
echo 现在可以运行测试了：
echo   make test
echo.
pause
