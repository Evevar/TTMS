package main

import (
	"TTMS/internal/web/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	baseGroup := r.Group("/ttms")

	baseGroup.POST("/user/create/", api.CreateUser)
	baseGroup.POST("/user/login/", api.UserLogin)
	baseGroup.GET("/user/all/", api.GetAllUser)
	return r
}
