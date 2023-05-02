package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID int64, communityID int64) (err error) {
	//引入事务操作，创建时间与初试分数赋值是一个事务，原子性操作
	// 用pipeline
	pipeline := client.Pipeline()

	// 帖子创建时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子初始分数：为帖子创建时间
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 创建为每个社区community创建帖子分类
	communityIDstr := strconv.FormatInt(communityID, 10)
	ckey := getRedisKey(KeyCommunitySetPF + communityIDstr)
	pipeline.SAdd(ckey, postID)

	_, err = pipeline.Exec()
	return err
}

// getIDsFromKey 根据缓存的key 查询ids
func getIDsFromKey(key string, page, size int64) ([]string, error) {
	// 确认，查询索引的起始点
	start := (page - 1) * size
	end := start + size
	// 3、zrevrange 按分数从大到小查询
	return client.ZRevRange(key, start, end).Result()
}

// GetPostInOrder 根据post_id按指定顺序获取帖子分数
func GetPostInOrder(p *models.ParamPostlist) ([]string, error) {
	// 从redis获取id
	// 根据用户请求中的参数order确定要查询的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 确认，查询索引的起始点
	return getIDsFromKey(key, p.Page, p.Size)
}

// GetCommunityPost 按社区获取帖子id列表 -->ids
func GetCommunityPost(p *models.ParamPostlist) ([]string, error) {
	// 使用zinterstore 把分区的帖子 set 和 按时间或分数的 zset 取交集
	// 生成新的zset ，针对新的zset 按之前的逻辑取数据
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 社区community的key
	cKey := getRedisKey(KeyCommunitySetPF + strconv.FormatInt(p.CommunityID, 10))
	// 利用缓存 key 减少 zinterstore 执行的次数
	key := orderKey + strconv.FormatInt(p.CommunityID, 10)
	if client.Exists(key).Val() < 1 {
		// 不存在，需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "Max",
		}, cKey, orderKey) // zinterstore计算
		pipeline.Expire(key, expireTime) // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	// 存在直接根据key查询ids
	return getIDsFromKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据id获取帖子投票分数
func GetPostVoteData(ids []string) (dataP []int64, dataN []int64, err error) {
	pipelineP := client.Pipeline()
	pipelineN := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipelineP.ZCount(key, "1", "1")   // 计算投赞成票的
		pipelineN.ZCount(key, "-1", "-1") // 计算投反对票的
	}
	cmdersP, err := pipelineP.Exec()
	if err != nil {
		return nil, nil, err
	}
	cmdersN, err := pipelineN.Exec()
	if err != nil {
		return nil, nil, err
	}
	dataP = make([]int64, 0, len(cmdersP))
	dataN = make([]int64, 0, len(cmdersN))
	for _, cmderP := range cmdersP {
		v := cmderP.(*redis.IntCmd).Val()
		dataP = append(dataP, v)
	}
	for _, cmderN := range cmdersN {
		v := cmderN.(*redis.IntCmd).Val()
		dataN = append(dataN, v)
	}
	return
}
