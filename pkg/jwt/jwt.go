package jwt

import (
	"TTMS/kitex_gen/user"
	"TTMS/pkg/consts"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

const TokenExpireDuration = JWTOverTime

var MySecret = []byte(JWTSecret)

var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: consts.RedisPassword,
		DB:       0, // use default DB
	})
}

func GenToken(userInfo *user.User) (string, error) {
	c := MyClaims{
		userInfo.Id,
		userInfo.UserType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "kangning",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完成的编码后的字符串token
	tokenStr, err := token.SignedString(MySecret)
	if err != nil {
		log.Panicln(err)
	}
	//写入redis
	userBytes, err := json.Marshal(userInfo)
	if err != nil {
		log.Panicln(err)
	}
	client.Set(context.Background(), tokenStr, userBytes, JWTOverTime)
	return tokenStr, err
}

// 解析JWT.
func ParseToken(tokenString string) (*MyClaims, error) {
	if _, err := client.GetEx(context.Background(), tokenString, JWTOverTime).Result(); err != nil {
		return nil, errors.New("token失效")
	}
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i any, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("JWT 解析错误")
}
