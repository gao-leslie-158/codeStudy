package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

// 基于用户投票的相关算法：阮一峰博客

// 本项目使用简化版的投票分数
// 投一票加432分，一天86400/200 --> 200张赞成票给帖子续一天 《redis实战》

/* 投票的几种情况
choice =1 时，两种情况
	1、之前没有投票，现在投赞成票--> 更新分数和投票记录 差值绝对值：1
	2、之前投反对票，现在投赞成票--> 更新分数和投票记录	差值绝对值：2

choice = 0时，两种情况
	1、之前投赞成票，现在取消投票--> 更新分数和投票记录 差值绝对值：1
	2、之前投反对票，现在取消投票--> 更新分数和投票记录 差值绝对值：1

choice = -1时，两种情况
	1、之前没有投票，现在投反对票--> 更新分数和投票记录 差值绝对值：1
	2、之前投赞成票，现在投反对票--> 更新分数和投票记录 差值绝对值：2

投票限制：
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期不允许投票
	1、到期之后，将redis中保存的投票情况保存到mysql中
	2、到期之后，删除 KeyPostVotedZSetPF
*/

// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("logic.VoteForPost:",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("choice", p.Choice))
	// 格式转换
	userIDStr := strconv.FormatInt(userID, 10) //int64 --> string
	//postIDStr := strconv.FormatInt(p.PostID, 10)
	return redis.VoteForPost(userIDStr, p.PostID, float64(p.Choice))
}
