package api

import (
	"cloud-restaurant/model/request"
	"cloud-restaurant/model/response"
	"cloud-restaurant/service"
	"cloud-restaurant/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

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
