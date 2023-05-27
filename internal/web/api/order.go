package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/order"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAllOrder(c *gin.Context) {
	req := &order.GetAllOrderRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.GetAllOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

func GetOrderAnalysis(c *gin.Context) {
	req := &order.GetOrderAnalysisRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.GetOrderAnalysis(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}

func CommitOrder(c *gin.Context) {
	req := &order.CommitOrderRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	resp, err := rpc.CommitOrder(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err)
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, resp)
}
