package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	APIBaseURL  string
	AdminToken  string // 管理端Token
	UserToken   string // 用户端Token
	DBHost      string
	DBPort      int
	DBUser      string
	DBPass      string
	DBName      string
	DBReadOnly  bool // 数据库是否只读（默认true）
	TestTimeout int
}

var Cfg Config

// Load 加载配置
func Load() {
	// 获取当前文件所在目录
	_, currentFilePath, _, _ := runtime.Caller(0)
	configDir := filepath.Dir(currentFilePath)
	envPath := filepath.Join(configDir, "..", ".env")

	// 尝试从多个位置加载.env文件
	envFiles := []string{
		envPath,
		".env",
		filepath.Join("..", ".env"),
	}

	var envLoadErr error
	for _, envFile := range envFiles {
		if err := godotenv.Load(envFile); err == nil {
			fmt.Printf("✅ .env file loaded from: %s\n", envFile)
			envLoadErr = nil
			break
		} else {
			envLoadErr = err
		}
	}

	if envLoadErr != nil {
		fmt.Printf("⚠️  Warning: .env file not found or error loading: %v\n", envLoadErr)
		fmt.Printf("📁 Searched paths:\n")
		for _, envFile := range envFiles {
			absPath, _ := filepath.Abs(envFile)
			fmt.Printf("   - %s\n", absPath)
		}
		fmt.Println("\n💡 To create .env file:")
		fmt.Println("   Windows: copy .env.example .env")
		fmt.Println("   Linux/Mac: cp .env.example .env")
		fmt.Println("   Then edit .env and fill in ADMIN_TOKEN and USER_TOKEN")
		panic("ADMIN_TOKEN and USER_TOKEN are required in .env file")
	}

	Cfg = Config{
		APIBaseURL:  getEnv("API_BASE_URL", "http://localhost:8080"),
		AdminToken:  getEnv("ADMIN_TOKEN", ""),
		UserToken:   getEnv("USER_TOKEN", ""),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnvInt("DB_PORT", 3306),
		DBUser:      getEnv("DB_USER", "root"),
		DBPass:      getEnv("DB_PASS", ""),
		DBName:      getEnv("DB_NAME", "my_stonks"),
		DBReadOnly:  getEnvBool("DB_READ_ONLY", true), // 默认只读
		TestTimeout: getEnvInt("TEST_TIMEOUT", 30),
	}

	// 验证必要的配置
	if Cfg.AdminToken == "" || Cfg.AdminToken == "your_admin_token_here" {
		panic("\n❌ ADMIN_TOKEN is required in .env file.\n" +
			"   Please edit .env and set a valid ADMIN_TOKEN\n" +
			"   Do not use the placeholder 'your_admin_token_here'")
	}
	if Cfg.UserToken == "" || Cfg.UserToken == "your_user_token_here" {
		panic("\n❌ USER_TOKEN is required in .env file.\n" +
			"   Please edit .env and set a valid USER_TOKEN\n" +
			"   Do not use the placeholder 'your_user_token_here'")
	}

	dbMode := "只读"
	if !Cfg.DBReadOnly {
		dbMode = "读写"
	}
	fmt.Printf("✅ Configuration loaded successfully (数据库模式: %s)\n", dbMode)
}

// GetDSN 获取数据库连接字符串
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return defaultVal
}
