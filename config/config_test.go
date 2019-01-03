package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	InitEnv()

	assert.NotEqual(t, Env.DBHost, "", "`config.Env.DBHost` should not be empty.")
	assert.NotEqual(t, Env.DBPort, "", "`config.Env.DBPort` should not be empty.")
	assert.NotEqual(t, Env.DBName, "", "`config.Env.DBName` should not be empty.")
	assert.NotEqual(t, Env.DBUser, "", "`config.Env.DBUser` should not be empty.")
	assert.NotEqual(t, Env.DBPass, "", "`config.Env.DBPass` should not be empty.")
	assert.NotEqual(t, Env.DBConn, "", "`config.Env.DBConn` should not be empty.")
}

func TestConfigEnvVars(t *testing.T) {

	expected := "default test value"
	actual := GetEnvVar("TEST_KEY", "default test value")
	description := "`config.GetEnvVar()` should output default if no env var exists."
	assert.Equal(t, actual, expected, description)
}
