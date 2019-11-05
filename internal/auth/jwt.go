package auth

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	headerKey = "x-access-token"
)

// Verify checks authentication
// header : x-access-token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o
func Verify(w http.ResponseWriter, r *http.Request) bool {
	header := r.Header.Get(headerKey)
	header = strings.TrimSpace(header)

	if header == "" {
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
