package main

import (
	"TTMS/internal/web/api"
	"TTMS/pkg/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	baseGroup := r.Group("/ttms")

	baseGroup.POST("/user/create/", api.CreateUser)
	baseGroup.POST("/user/login/", api.UserLogin)
	baseGroup.POST("/user/verify/", api.GetVerification)
	baseGroup.POST("/user/forget/", api.ForgetPassword)
	//以上内容不需要Token鉴权
	baseGroup.GET("/user/all/", AuthMiddleware(), api.GetAllUser)
	baseGroup.POST("/user/change", AuthMiddleware(), api.ChangeUserPassword)
	baseGroup.POST("/user/delete/", AuthMiddleware(), api.DeleteUser)
	baseGroup.GET("/user/info/", AuthMiddleware(), api.GetUserInfo)
	baseGroup.POST("/user/bind/", AuthMiddleware(), api.BindEmail)

	studioGroup := baseGroup.Group("/studio", AuthMiddleware())
	studioGroup.POST("/add/", api.AddStudio)
	studioGroup.GET("/all/", api.GetAllStudio)
	studioGroup.GET("/info/", api.GetStudio)
	studioGroup.POST("/update/", api.UpdateStudio)
	studioGroup.POST("/delete/", api.DeleteStudio)

	seatGroup := baseGroup.Group("/seat", AuthMiddleware())
	seatGroup.POST("/add/", api.AddSeat)
	seatGroup.POST("/update/", api.UpdateSeat)
	seatGroup.POST("/delete/", api.DeleteSeat)
	seatGroup.GET("/all/", api.GetAllSeat)

	playGroup := baseGroup.Group("/play", AuthMiddleware())
	playGroup.POST("/add/", api.AddPlay)
	playGroup.POST("/update/", api.UpdatePlay)
	playGroup.POST("/delete/", api.DeletePlay)
	playGroup.GET("/all/", api.GetAllPlay)

	scheduleGroup := baseGroup.Group("/schedule", AuthMiddleware())
	scheduleGroup.POST("/add/", api.AddSchedule)
	scheduleGroup.POST("/update/", api.UpdateSchedule)
	scheduleGroup.POST("/delete/", api.DeleteSchedule)
	scheduleGroup.GET("/all/", api.GetAllSchedule)

	ticketGroup := baseGroup.Group("/ticket", AuthMiddleware())
	ticketGroup.POST("/update/", api.UpdateTicket)
	ticketGroup.GET("/all/", api.GetAllTicket)
	ticketGroup.POST("/buy/", api.BuyTicket)
	ticketGroup.POST("/return/", api.ReturnTicket)

	orderGroup := baseGroup.Group("/order", AuthMiddleware())
	orderGroup.POST("/commit/", api.CommitOrder)
	orderGroup.GET("/all/", api.GetAllOrder)
	orderGroup.GET("/analysis/", api.GetOrderAnalysis)
	return r
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从HTTP请求头中读取token
		tokenString := c.GetHeader("Authorization")
		// 确保token非空
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// 解析token
		claim, err := jwt.ParseToken(tokenString)

		// 处理token解析错误
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// 将解析后的token存储到gin的上下文中，以便后续使用
		c.Set("ID", claim.ID)
		c.Set("UserType", claim.UserType)
		c.Set("Token", tokenString)
		// 继续处理请求
		c.Next()
	}
}
