package config

import "os"

// Env set configuration struct
var Env Environment

// Environment struct for storing configuration settings.
type Environment struct {
	DatabaseHost string
	DatabasePort string
	DatabaseName string
	DatabaseUser string
	DatabasePass string
}

// LoadConfig loads configuration settings to be used through out the app.
func LoadConfig() {
	Env.DatabaseHost = GetEnv("DATABASE_HOST", "0.0.0.0")
	Env.DatabasePort = GetEnv("DATABASE_PORT", "5432")
	Env.DatabaseName = GetEnv("DATABASE_NAME", "80sMixtapeAPI")
	Env.DatabaseUser = GetEnv("DATABASE_USER", "database123")
	Env.DatabasePass = GetEnv("DATABASE_PASS", "database123")
}

// GetEnv get the environmental variable for a key or return the specified default value.
func GetEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}
	return val
}
