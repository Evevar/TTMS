package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JWTSecret = "kangning"

	JWTOverTime = time.Hour * 2
)

type MyClaims struct {
	ID       int64
	UserType int32
	jwt.StandardClaims
}
