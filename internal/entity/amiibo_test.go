package entity

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAmiibo_Build(t *testing.T) {
	viper.Set("key-retail", "../../pkg/amiibo/key_retail.bin")
	amiibo := NewAmiibo("Mario")
	amiibo.ID = "0x0000000000000002"
	assert.NoError(t, amiibo.Build())
	assert.NotEmpty(t, amiibo.Path)
}
