package dao

import (
	"cloud-restaurant/global"
	"cloud-restaurant/model"
)

func ValidateMsgCode(phone string, code string) (sms *model.Sms) {
	sms = new(model.Sms)
	err := global.DB.Where("phone = ? AND  code = ?", phone, code).First(&sms).Error
	if err != nil {
		return nil
	}
	return sms
}

func QueryByPhone(phone string) (member *model.Member) {
	member = new(model.Member)
	err := global.DB.Where("phone", phone).First(&member).Error
	if err != nil {
		return nil
	}
	return member
}
func InsertMember(user *model.Member) (err error) {
	err = global.DB.Save(&user).Error
	return
}

func Login(name string, pwd string) (member *model.Member) {
	member = new(model.Member)
	err := global.DB.Where("user_name = ? AND  password = ?", name, pwd).First(&member).Error
	if err != nil {
		return nil
	}
	return
}
