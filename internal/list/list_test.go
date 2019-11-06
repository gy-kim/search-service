package list

import (
	"context"
	"errors"
	"testing"

	"github.com/gy-kim/search-service/internal/data"
	"github.com/gy-kim/search-service/logging"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLister_Do(t *testing.T) {
	ctx := context.Background()
	happyPathProducts := []*data.Product{
		{
			ID:    "first-id",
			Type:  "black_shoes",
			Brand: "adidas",
			Name:  "black_shoes_adidas",
		},
	}

	scenarios := []struct {
		desc           string
		mdao           func() *mockListDAO
		expectedError  bool
		expectedResult []*data.Product
	}{
		{
			desc: "happy path",
			mdao: func() *mockListDAO {
				dao := &mockListDAO{}
				dao.On("GetProducts", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(happyPathProducts, int64(1), nil).Once()
				return dao
			},
			expectedError:  false,
			expectedResult: happyPathProducts,
		},
		{
			desc: "error load data",
			mdao: func() *mockListDAO {
				dao := &mockListDAO{}
				dao.On("GetProducts", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), errors.New("something error")).Once()
				return dao
			},
			expectedError:  true,
			expectedResult: nil,
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			dao := s.mdao()

			lister := &Lister{
				dao: dao,
				cfg: &testConfig{},
			}

			result, _, err := lister.Do(ctx, "", nil, nil, 0)

			require.Equal(t, s.expectedError, err != nil)
			assert.Equal(t, s.expectedResult, result)
			assert.True(t, dao.AssertExpectations(t))
		})
	}
}

type testConfig struct {
	url string
}

func (t *testConfig) Logger() logging.Logger { return &logging.LoggerStdOut{} }

func (t *testConfig) DataURL() string { return t.url }
