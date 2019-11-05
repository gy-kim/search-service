package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gy-kim/search-service/logging"
)

const (
	headerKey = "x-access-token"
)

// Authentication contains verify token.
type Authentication interface {
	Verify(w http.ResponseWriter, r *http.Request) bool
}

// JWTAuth verify authentication
type JWTAuth struct {
	cfg Config
}

// NewJWTAuth creates and initialize Auth
func NewJWTAuth(cfg Config) *JWTAuth {
	return &JWTAuth{
		cfg: cfg,
	}
}

// Verify checks authentication
// header : x-access-token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o
func (a *JWTAuth) Verify(w http.ResponseWriter, r *http.Request) bool {
	header := r.Header.Get(headerKey)
	header = strings.TrimSpace(header)

	if header == "" {
		err := fmt.Errorf("Missing auth token")
		a.cfg.Logger().Error("[Auth] failed verify token. %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	token := &Token{}

	_, err := jwt.ParseWithClaims(header, token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

// Config is the configuration of list package.
type Config interface {
	Logger() logging.Logger
}
