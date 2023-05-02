package models

import (
	"time"
)

// 内存对齐：相同类型的按循序放在一起
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情
type ApiPostDetail struct {
	AuthorName       string                    `json:"user_name" db:"user_name"`
	VotePosNum       int64                     `json:"vote_pos_num,string"`
	VoteNegNum       int64                     `json:"vote_neg_num,string"`
	*Post            `json:"post_details"`     // 嵌入帖子结构体
	*CommunityDetail `json:"community_detail"` // 嵌入社区结构体
}
