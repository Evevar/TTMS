package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
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
