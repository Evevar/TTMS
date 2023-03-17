package api

import (
	"TTMS/internal/web/rpc"
	user "TTMS/kitex_gen/user"
	"TTMS/pkg/jwt"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	req := &user.CreateUserRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	//_, err := jwt.ParseToken(req.Token)
	//if err != nil {
	//	c.JSON(http.StatusOK, user.CreateUserResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
	//	return
	//}

	resp, err := rpc.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, resp)
}
func UserLogin(c *gin.Context) {
	req := &user.UserLoginRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	fmt.Println(" LoginUser ", req)
	resp, err := rpc.UserLogin(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	Token, err := jwt.GenToken(resp.UserInfo)
	if err != nil {
		log.Println("JWT 生成错误", err)
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMessage = err.Error()
	} else {
		resp.Token = Token
	}
	resp.UserInfo = nil
	log.Println("resp = ", resp)
	c.JSON(http.StatusOK, resp)
}
func GetAllUser(c *gin.Context) {
	req := &user.GetAllUserRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	//token验证
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, user.GetAllUserResponse{BaseResp: &user.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	//接收resp
	resp, err := rpc.GetAllUser(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, resp)
}
