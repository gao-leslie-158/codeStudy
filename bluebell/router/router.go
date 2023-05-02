package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"bluebell/settings"
	"net/http"
	"strconv"

	swaggerFiles "github.com/swaggo/files"

	"go.uber.org/zap"

	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Setup(cfg *settings.AppConfig) *gin.Engine {
	if cfg.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	}
	r := gin.New()
	// 注册gin-swagger 接口路由
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//v1 := r.Group("/api/v1", middlewares.RateLimitMiddleware(time.Second, 4))
	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHander)
	// 登录业务逻辑
	v1.POST("/login", controller.LoginHander)

	v1.Use(middlewares.JWTAuthMiddleware()) //应用JWT认证中间件
	pprof.Register(r)                       // 注册pprof路由
	// JWT token校验
	//v1.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	// 如果是登录的用户,判断请求头中是否有 有效的JWT
	//	// 登录才能访问的地方，权限、业务体验
	//	c.String(http.StatusOK, "登录且认证成功")
	//})

	{
		v1.GET("/community", controller.CommunityHander)
		v1.GET("/community/:id", controller.CommunityDetailHander)

		v1.POST("/post", controller.CreatePostHander)       // 创建帖子
		v1.GET("/post/:id", controller.GetPostDetailHander) // 根据post_id获取帖子详情
		v1.GET("/posts", controller.GetPostListHander)      // 分页获取帖子列表，
		//v1.GET("/posts/order", controller.GetPostListHander2)         // 按时间或者分数获取帖子列表
		//v1.GET("/posts/community", controller.GetCommunityPostHander) // 按community获取帖子列表
		v1.GET("/posts2", controller.GetPostListHander2) // 可根据时间、分数、社区分类

		v1.POST("/vote", controller.PostVoteHander) // 帖子投票
	}
	// 激活路由
	err := r.Run(":" + strconv.Itoa(cfg.Port))
	if err != nil {
		zap.L().Error("r.Run() err:", zap.Error(err))
		return nil
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
