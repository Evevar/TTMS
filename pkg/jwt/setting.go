package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	ID       int64
	UserType int32
	jwt.StandardClaims
}
