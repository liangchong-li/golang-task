package model

import "gorm.io/gorm"

/*
存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
*/

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint

	User     User
	Comments []Comment
}
