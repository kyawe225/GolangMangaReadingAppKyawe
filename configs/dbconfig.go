package configs

import (
	"os"
)

type DbConfigs struct {
	Password    string
	Port        string
	User        string
	Host        string
	SslDisabled string
	Database    string
}

// GetDbConfigs returns database configuration with default values or from environment
func GetDbConfigs() *DbConfigs {
	return &DbConfigs{
		Password:    getEnvOrDefault("DB_PASSWORD", "kyawe"),
		Port:        getEnvOrDefault("DB_PORT", "5432"),
		User:        getEnvOrDefault("DB_USER", "postgres"),
		Host:        getEnvOrDefault("DB_HOST", "db"),
		SslDisabled: getEnvOrDefault("DB_SSL_MODE", "disable"),
		Database:    getEnvOrDefault("DB_DATABASE", "manga_database"),
	}
}

// GetConnectionString returns formatted connection string for PostgreSQL
func (db *DbConfigs) GetConnectionString() string {
	return "user=" + db.User + " password=" + db.Password + " host=" + db.Host + " port=" + db.Port + " database=" + db.Database + " sslmode=" + db.SslDisabled
}

// getEnvOrDefault gets environment variable or returns default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
