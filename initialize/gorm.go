package initialize

import (
	"cloud-restaurant/model"
	"cloud-restaurant/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	config := utils.GetConfig().Mysql
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Path + ")/" + config.Database + "?" + config.Config
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
	}
	logrus.Info("数据库连接成功")
	return db

}
func InitMysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Sms{},
		model.Member{},
	)
	if err != nil {
		logrus.Error("数据库表初始化失败")
		os.Exit(0)
	}
	logrus.Info("数据库表初始化成功")
}
