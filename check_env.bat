@echo off
chcp 65001 >nul
cls
echo ========================================
echo 检查 .env 配置
echo ========================================
echo.

if not exist .env (
    echo ❌ .env 文件不存在
    echo.
    echo 解决方法：
    echo   运行：quick_setup.bat
    echo   或执行：copy .env.example .env
    echo.
    pause
    exit /b 1
)

echo ✅ .env 文件存在
echo.
echo 当前配置：
echo.

type .env
echo.

echo ========================================
echo 验证配置
echo ========================================
echo.

REM 检查ADMIN_TOKEN
findstr /C:"ADMIN_TOKEN=" .env | findstr /C:"your_admin_token_here" >nul
if errorlevel 1 (
    echo ✅ ADMIN_TOKEN 已配置
) else (
    echo ❌ ADMIN_TOKEN 仍为占位符，需要填写真实Token
)

REM 检查USER_TOKEN
findstr /C:"USER_TOKEN=" .env | findstr /C:"your_user_token_here" >nul
if errorlevel 1 (
    echo ✅ USER_TOKEN 已配置
) else (
    echo ❌ USER_TOKEN 仍为占位符，需要填写真实Token
)

echo.
echo ========================================
pause
