package service

import (
	"cloud-restaurant/dao"
	"cloud-restaurant/model/request"
)

func Login(login request.Login) bool {
	member := dao.Login(login.Name, login.Password)
	if member != nil {
		return true
	}
	return false
}
