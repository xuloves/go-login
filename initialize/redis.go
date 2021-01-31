package initialize

import (
	"cloud-restaurant/global"
	"cloud-restaurant/utils"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"os"
)

func InitRedis() {
	config := utils.GetConfig().Redis
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.DB,       // use default DB
	})
	_, err := global.Redis.Ping().Result()
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
	}
	logrus.Info("redis连接成功")

}
