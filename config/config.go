package config

import "os"

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}
type AppConfig struct {
	ENV                string
	AppName            string
	JWTSecret          []byte
	JWTExpireInMinutes int64
	DBConfig           dbConfig
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var Config = AppConfig{
	ENV:                getENV("ENV", "deploy"),
	AppName:            "apps",
	JWTSecret:          []byte("very-secret"),
	JWTExpireInMinutes: 10000,
	DBConfig: dbConfig{
		Host:     getENV("DB_HOST", "localhost"),
		User:     getENV("DB_USER", "postgres"),
		Password: getENV("DB_PASSWORD", "postgres"),
		DBName:   getENV("DB_NAME", "wallet_db_aaronlee"),
		Port:     getENV("DB_PORT", "5432"),
	},
}
