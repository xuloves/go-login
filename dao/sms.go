package dao

import (
	"cloud-restaurant/global"
	"cloud-restaurant/model"
)

//插入短信
func InsertSms(sms *model.Sms) (err error) {
	err = global.DB.Save(&sms).Error
	return
}
