package core

import (
	"cloud-restaurant/global"
	"cloud-restaurant/initialize"
	"cloud-restaurant/utils"
	"github.com/sirupsen/logrus"
	"os"
)

func Start() {
	//读取配置
	config, err := utils.ParseConfig("./config/config.yaml")
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
	}
	//连接mysql数据库
	global.DB = initialize.Gorm()
	//连接redis
	initialize.InitRedis()
	//初始化数据表
	initialize.InitMysqlTables(global.DB)
	//初始化路由
	engine := initialize.Routers()
	//启动服务
	engine.Run(config.Host + ":" + config.Port)

}
