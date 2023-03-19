package api

import (
	"TTMS/internal/web/rpc"
	"TTMS/kitex_gen/studio"
	"TTMS/pkg/jwt"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddStudio(c *gin.Context) {
	req := &studio.AddStudioRequest{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusOK, "bind error")
		return
	}
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		c.JSON(http.StatusOK, studio.AddStudioResponse{BaseResp: &studio.BaseResp{StatusCode: 1, StatusMessage: err.Error()}})
		return
	}

	resp, err := rpc.AddStudio(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, resp)
}
