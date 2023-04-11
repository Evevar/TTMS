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
	baseGroup.POST("/user/verify/", api.GetVerification)
	baseGroup.POST("/user/bind/", api.BindEmail)
	baseGroup.POST("/user/forget/", api.ForgetPassword)

	baseGroup.POST("/studio/add/", api.AddStudio)
	baseGroup.GET("/studio/all/", api.GetAllStudio)
	baseGroup.GET("/studio/info/", api.GetStudio)
	baseGroup.POST("/studio/update/", api.UpdateStudio)
	baseGroup.POST("/studio/delete/", api.DeleteStudio)

	baseGroup.POST("/seat/add/", api.AddSeat)
	baseGroup.POST("/seat/update/", api.UpdateSeat)
	baseGroup.POST("/seat/delete/", api.DeleteSeat)
	baseGroup.GET("/seat/all/", api.GetAllSeat)

	baseGroup.POST("/play/add/", api.AddPlay)
	baseGroup.POST("/play/update/", api.UpdatePlay)
	baseGroup.POST("/play/delete/", api.DeletePlay)
	baseGroup.GET("/play/all/", api.GetAllPlay)

	baseGroup.POST("/schedule/add/", api.AddSchedule)
	baseGroup.POST("/schedule/update/", api.UpdateSchedule)
	baseGroup.POST("/schedule/delete/", api.DeleteSchedule)
	baseGroup.GET("/schedule/all/", api.GetAllSchedule)

	baseGroup.POST("/ticket/update/", api.UpdateTicket)
	baseGroup.GET("/ticket/all/", api.GetAllTicket)
	baseGroup.POST("/ticket/buy/", api.BuyTicket)
	baseGroup.POST("/ticket/return/", api.ReturnTicket)
	return r
}
