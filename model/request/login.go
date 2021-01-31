package request

//用户名，密码和验证码登录
type Login struct {
	Name     string `json:"name"` //用户名
	Password string `json:"pwd"`  //密码

}
