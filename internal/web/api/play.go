package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/play"
	"TTMS/pkg/jwt"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddPlay(c *gin.Context) {
	req := &play.AddPlayRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.AddPlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.AddPlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func UpdatePlay(c *gin.Context) {
	req := &play.UpdatePlayRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.UpdatePlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.UpdatePlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func DeletePlay(c *gin.Context) {
	req := &play.DeletePlayRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.DeletePlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.DeletePlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func GetAllPlay(c *gin.Context) {
	req := &play.GetAllPlayRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.AddPlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.GetAllPlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func AddSchedule(c *gin.Context) {
	req := &play.AddScheduleRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.AddPlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.AddSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func UpdateSchedule(c *gin.Context) {
	req := &play.UpdateScheduleRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.UpdateScheduleResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.UpdateSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func DeleteSchedule(c *gin.Context) {
	req := &play.DeleteScheduleRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.DeleteScheduleResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.DeleteSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
func GetAllSchedule(c *gin.Context) {
	req := &play.GetAllScheduleRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.GetAllScheduleResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.GetAllSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
