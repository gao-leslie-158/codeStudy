package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 定义注册参数结构体
type ParamSignUp struct {
	// binding是gin框架用作参数校验的tag
	Username   string `json:"username" binding:"required"` // required：需要的
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password" ` // eqfeild判断一个字段是不是等于另一个字段
}

// ParamLogin 定义登录参数结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"` // required：需要的
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据参数
type ParamVoteData struct {
	// UserID 能获取到
	PostID string `json:"post_id" binding:"required"`
	Choice int8   `json:"choice,string" binding:"oneof=1 0 -1"` //赞成(1),反对(-1),取消票(0)
	// required 会过滤掉0值
}

// ParamPostlist 获取帖子列表query string参数
type ParamPostlist struct {
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
	CommunityID int64  `json:"community_id" form:"community_id"` //可以为空
}
