package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHander 处理注册请求参数的函数的函数
func SignUpHander(c *gin.Context) {
	// 1、获取参数和参数校验
	p := new(models.ParamSignUp)
	// 这里只能做格式校验
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 使用翻译器翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2、业务处理
	if err := logic.SignUp(p); err != nil {
		// p为请求参数
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, CodeSuccess, "signup success")
	return
}

// LoginHander 处理登录请求参数的函数
func LoginHander(c *gin.Context) {
	// 1、获取参数和参数校验
	p := new(models.ParamLogin)
	// 这里只能做格式校验
	if err := c.ShouldBindJSON(&p); err != nil {
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		zap.L().Error("Login with invalid param", zap.Error(errs))
		// 请求参数有误，直接返回响应
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 使用翻译器翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2、业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login() failed...", zap.String("username", p.Username), zap.Error(err))
		ResponseError(c, CodeInvalidPssword)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, CodeSuccess, gin.H{
		"user_id":  fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1（json能表示的最大值）,而int64: 1<<63-1，会失真
		"username": user.Username,
		"token":    user.Token,
	})
	return
}
