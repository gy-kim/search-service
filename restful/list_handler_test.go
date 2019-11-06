package restful

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gy-kim/search-service/internal/data"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListHandler_ServeHTTP(t *testing.T) {
	url := "/v1/products?q=black_shoes&filter=brand:adidas&sort=name"
	scenarios := []struct {
		desc           string
		inRequest      func() *http.Request
		inService      func() *MockListService
		expectedStatus int
	}{
		{
			desc: "happy path",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", url, nil)
				require.NoError(t, err)
				return req
			},
			inService: func() *MockListService {
				mockResult := []*data.Product{
					{
						ID:    "product_id_1",
						Type:  "black_shoes",
						Brand: "adidas",
						Name:  "adidas_black_shoes_product_id_1",
					},
					{
						ID:    "product_id_2",
						Type:  "black_shoes",
						Brand: "adidas",
						Name:  "adidas_black_shoes_product_id_2",
					},
				}
				mck := &MockListService{}
				mck.On("Do", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResult, int64(2), nil).Once()
				return mck
			},
			expectedStatus: http.StatusOK,
		},
		{
			desc: "service failed",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", url, nil)
				require.NoError(t, err)
				return req
			},
			inService: func() *MockListService {
				mck := &MockListService{}
				mck.On("Do", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, int64(0), errors.New("something error")).Once()
				return mck
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			mservice := s.inService()
			handler := NewListHandler(&testConfig{}, mservice)
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, s.inRequest())

			require.Equal(t, s.expectedStatus, response.Code)
			assert.True(t, mservice.AssertExpectations(t))
		})
	}
}
