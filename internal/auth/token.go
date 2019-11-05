package auth

import jwt "github.com/dgrijalva/jwt-go"

// Token is strucure token.
type Token struct {
	*jwt.StandardClaims
}
