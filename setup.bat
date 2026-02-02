@echo off
chcp 65001 >nul
echo ========================================
echo MyStonks API 测试项目配置向导
echo ========================================
echo.

REM 检查.env文件是否存在
if exist .env (
    echo .env 文件已存在
    echo.
    choice /C YN /M "是否要重新配置？"
    if errorlevel 2 goto end
    echo.
)

echo 正在创建/更新 .env 文件...
echo.

REM 获取用户输入
set /p API_BASE_URL="请输入API地址 [默认: http://localhost:8080]: "
if "%API_BASE_URL%"=="" set API_BASE_URL=http://localhost:8080

echo.
echo ========================================
echo Token配置
echo ========================================
echo.
echo Token可以从以下方式获取：
echo 1. 登录接口返回的 data.token
echo 2. 浏览器开发者工具（F12）的Network标签中
echo.
echo Token格式示例（JWT）：
echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTk5MjAwMDB9.XXXX
echo.

set /p ADMIN_TOKEN="请输入管理端Token (ADMIN_TOKEN): "

if "%ADMIN_TOKEN%"=="" (
    echo 错误：管理端Token不能为空
    pause
    exit /b 1
)

echo.
set /p USER_TOKEN="请输入用户端Token (USER_TOKEN): "

if "%USER_TOKEN%"=="" (
    echo 错误：用户端Token不能为空
    pause
    exit /b 1
)

echo.
echo ========================================
echo 数据库配置
echo ========================================
echo.

set /p DB_HOST="请输入数据库地址 [默认: localhost]: "
if "%DB_HOST%"=="" set DB_HOST=localhost

set /p DB_PORT="请输入数据库端口 [默认: 3306]: "
if "%DB_PORT%"=="" set DB_PORT=3306

set /p DB_USER="请输入数据库用户名 [默认: root]: "
if "%DB_USER%"=="" set DB_USER=root

set /p DB_PASS="请输入数据库密码: "

set /p DB_NAME="请输入数据库名 [默认: my_stonks]: "
if "%DB_NAME%"=="" set DB_NAME=my_stonks

echo.
REM 写入.env文件
(
    echo # API基础配置
    echo API_BASE_URL=%API_BASE_URL%
    echo.
    echo # 管理端 Token ^(用于 /api/v1/admin/* 接口^)
    echo ADMIN_TOKEN=%ADMIN_TOKEN%
    echo.
    echo # 用户端 Token ^(用于 /api/v1/wth/* 等用户端接口^)
    echo USER_TOKEN=%USER_TOKEN%
    echo.
    echo # 数据库配置
    echo DB_HOST=%DB_HOST%
    echo DB_PORT=%DB_PORT%
    echo DB_USER=%DB_USER%
    echo DB_PASS=%DB_PASS%
    echo DB_NAME=%DB_NAME%
    echo.
    echo # 测试配置
    echo TEST_TIMEOUT=30
) > .env

echo.
echo ========================================
echo 配置完成！
echo ========================================
echo.
echo .env 文件已创建，包含以下配置：
echo - API_BASE_URL: %API_BASE_URL%
echo - ADMIN_TOKEN: 已设置（长度: %TOKEN_LENGTH%）
echo - USER_TOKEN: 已设置（长度: %TOKEN_LENGTH%）
echo - 数据库: %DB_USER%@%DB_HOST%:%DB_PORT%/%DB_NAME%
echo.
echo 现在可以运行测试了：
echo   make test
echo.

:end
pause
