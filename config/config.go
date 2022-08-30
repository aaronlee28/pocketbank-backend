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
	ENV:                getENV("ENV", "not testing"),
	AppName:            "apps",
	JWTSecret:          []byte("very-secret"),
	JWTExpireInMinutes: 10000,
	DBConfig: dbConfig{
		Host:     getENV("DB_HOST", "localhost"),
		User:     getENV("DB_USER", "postgres"),
		Password: getENV("DB_PASSWORD", "postgres"),
		DBName:   getENV("DB_NAME", "wallet_db_aaronlee"),
		Port:     getENV("DB_PORT", "5432"),
		//Host:     getENV("DB_HOST", "ec2-34-199-68-114.compute-1.amazonaws.com"),
		//User:     getENV("DB_USER", "buwbstfwugrqur"),
		//Password: getENV("DB_PASSWORD", "93d2bee55e0f6ea11d74e3812ca4a9cbf0696100de2aaf3e09f094463c5fc031"),
		//DBName:   getENV("DB_NAME", "depet60r4b0q6n"),
		//Port:     getENV("DB_PORT", "5432"),
	},
}
