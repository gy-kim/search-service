// +build integration

package data

// import (
// 	"context"
// 	"fmt"
// 	"strings"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestDAO_Integration_InsertProduct(t *testing.T) {
// 	ctx := context.Background()
// 	cfg := &testConfig{
// 		url: "http://127.0.0.1:9200",
// 	}

// 	id, err := uuid.NewRandom()
// 	require.NoError(t, err)

// 	t.Logf("uuid:%s", id.String())

// 	product := &Product{
// 		ID:    id.String(),
// 		Type:  "black_shoes",
// 		Brand: "adidas",
// 		Name:  fmt.Sprintf("adidas_%s", strings.Split(id.String(), "-")[0]),
// 	}

// 	dao := NewDAO(cfg)
// 	err = dao.InsertProduct(ctx, product)

// 	assert.NoError(t, err)
// }
