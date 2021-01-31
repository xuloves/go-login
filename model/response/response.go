package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	FAILED  = 1
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Success(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}
func Ok(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}
func Error(c *gin.Context) {
	Result(FAILED, map[string]interface{}{}, "操作失败", c)
}
func Failed(message string, c *gin.Context) {
	Result(FAILED, map[string]interface{}{}, message, c)
}
