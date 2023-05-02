package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHander 创建帖子接口
// @Summary 用户创建帖子接口
// @Description 用户必须登录才能创建帖子
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token令牌"
// @Param p body models.Post true "创建帖子参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post [post]
func CreatePostHander(c *gin.Context) {

	// 1、获取参数及参数校验
	// c.ShouldBindJSON()  //validator-->binding tag
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("create post failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 取到当前法请求的用户的 ID
	userID, err := getCurUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2、创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, CodeSuccess, "创建帖子成功！")
}

// GetPostDetailHander 获取帖子详情接口
// @Summary 根据post_id获取帖子详情的接口
// @Description 用户必须登录才能获取帖子详情
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token令牌"
// @Param object query int64 false "帖子id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post/:id [get]
func GetPostDetailHander(c *gin.Context) {
	// 1、获取帖子id --> post_id
	idStr := c.Param("id") // 获取URL参数
	postID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2、根据 id 获取帖子数据
	postDetail, err := logic.GetPostDetail(postID)
	if err != nil {
		zap.L().Error("logic.GetPostDetail(postID)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, postDetail)
}

// GetPostListHander 分页获取帖子列表的接口
// @Summary 根据page，size获取帖子详情的接口
// @Description 用户必须登录后根据page和size请求参数分页获取帖子详情列表
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token令牌"
// @Param object query string true "帖子起始页page和页数大小size,有默认值"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts [get]
func GetPostListHander(c *gin.Context) {
	// 获取页面参数信息
	pageNum, pageSize := getPageInfo(c)

	// 获取数据
	postList, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, CodeSuccess, postList)
}

// GetPostListHander2 根据前端传来的参数：可根据时间、分数、社区分类获取帖子详情
// @Summary 根据时间、帖子投票数或社区分类获取帖子详情的接口
// @Description 用户必须登录后根据需求获取帖子详情列表
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Token令牌"
// @Param object query models.ParamPostlist  false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHander2(c *gin.Context) {
	// 1、获取请求的query string参数
	// 2、去redis查询id列表
	// 3、根据post_id无数据库查询帖子详细信息
	//zap.L().Debug("GetPostListHander2(c *gin.Context)")
	// GET请求参数/api/v1/post-sort?page=1&size=2&order=time , 这是query string
	// 初始化结构体时，设置默认参数
	p := &models.ParamPostlist{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	// 1、获取参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("c.ShouldBindQuery(p) with invalid param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取数据
	//postList, err := logic.GetPostList2(p)
	postList, err := logic.GetPostListCompose(p) // 将按community和排序两个接口合而为一
	if err != nil {
		zap.L().Error("logic.GetPostList2() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, postList)
	return
}

// GetCommunityPostHander 根据社区查询帖子
func GetCommunityPostHander(c *gin.Context) {
	// GET请求参数/api/v1/post-sort?page=1&size=2&order=time , 这是query string
	// 初始化结构体时，设置默认参数
	p := &models.ParamPostlist{
		Page:        1,
		Size:        10,
		Order:       "time",
		CommunityID: 0,
	}
	// 1、获取参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("c.ShouldBindQuery(p) with invalid param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取数据
	postList, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess, postList)
	return
}
