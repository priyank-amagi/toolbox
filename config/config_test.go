package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJsonApiHost(t *testing.T) {
	testJsonHost := "json_api_host_1"
	err := os.Setenv("JSON_API_HOST", testJsonHost)
	assert.Equal(t, nil, err)
	assert.Equal(t, testJsonHost, GetJsonApiHost())

}
