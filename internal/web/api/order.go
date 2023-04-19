package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/order"
	"TTMS/pkg/jwt"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GetAllOrderRequest struct {
	Token string
}

func GetAllOrder(c *gin.Context) {
	receive := &GetAllOrderRequest{}
	if err := c.Bind(receive); err != nil {
		log.Println("err = ", err, " receive = ", receive)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	claim, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, order.GetAllOrderResponse{BaseResp: &order.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	fmt.Println("receive = ", receive)
	req := &order.GetAllOrderRequest{
		UserId: claim.ID,
	}
	fmt.Println("req = ", req)
	resp, err := rpc.GetAllOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type GetOrderAnalysisRequest struct {
	PlayId int64
	Token  string
}

func GetOrderAnalysis(c *gin.Context) {
	receive := &GetOrderAnalysisRequest{}
	if err := c.Bind(receive); err != nil {
		log.Println("err = ", err, " receive = ", receive)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, order.GetOrderAnalysisResponse{BaseResp: &order.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	req := &order.GetOrderAnalysisRequest{
		PlayId: receive.PlayId,
	}
	resp, err := rpc.GetOrderAnalysis(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type CommitOrderRequest struct {
	ScheduleId int64
	SeatRow    int32
	SeatCol    int32
	Token      string
}

func CommitOrder(c *gin.Context) {
	receive := &CommitOrderRequest{}
	if err := c.Bind(receive); err != nil {
		log.Println("err = ", err, " receive = ", receive)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	claim, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, order.CommitOrderResponse{BaseResp: &order.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	req := &order.CommitOrderRequest{
		UserId:     claim.ID,
		ScheduleId: receive.ScheduleId,
		SeatRow:    receive.SeatRow,
		SeatCol:    receive.SeatCol,
	}
	resp, err := rpc.CommitOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
