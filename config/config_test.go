package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	// InitEnv()

	assert.NotEqual(t, Env.DBHost, "", "`config.Env.DBHost` should not be empty.")
}
