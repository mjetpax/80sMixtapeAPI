package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Env set configuration struct
var Env Environment

// Environment struct for storing configuration settings.
type Environment struct {
	DB     *sql.DB
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
	DBConn string
}

// InitEnv initializes configuration settings to be used through out the app.
func InitEnv() {
	Env.DBHost = GetEnvVar("DATABASE_HOST", "0.0.0.0")
	Env.DBPort = GetEnvVar("DATABASE_PORT", "5432")
	Env.DBName = GetEnvVar("DATABASE_NAME", "80sMixtapeAPI")
	Env.DBUser = GetEnvVar("DATABASE_USER", "database123")
	Env.DBPass = GetEnvVar("DATABASE_PASS", "database123")
	Env.DBConn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Env.DBHost, Env.DBPort, Env.DBUser, Env.DBPass, Env.DBName)

	db, err := sql.Open("postgres", Env.DBConn)
	if err != nil {
		log.Panic(err)
	}

	Env.DB = db
}

// GetEnvVar get the environmental variable for a key or return the specified default value.
func GetEnvVar(key, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}
	return val
}
