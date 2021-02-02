package router

import (
	"cloud-restaurant/api"
	"github.com/gin-gonic/gin"
)

func InitMemberRouter(Router *gin.RouterGroup) {
	memberRouter := Router.Group("api")
	{
		memberRouter.GET("sendMsg", api.SendMessage)
		memberRouter.OPTIONS("msgLogin", api.MsgLogin)
		memberRouter.POST("login", api.Login)
	}
}
