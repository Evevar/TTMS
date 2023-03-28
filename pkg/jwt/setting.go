package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	JWTSecret = "kangning"
	// JWTOverTime 先弄成200小时，之后改回来
	JWTOverTime = time.Hour * 200
)

type MyClaims struct {
	ID       int64
	UserType int32
	jwt.StandardClaims
}
