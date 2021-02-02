package initialize

import (
	"cloud-restaurant/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()

	group := Router.Group("")
	{
		router.InitMemberRouter(group)
	}
	return Router
}
