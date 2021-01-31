package api

import (
	"cloud-restaurant/model/response"
	"cloud-restaurant/service"
	"github.com/gin-gonic/gin"
)

func SendMessage(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		response.Failed("参数解析失败", context)
		return
	}

	code := service.SendMsg(phone)
	if code == 1 {
		response.Success("验证码发送成功", context)
		return
	} else if code == 2 {
		response.Success("验证码未过期，请勿重复发送", context)
		return
	}
	response.Failed("验证码发送失败", context)
}
