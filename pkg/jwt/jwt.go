package jwt

import (
	"TTMS/configs/consts"
	"TTMS/kitex_gen/user"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
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
	num, err1 := client.LRem(context.Background(), "black", 0, tokenStr).Result()
	log.Println("client.LRem return ", num, err1)
	return tokenStr, err
}

// 解析JWT.
func ParseToken(tokenString string) (*MyClaims, error) {
	arr, _ := client.LRange(context.Background(), "black", 0, -1).Result()
	for _, v := range arr {
		if strings.Compare(v, tokenString) == 0 { //该token在黑名单中
			return nil, errors.New("更改密码后需重新登录")
		}
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

// 废弃JWT(更改密码后)
func DiscardToken(tokenString string) {
	_, err := client.LPush(context.Background(), "black", tokenString).Result()
	if err != nil {
		log.Println(err)
	}
}
