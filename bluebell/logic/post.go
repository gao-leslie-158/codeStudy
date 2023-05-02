package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	// 1、生成postID
	p.ID = snowflake.GenID()
	// 2、保存到数据库
	if err = mysql.CreatePost(p); err != nil {
		return
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
}

// GetPostDetail 查询并组合接口想要的数据
func GetPostDetail(postID int64) (postDetail *models.ApiPostDetail, err error) {
	postDetail = new(models.ApiPostDetail)
	// 查询帖子数据
	post, err := mysql.GetPostDetailByID(postID)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(postID) failed",
			zap.Int64("post_id", postID),
			zap.Error(err))
		return
	}
	// 根据AuthorID查询作者信息
	user, err := mysql.GetUserDetalByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserDetalByID(post.AuthorID) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}
	// 根据CommunityID查询社区信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	// 组合想用的数据
	postDetail = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// GetPostList 获取帖子列表
func GetPostList(pageNum int64, pageSize int64) (postList []*models.ApiPostDetail, err error) {
	//以列表返回数据库中的所有帖子
	posts, err := mysql.GetPostList(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	postList = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 根据AuthorID查询作者信息
		user, err := mysql.GetUserDetalByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserDetalByID(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据CommunityID查询社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		// 组合想用的数据
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		postList = append(postList, postDetail)
	}
	return
}

// GetPostList2 按用户请求参数获取帖子列表
func GetPostList2(p *models.ParamPostlist) (postList []*models.ApiPostDetail, err error) {
	// 2、去redis查询 post_id 列表
	ids, err := redis.GetPostInOrder(p)
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostInOrder(p) success but len(ids) == 0")
		return
	}
	// 3、根据post_ids去mysql数据库查询帖子详情信息
	// 返回的数据按照我给定的顺序
	posts, err := mysql.GetPostListByIDs(ids)
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
	if err != nil {
		return
	}
	// 提前查询好帖子的投票数
	votePositive, voteNegtive, err := redis.GetPostVoteData(ids)
	postList = make([]*models.ApiPostDetail, 0, len(posts))
	// 将帖子的作者和社区信息查询出
	for idx, post := range posts {
		// 根据AuthorID查询作者信息
		user, err := mysql.GetUserDetalByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserDetalByID(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据CommunityID查询社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		// 组合想用的数据
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VotePosNum:      votePositive[idx],
			VoteNegNum:      voteNegtive[idx],
			Post:            post,
			CommunityDetail: community,
		}
		postList = append(postList, postDetail)
	}
	return
}

// GetCommunityPostList 按社区community分类得到帖子
func GetCommunityPostList(p *models.ParamPostlist) (postList []*models.ApiPostDetail, err error) {
	// 2、去redis查询 post_id 列表
	ids, err := redis.GetCommunityPost(p)
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostInOrder(p) success but len(ids) == 0")
		return
	}
	// 3、根据post_ids去mysql数据库查询帖子详情信息
	// 返回的数据按照我给定的顺序
	posts, err := mysql.GetPostListByIDs(ids)
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
	if err != nil {
		return
	}
	// 提前查询好帖子的投票数
	votePositive, voteNegtive, err := redis.GetPostVoteData(ids)
	postList = make([]*models.ApiPostDetail, 0, len(posts))
	// 将帖子的作者和社区信息查询出
	for idx, post := range posts {
		// 根据AuthorID查询作者信息
		user, err := mysql.GetUserDetalByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserDetalByID(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据CommunityID查询社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		// 组合想用的数据
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VotePosNum:      votePositive[idx],
			VoteNegNum:      voteNegtive[idx],
			Post:            post,
			CommunityDetail: community,
		}
		postList = append(postList, postDetail)
	}
	return
}

// 将GetPostList2与GetCommunityPostList两个接口合二为一
func GetPostListCompose(p *models.ParamPostlist) (postList []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		// 只按时间或者分数排序
		return GetPostList2(p)
	}
	// 同时按社区分类
	return GetCommunityPostList(p)
}
