package config

import "os"

type AppConfig struct {
	ENV                string
	AppName            string
	JWTSecret          []byte
	JWTExpireInMinutes int64
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
}
