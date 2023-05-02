package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code":1000, 	// 程序中的错误码
	"msg": xx,		// 提示信息
	"data": {},		// 数据
}
*/

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, code ResCode, data interface{}) {
	//var code ResCode = CodeSuccess
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	})
}

// ResponseErrorWithMsg 自定义错误
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
