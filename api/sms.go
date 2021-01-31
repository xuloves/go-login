package api

import (
	"cloud-restaurant/model/request"
	"cloud-restaurant/model/response"
	"cloud-restaurant/service"
	"cloud-restaurant/utils"
	"encoding/json"
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

func MsgLogin(context *gin.Context) {
	var msgLogin request.MsgLogin
	err := utils.Decode(context.Request.Body, &msgLogin)

	if err != nil {
		response.Failed("参数解析失败", context)
		return
	}
	member := service.MsgLogin(msgLogin)
	member.Password = ""
	if member != nil {
		response.Ok(member, "登陆成功", context)
		return
	}
	response.Failed("登录失败", context)
}

func Login(context *gin.Context) {
	var login request.Login
	data, _ := context.GetRawData() // 从c.Request.Body读取请求数据
	err := json.Unmarshal(data, &login)
	if err != nil {
		response.Failed("参数解析失败", context)
		return
	}
	b := service.Login(login)
	if b {
		response.Success("登陆成功", context)
		return
	}
	response.Failed("登录失败", context)
}
