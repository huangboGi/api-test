@echo off
chcp 65001 >nul
cls
echo ========================================
echo 从 settings.yml 同步数据库配置
echo ========================================
echo.

REM 检查 settings.yml 文件
set SETTINGS_PATH=d:\project\my_stonks_background_dev\config\settings.yml
if not exist "%SETTINGS_PATH%" (
    echo ❌ 错误：找不到 settings.yml 文件
    echo    路径：%SETTINGS_PATH%
    echo.
    pause
    exit /b 1
)

echo ✅ 找到 settings.yml 文件
echo    路径：%SETTINGS_PATH%
echo.

REM 解析 settings.yml 并提取数据库配置
echo 正在解析配置文件...

for /f "tokens=2 delims=: " %%a in ('findstr /C:"source:" "%SETTINGS_PATH%" ^| findstr /C:"*"":^') do (
    set DSN=%%b
)

if "%DSN%"=="" (
    echo ❌ 错误：无法从 settings.yml 中提取数据库配置
    pause
    exit /b 1
)

echo ✅ 提取到数据库连接字符串
echo.

REM 解析 DSN 字符串
REM 格式: admin:password@tcp(host:port)/database?charset=utf8&parseTime=True&loc=Local&timeout=1000ms&compress=true

REM 提取用户名和密码 (admin:password@tcp...)
for /f "tokens=1 delims=@" %%a in ("%DSN%") do (
    set USER_PASS=%%a
)

for /f "tokens=1,2 delims=:" %%a in ("%USER_PASS%") do (
    set DB_USER=%%a
    set DB_PASS=%%b
)

REM 提取主机、端口、数据库名 (tcp(host:port)/database?...)
for /f "tokens=2 delims=/" %%a in ("%DSN%") do (
    set HOST_DB_PART=%%a
)

REM 提取主机:port
set HOST_DB_PART=%HOST_DB_PART:tcp(=%
set HOST_DB_PART=%HOST_DB_PART:)=%

for /f "tokens=1,2 delims=:)" %%a in ("%HOST_DB_PART%") do (
    set DB_HOST=%%a
    set DB_PORT=%%b
)

echo 数据库配置解析结果：
echo   主机: %DB_HOST%
echo   端口: %DB_PORT%
echo   用户: %DB_USER%
echo   数据库: 将自动识别
echo.

REM 生成 .env 文件
echo 正在生成 .env 文件...

(
    echo # API基础配置
    echo API_BASE_URL=http://localhost:8080
    echo.
    echo # 管理端 Token ^(用于 /api/v1/admin/* 接口^)
    echo ADMIN_TOKEN=your_admin_token_here
    echo.
    echo # 用户端 Token ^(用于 /api/v1/wth/* 等用户端接口^)
    echo USER_TOKEN=your_user_token_here
    echo.
    echo # 数据库配置（从 settings.yml 自动同步^)
    echo DB_HOST=%DB_HOST%
    echo DB_PORT=%DB_PORT%
    echo DB_USER=%DB_USER%
    echo DB_PASS=%DB_PASS%
    echo DB_NAME=stonks
    echo.
    echo # 数据库只读模式（推荐：true^)
    echo # true: 只读模式，禁止所有增删改操作（推荐用于测试^）
    echo # false: 读写模式，允许所有操作（不推荐^）
    echo DB_READ_ONLY=true
    echo.
    echo # 测试配置
    echo TEST_TIMEOUT=30
) > .env

echo.
echo ========================================
echo ✅ 数据库配置同步成功！
echo ========================================
echo.
echo 数据库配置：
echo   主机: %DB_HOST%
echo   端口: %DB_PORT%
echo   用户: %DB_USER%
echo   数据库: stonks
echo   模式: 只读 (DB_READ_ONLY=true)
echo.
echo ⚠️  接下来的操作：
echo 1. 编辑 .env 文件
echo 2. 填写真实的 ADMIN_TOKEN 和 USER_TOKEN
echo 3. 运行测试: make test
echo ========================================
pause
