package router

import (
	"cloud-restaurant/api"
	"github.com/gin-gonic/gin"
)

func InitSmsRouter(Router *gin.RouterGroup) {
	smsRouter := Router.Group("api")
	{
		smsRouter.GET("sendMsg", api.SendMessage)
		smsRouter.OPTIONS("msgLogin", api.MsgLogin)
		smsRouter.POST("login", api.Login)
	}
}
