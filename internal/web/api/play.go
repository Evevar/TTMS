package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/play"
	"TTMS/pkg/jwt"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddPlayRequest struct {
	Name      string
	Type      uint32
	Area      string
	Rating    uint32
	Duration  string
	StartDate string
	EndDate   string
	Price     int64
	Token     string
}

func AddPlay(c *gin.Context) {
	receive := &AddPlayRequest{}
	if err := c.Bind(receive); err != nil {
		log.Println("err = ", err, " receive = ", receive)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.AddPlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	req := &play.AddPlayRequest{
		Name:      receive.Name,
		Type:      receive.Type,
		Area:      receive.Area,
		Rating:    receive.Rating,
		Duration:  receive.Duration,
		StartDate: receive.StartDate,
		EndDate:   receive.EndDate,
		Price:     receive.Price,
	}
	resp, err := rpc.AddPlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type UpdatePlayRequest struct {
	Name      string
	Type      uint32
	Area      string
	Rating    uint32
	Duration  string
	StartDate string
	EndDate   string
	Price     int64
	Token     string
}

func UpdatePlay(c *gin.Context) {
	receive := &UpdatePlayRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.UpdatePlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &play.UpdatePlayRequest{
		Name:      receive.Name,
		Type:      receive.Type,
		Area:      receive.Area,
		Rating:    receive.Rating,
		Duration:  receive.Duration,
		StartDate: receive.StartDate,
		EndDate:   receive.EndDate,
		Price:     receive.Price,
	}
	resp, err := rpc.UpdatePlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type DeletePlayRequest struct {
	Id    int64
	Token string
}

func DeletePlay(c *gin.Context) {
	receive := &DeletePlayRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.DeletePlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &play.DeletePlayRequest{
		Id: receive.Id,
	}
	resp, err := rpc.DeletePlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type GetAllPlayRequest struct {
	Current  int32
	PageSize int32
	Token    string
}

func GetAllPlay(c *gin.Context) {
	receive := &GetAllPlayRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.GetAllPlayResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &play.GetAllPlayRequest{
		Current:  receive.Current,
		PageSize: receive.PageSize,
	}
	resp, err := rpc.GetAllPlay(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type AddScheduleRequest struct {
	PlayId   int64
	StudioId int64
	ShowTime string
	Token    string
}

func AddSchedule(c *gin.Context) {
	receive := &AddScheduleRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.AddScheduleResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &play.AddScheduleRequest{
		PlayId:   receive.PlayId,
		StudioId: receive.StudioId,
		ShowTime: receive.ShowTime,
	}
	resp, err := rpc.AddSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type UpdateScheduleRequest struct {
	PlayId   int64
	StudioId int64
	ShowTime string
	Token    string
}

func UpdateSchedule(c *gin.Context) {
	receive := &UpdateScheduleRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.UpdateScheduleResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &play.UpdateScheduleRequest{
		PlayId:   receive.PlayId,
		StudioId: receive.StudioId,
		ShowTime: receive.ShowTime,
	}
	resp, err := rpc.UpdateSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type DeleteScheduleRequest struct {
	Id    int64
	Token string
}

func DeleteSchedule(c *gin.Context) {
	receive := &DeleteScheduleRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.DeleteScheduleResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &play.DeleteScheduleRequest{
		Id: receive.Id,
	}
	resp, err := rpc.DeleteSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type GetAllScheduleRequest struct {
	Current  int32
	PageSize int32
	Token    string
}

func GetAllSchedule(c *gin.Context) {
	receive := &GetAllScheduleRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, play.GetAllScheduleResponse{BaseResp: &play.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &play.GetAllScheduleRequest{
		Current:  receive.Current,
		PageSize: receive.PageSize,
	}
	resp, err := rpc.GetAllSchedule(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
