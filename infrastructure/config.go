package infrastructure

import (
	"fmt"
	"os"
)

// DBConfig 構造体
type DBConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

// LoadDBConfig 環境変数をロードして設定を取得
func LoadDBConfig() *DBConfig {

	return &DBConfig{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("MYSQL_ROOT_PASSWORD", ""),
		DBName:     getEnv("MYSQL_DATABASE", "db"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
	}
}

func (c *DBConfig) GetDSN() string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	)

	return dsn
}

// DBConfig 構造体
type AUTH0Config struct {
	Domain   string `json:"domain"`
	ClientID string `json:"clientId"`
	Audience string `json:"audience"`
}

func LoadAUTH0Config() *AUTH0Config {
	return &AUTH0Config{
		Domain:   getEnv("AUTH0_DOMAIN", ""),
		ClientID: getEnv("AUTH0_WEBAPP_CLIENT_ID", ""),
		Audience: getEnv("AUTH0_API_AUDIENCE", ""),
	}
}

// getEnv 環境変数を取得（デフォルト値あり）
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
