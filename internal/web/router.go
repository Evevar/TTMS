package main

import (
	"TTMS/internal/web/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	baseGroup := r.Group("/TTMS")

	baseGroup.POST("/user/CreateUser/", api.CreateUser)
	return r
}
