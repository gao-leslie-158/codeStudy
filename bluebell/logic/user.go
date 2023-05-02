package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码

// SignUp 处理注册业务逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 1、判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2、生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3、保存进数据库
	return mysql.InsertUser(user)
}

// Login 处理登录业务逻辑
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// user传递的是一个指针，能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT的token
	atoken, _, err := jwt.GenToken(user.UserID)
	if err != nil {
		return
	}
	user.Token = atoken
	return
}
