package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gy-kim/search-service/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWTAuth_Verify(t *testing.T) {
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o"
	url := "/v1/products"
	scenarios := []struct {
		desc           string
		inRequest      func() *http.Request
		expectedResult bool
	}{
		{
			desc: "happy path",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", url, nil)
				require.NoError(t, err)
				req.Header.Add(headerKey, validToken)

				return req
			},
			expectedResult: true,
		},
		{
			desc: "missing token",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", url, nil)
				require.NoError(t, err)

				return req
			},
			expectedResult: false,
		},
		{
			desc: "invalid token",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", url, nil)
				require.NoError(t, err)
				req.Header.Add(headerKey, "invalid token")

				return req
			},
			expectedResult: false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			response := httptest.NewRecorder()
			jwtAuth := NewJWTAuth(&testConfig{})

			result := jwtAuth.Verify(response, s.inRequest())
			assert.Equal(t, s.expectedResult, result)
		})
	}
}

type testConfig struct{}

func (t *testConfig) Logger() logging.Logger {
	return &logging.LoggerStdOut{}
}
