package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/ticket"
	"TTMS/pkg/jwt"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UpdateTicketRequest struct {
	ScheduleId int64
	SeatRow    int32
	SeatCol    int32
	Price      int32
	Token      string
}

func UpdateTicket(c *gin.Context) {
	receive := &UpdateTicketRequest{}
	if err := c.Bind(receive); err != nil {
		log.Println("err = ", err, " receive = ", receive)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, ticket.UpdateTicketResponse{BaseResp: &ticket.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	req := &ticket.UpdateTicketRequest{
		ScheduleId: receive.ScheduleId,
		SeatRow:    receive.SeatRow,
		SeatCol:    receive.SeatCol,
		Price:      receive.Price,
	}
	resp, err := rpc.UpdateTicket(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type GetAllTicketRequest struct {
	ScheduleId int64
	Token      string
}

func GetAllTicket(c *gin.Context) {
	receive := &GetAllTicketRequest{}
	if err := c.Bind(receive); err != nil {
		log.Println("err = ", err, " receive = ", receive)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, ticket.GetAllTicketResponse{BaseResp: &ticket.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	req := &ticket.GetAllTicketRequest{
		ScheduleId: receive.ScheduleId,
	}
	resp, err := rpc.GetAllTicket(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type BuyTicketRequest struct {
	ScheduleId int64
	SeatRow    int32
	SeatCol    int32
	Token      string
}

func BuyTicket(c *gin.Context) {
	receive := &BuyTicketRequest{}
	if err := c.Bind(receive); err != nil {
		log.Println("err = ", err, " receive = ", receive)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	claim, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, ticket.BuyTicketResponse{BaseResp: &ticket.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	req := &ticket.BuyTicketRequest{
		ScheduleId: receive.ScheduleId,
		SeatRow:    receive.SeatRow,
		SeatCol:    receive.SeatCol,
		UserId:     claim.ID,
	}
	resp, err := rpc.BuyTicket(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

type ReturnTicketRequest struct {
	ScheduleId int64
	SeatRow    int32
	SeatCol    int32
	Token      string
}

func ReturnTicket(c *gin.Context) {
	receive := &ReturnTicketRequest{}
	if err := c.Bind(receive); err != nil {
		log.Println("err = ", err, " receive = ", receive)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	claim, err := jwt.ParseToken(receive.Token)
	if err != nil {
		c.JSON(http.StatusOK, ticket.ReturnTicketResponse{BaseResp: &ticket.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}
	req := &ticket.ReturnTicketRequest{
		UserId:     claim.ID,
		ScheduleId: receive.ScheduleId,
		SeatRow:    receive.SeatRow,
		SeatCol:    receive.SeatCol,
	}
	resp, err := rpc.ReturnTicket(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
