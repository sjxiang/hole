package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	cfg, err := GetConfig()
	assert.Nil(t, err)
	if assert.NotNil(t, cfg) {
		assert.Equal(t, "127.0.0.1", cfg.Addr)
		assert.Equal(t, "5432", cfg.Port)
		assert.Equal(t, "", cfg.User)
		assert.Equal(t, "", cfg.Password)
		assert.Equal(t, "illa", cfg.Database)
	}
}
