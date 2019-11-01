// +build integration

package data

import (
	"testing"

	"github.com/gy-kim/search-service/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestData_Integration_getClient(t *testing.T) {
	result, err := getClient(&testConfig{})
	require.NoError(t, err)
	assert.NotNil(t, result)
	t.Log(result)
}

type testConfig struct{}

func (t *testConfig) Logger() logging.Logger {
	return &logging.LoggerStdOut{}
}

func (t *testConfig) DataURL() string {
	return "http://elasticsearch:9200"
}
