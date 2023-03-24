package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/studio"
	"TTMS/pkg/jwt"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddSeat(c *gin.Context) {
	req := &studio.AddSeatRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.AddSeatResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.AddSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func GetAllSeat(c *gin.Context) {
	req := &studio.GetAllSeatRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.GetAllSeatResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	fmt.Println(req)
	resp, err := rpc.GetAllSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func UpdateSeat(c *gin.Context) {
	req := &studio.UpdateSeatRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.UpdateSeatResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.UpdateSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func DeleteSeat(c *gin.Context) {
	req := &studio.DeleteSeatRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.DeleteSeatResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.DeleteSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
