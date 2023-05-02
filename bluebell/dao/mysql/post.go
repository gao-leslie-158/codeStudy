package mysql

import (
	"bluebell/models"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

// CreatePost 将帖子信息写入数据库
func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post (" +
		"post_id,title,content,author_id,community_id)" +
		"values(?, ?, ?, ?, ?)"
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostDetailByID 通过postID 查询数据库中的单个帖子数据信息
func GetPostDetailByID(postID int64) (postDetail *models.Post, err error) {
	sqlStr := "select" +
		" post_id, community_id,author_id,title,content" +
		" from post" +
		" where post_id = ?"
	postDetail = new(models.Post)
	if err = db.Get(postDetail, sqlStr, postID); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}

// GetPostList 返回数据库中的所有帖子信息
func GetPostList(pageNum int64, pageSize int64) (postList []*models.Post, err error) {
	sqlStr := `select
				post_id,title,content,community_id,author_id,create_time
				from post
				order by create_time
				desc 
				limit ?,?`
	postList = make([]*models.Post, 0, pageSize)
	err = db.Select(&postList, sqlStr, (pageNum-1)*pageSize, pageSize)
	return
}

// GetPostListByIDlist 根据给定的post_id列表查询帖子信息
func GetPostListByIDs(ids []string) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, len(ids))
	sqlStr := `select post_id,title,content,community_id,create_time,author_id
				 from post
				where post_id in (?)
				order by find_in_set(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	//query为查询语句，args为参数列表
	if err != nil {
		return nil, err
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&posts, query, args...)
	return
}
