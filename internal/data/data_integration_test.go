// +build integration

package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestData_Integration_getClient(t *testing.T) {
	cfg := &testConfig{url: "http://127.0.0.1:9200"}
	result, err := getClient(cfg)

	require.NoError(t, err)
	assert.NotNil(t, result)
	t.Log(result)
}
