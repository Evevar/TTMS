package api

import (
	"TTMS/internal/web/rpc"
	user "TTMS/kitex_gen/user"
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
	}
	fmt.Println(" CreateUser ", req)
	resp, err := rpc.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}
	c.JSON(http.StatusOK, resp)
}
func UserLogin(c *gin.Context) {

}
