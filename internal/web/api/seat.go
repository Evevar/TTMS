package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/studio"
	"TTMS/pkg/jwt"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddSeatRequest struct {
	StudioId int64
	Row      int64
	Col      int64
	Status   int64
	Token    string
}

func AddSeat(c *gin.Context) {
	receive := &AddSeatRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.AddSeatResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.AddSeatRequest{
		StudioId: receive.StudioId,
		Row:      receive.Row,
		Col:      receive.Col,
		Status:   receive.Status,
	}
	resp, err := rpc.AddSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type GetAllSeatRequest struct {
	StudioId int64
	Current  int32
	PageSize int32
	Token    string
}

func GetAllSeat(c *gin.Context) {
	receive := &GetAllSeatRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.GetAllSeatResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.GetAllSeatRequest{
		StudioId: receive.StudioId,
		Current:  receive.Current,
		PageSize: receive.PageSize,
	}
	resp, err := rpc.GetAllSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type UpdateSeatRequest struct {
	StudioId int64
	Row      int64
	Col      int64
	Status   int64
	Token    string
}

func UpdateSeat(c *gin.Context) {
	receive := &UpdateSeatRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.UpdateSeatResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.UpdateSeatRequest{
		StudioId: receive.StudioId,
		Row:      receive.Row,
		Col:      receive.Col,
		Status:   receive.Status,
	}
	resp, err := rpc.UpdateSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type DeleteSeatRequest struct {
	StudioId int64
	Row      int64
	Col      int64
	Token    string
}

func DeleteSeat(c *gin.Context) {
	receive := &DeleteSeatRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.DeleteSeatResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.DeleteSeatRequest{
		StudioId: receive.StudioId,
		Row:      receive.Row,
		Col:      receive.Col,
	}
	resp, err := rpc.DeleteSeat(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
