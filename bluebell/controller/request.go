package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "UserID"
)

var (
	ErrorUserNOtLogin = errors.New("当前用户未登录")
)

// getCurUserID 获取当前用户的ID
func getCurUserID(c *gin.Context) (UserID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNOtLogin
		return
	}
	UserID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNOtLogin
		return
	}
	return
}

// getPageInfo 获取页面参数信息
func getPageInfo(c *gin.Context) (int64, int64) {
	// 获取分页参数
	pageNumStr := c.Query("page")
	pagesizeStr := c.Query("size")

	var (
		pageNum  int64
		pageSize int64
		err      error
	)
	pageNum, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	pageSize, err = strconv.ParseInt(pagesizeStr, 10, 64)
	if err != nil {
		pageSize = 3
	}
	return pageNum, pageSize
}
