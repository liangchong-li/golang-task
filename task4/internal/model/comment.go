package model

import "gorm.io/gorm"

/*
comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
*/

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	PostID  uint

	User User
	Post Post
	//Post Post `gorm:"foreignKey:PostID"`
}
