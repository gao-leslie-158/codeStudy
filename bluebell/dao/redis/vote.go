package redis

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/go-redis/redis"
)

// 基于用户投票的相关算法：阮一峰博客

// 本项目使用简化版的投票分数
// 投一票加432分，一天86400/200 --> 200张赞成票给帖子续一天 《redis实战》

/* 投票的几种情况
choice =1 时，两种情况
	1、之前没有投票，现在投赞成票--> 更新分数和投票记录 差值：1 +864
	2、之前投反对票，现在投赞成票--> 更新分数和投票记录	差值：2  +864

choice = 0时，两种情况
	1、之前投反对票，现在取消投票--> 更新分数和投票记录 差值：1 +864
	2、之前投赞成票，现在取消投票--> 更新分数和投票记录 差值：1 -864

choice = -1时，两种情况
	1、之前没有投票，现在投反对票--> 更新分数和投票记录 差值：1 -864
	2、之前投赞成票，现在投反对票--> 更新分数和投票记录 差值：2 -864

投票限制：
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期不允许投票
	1、到期之后，将redis中保存的投票情况保存到mysql中
	2、到期之后，删除 KeyPostVotedZSetPF
*/

const (
	oneWeekInSecond = 7 * 24 * 3600
	scorePerVote    = (24 * 3600) / 100 // 每一票加多少分，每投100票多续命一天
	expireTime      = 60 * time.Second
)

var (
	ErrorVoteTimeExpire = errors.New("投票已经结束")
	ErrorVoteRepeated   = errors.New("不允许重复投票")
)

// VoteForPost 为帖子投票
func VoteForPost(userID, postID string, choice float64) error {
	// 1、判断投票的限制
	// 去redis取发布时间
	postTime := client.ZScore(getRedisKey(KeyPostScoreZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSecond {
		return ErrorVoteTimeExpire
	}
	fmt.Println("postTime:", postTime)
	// 2、更新帖子分数
	// 先查当前用户该帖子之前的投票记录
	oc := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	var op float64
	if oc < choice {
		op = 1
	} else if oc == choice {
		return ErrorVoteRepeated
	} else {
		op = -1
	}
	diff := math.Abs(oc - choice)
	//用pipeline事务操作
	pipeline := client.Pipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 3、记录用户为该帖子投过票
	if choice == 0 {
		// 移除用户投票记录
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID).Result()
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  choice, //当前用户投票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
