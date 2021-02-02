package service

import (
	"cloud-restaurant/dao"
	"cloud-restaurant/global"
	"cloud-restaurant/model"
	"cloud-restaurant/model/request"
	"cloud-restaurant/utils"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func SendMsg(phone string) int {
	//1.产生一个验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	//redis获取看是否过期
	_, err := global.Redis.Get(phone).Result()
	if err == redis.Nil {
		//redis未存验证码
		//2.调用阿里云sdk 完成发送
		config := utils.GetConfig().Sms
		client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AccessKeyID, config.AccessKeySecret)
		if err != nil {
			return 0
		}
		request := dysmsapi.CreateSendSmsRequest()
		request.Scheme = "https"
		request.SignName = config.SignName
		request.TemplateCode = config.TemplateCode
		request.PhoneNumbers = phone
		par, err := json.Marshal(map[string]interface{}{
			"code": code,
		})
		request.TemplateParam = string(par)
		response, err := client.SendSms(request)
		fmt.Println(response)

		if err != nil {
			return 0
		}

		//3.接收返回结果，并判断发送状态
		//短信验证码发送成功
		if response.Code == "OK" {
			//存入redis
			global.Redis.Set(phone, code, 3*time.Minute)
			//将验证码保存到数据库中
			smsCode := model.Sms{Phone: phone, Code: code, BizId: response.BizId, CreateTime: time.Now().Unix()}
			err := dao.InsertSms(&smsCode)
			if err != nil {
				return 0
			}
			return 1
		}

	} else if err != nil {
		logrus.Error(err.Error())
		return 0
	}
	//已经有验证码
	return 2
}

func MsgLogin(msgLogin request.MsgLogin) *model.Member {

	//1.获取到手机号和验证码

	//2.验证手机号+验证码是否正确
	sms := dao.ValidateMsgCode(msgLogin.Phone, msgLogin.Code)
	if sms == nil {
		return nil
	}

	//3、根据手机号member表中查询记录
	member := dao.QueryByPhone(msgLogin.Phone)
	if member != nil {
		return member
	}
	//4.新创建一个member记录，并保存
	user := model.Member{}
	user.UserName = msgLogin.Phone
	user.Mobile = msgLogin.Phone
	user.RegisterTime = time.Now().Unix()
	dao.InsertMember(&user)
	return &user
}
