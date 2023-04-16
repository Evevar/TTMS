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

type AddStudioRequest struct {
	Name      string
	RowsCount int64
	ColsCount int64
	Token     string
}

func AddStudio(c *gin.Context) {
	receive := &AddStudioRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.AddStudioResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.AddStudioRequest{
		Name:      receive.Name,
		RowsCount: receive.RowsCount,
		ColsCount: receive.ColsCount,
	}
	resp, err := rpc.AddStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type GetAllStudioRequest struct {
	Current  int32
	PageSize int32
	Token    string
}

func GetAllStudio(c *gin.Context) {
	receive := &GetAllStudioRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.GetAllStudioResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.GetAllStudioRequest{
		Current:  receive.Current,
		PageSize: receive.PageSize,
	}
	resp, err := rpc.GetAllStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type GetStudioRequest struct {
	Id    int64
	Token string
}

func GetStudio(c *gin.Context) {
	receive := &GetStudioRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.GetStudioResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.GetStudioRequest{
		Id: receive.Id,
	}
	resp, err := rpc.GetStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type UpdateStudioRequest struct {
	Id        int64
	Name      string
	RowsCount int64
	ColsCount int64
	Token     string
}

func UpdateStudio(c *gin.Context) {
	receive := &UpdateStudioRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.UpdateStudioResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.UpdateStudioRequest{
		Id:        receive.Id,
		Name:      receive.Name,
		RowsCount: receive.RowsCount,
		ColsCount: receive.ColsCount,
	}
	resp, err := rpc.UpdateStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type DeleteStudioRequest struct {
	Id    int64
	Token string
}

func DeleteStudio(c *gin.Context) {
	receive := &DeleteStudioRequest{}
	if err := c.Bind(receive); err != nil {
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.DeleteStudioResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	req := &studio.DeleteStudioRequest{
		Id: receive.Id,
	}
	resp, err := rpc.DeleteStudio(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
