package config

import (
	"os"
	"sync"
)

type Config struct {
	Server  Server
	DB      DBConfig
	Mailgun Mailgun
	JWT     JWT
}

type DBConfig struct {
	Name     string
	User     string
	Password string
	Port     string
	Host     string
}

type Server struct {
	Port string
}

type Mailgun struct {
	Domain          string
	Private_Key     string
	Sender_email    string
	Recipient_email string
}

type JWT struct {
	Secret string
}

var (
	once   sync.Once
	config Config
)

func GetConfig() *Config {
	// goroutine実行中でも一度だけ実行される
	once.Do(func() {
		// DBConfig
		config.DB.Name = getEnv("DB_NAME", "")
		config.DB.User = getEnv("DB_USER", "")
		config.DB.Password = getEnv("DB_PASSWORD", "")
		config.DB.Port = getEnv("DB_PORT", "")
		config.DB.Host = getEnv("DB_HOST", "")

		// Server
		config.Server.Port = getEnv("SERVER_PORT", "")

		// Mailgun
		config.Mailgun.Domain = getEnv("MAILGUN_DOMAIN", "")
		config.Mailgun.Private_Key = getEnv("MAILGUN_PRIVATE_API_KEY", "")
		config.Mailgun.Sender_email = getEnv("SENDER_EMAIL", "")
		config.Mailgun.Recipient_email = getEnv("RECIPIENT_EMAIL", "")

		// JWT
		config.JWT.Secret = getEnv("JWT_SECRET", "")
	})
	return &config
}

// getEnv は環境変数を取得し、指定されたデフォルト値を返す
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
