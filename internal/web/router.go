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
	baseGroup.POST("/user/change", api.ChangeUserPassword)
	baseGroup.POST("/user/delete/", api.DeleteUser)
	baseGroup.GET("/user/info/", api.GetUserInfo)

	baseGroup.POST("/studio/add/", api.AddStudio)
	baseGroup.GET("/studio/all/", api.GetAllStudio)
	baseGroup.POST("/studio/update/", api.UpdateStudio)
	baseGroup.POST("/studio/delete/", api.DeleteStudio)

	baseGroup.POST("/seat/add/", api.AddSeat)
	baseGroup.POST("/seat/update/", api.UpdateSeat)
	baseGroup.POST("/seat/delete/", api.DeleteSeat)
	baseGroup.GET("/seat/all/", api.GetAllSeat)

	return r
}
