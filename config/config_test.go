package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                                 App Config                                 */
/* -------------------------------------------------------------------------- */
func TestLoadAppConfig(t *testing.T) {
	LoadConfigs()
	assert.NotEmpty(t, App.Env)
	assert.NotEmpty(t, App.RootPath)
}

func TestLoadDBConfig(t *testing.T) {
	LoadConfigs()

	assert.NotEmpty(t, DB.Host)
	assert.NotEmpty(t, DB.Database)
	assert.NotEmpty(t, DB.Port)
	assert.NotEmpty(t, DB.Username)
	assert.NotEmpty(t, DB.Password)
}
