@echo off
chcp 65001 >nul
echo 正在创建 .env 文件...
echo.

REM 复制.env.example到.env
if not exist .env.example (
    echo 错误：.env.example 文件不存在
    pause
    exit /b 1
)

copy .env.example .env >nul
echo .env 文件已创建！
echo.
echo ========================================
echo 重要提示
echo ========================================
echo.
echo 请编辑 .env 文件，填写以下信息：
echo.
echo 1. ADMIN_TOKEN - 管理端Token
echo    用于：/api/v1/admin/* 接口
echo    获取方式：登录接口或浏览器开发者工具
echo.
echo 2. USER_TOKEN - 用户端Token
echo    用于：/api/v1/wth/* 等用户端接口
echo    获取方式：登录接口或浏览器开发者工具
echo.
echo 3. 数据库配置
echo    DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME
echo.
echo Token格式示例（JWT）：
echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
echo.
echo 编辑完成后，运行以下命令开始测试：
echo   make test
echo.
pause
