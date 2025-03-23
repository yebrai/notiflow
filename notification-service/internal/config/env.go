package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvAsInt(key string, defaultValue int) int {
	valueStr := GetEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func GetEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := GetEnv(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func GetEnvAsBool(key string, defaultValue bool) bool {
	valueStr := GetEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Email    EmailConfig
}

type ServerConfig struct {
	Port            string
	GinMode         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type DatabaseConfig struct {
	Host         string
	Port         string
	Name         string
	User         string
	Password     string
	SSLMode      string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  time.Duration
}

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
}

func LoadConfig() Config {
	return Config{
		Server: ServerConfig{
			Port:            GetEnv("PORT", "8080"),
			GinMode:         GetEnv("GIN_MODE", "debug"),
			ReadTimeout:     GetEnvAsDuration("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout:    GetEnvAsDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
			IdleTimeout:     GetEnvAsDuration("SERVER_IDLE_TIMEOUT", 60*time.Second),
			ShutdownTimeout: GetEnvAsDuration("SERVER_SHUTDOWN_TIMEOUT", 5*time.Second),
		},
		Database: DatabaseConfig{
			Host:         GetEnv("DB_HOST", "localhost"),
			Port:         GetEnv("DB_PORT", "5432"),
			Name:         GetEnv("DB_NAME", "notification_db"),
			User:         GetEnv("DB_USER", "postgres"),
			Password:     GetEnv("DB_PASSWORD", "admin"),
			SSLMode:      GetEnv("DB_SSL_MODE", "disable"),
			MaxOpenConns: GetEnvAsInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns: GetEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			MaxLifetime:  GetEnvAsDuration("DB_CONN_MAX_LIFETIME", 15*time.Minute),
		},
		Email: EmailConfig{
			SMTPHost:     GetEnv("SMTP_HOST", ""),
			SMTPPort:     GetEnv("SMTP_PORT", "587"),
			SMTPUsername: GetEnv("SMTP_USERNAME", ""),
			SMTPPassword: GetEnv("SMTP_PASSWORD", ""),
			FromEmail:    GetEnv("SMTP_FROM", "notifications@example.com"),
		},
	}
}

// Validate valida que la configuración tenga valores válidos
func (c *Config) Validate() []string {
	var missingVars []string

	if c.Database.Host == "" {
		missingVars = append(missingVars, "DB_HOST")
	}

	return missingVars
}
