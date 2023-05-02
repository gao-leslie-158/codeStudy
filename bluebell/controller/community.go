package controller

import (
	"bluebell/logic"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//--------- 跟社区相关的 -----------

func CommunityHander(c *gin.Context) {
	// 查询到所有的社区(community_id,community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() fialed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易暴露服务端具体错误，而是存在日志里
		return
	}
	fmt.Println(c)
	ResponseSuccess(c, CodeSuccess, data)
}

func CommunityDetailHander(c *gin.Context) {
	// 1、获取社区id
	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() fialed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易暴露服务端具体错误，而是存在日志里
		return
	}
	ResponseSuccess(c, CodeSuccess, data)
}
