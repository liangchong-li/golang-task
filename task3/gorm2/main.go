package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。

type User struct {
	gorm.Model
	Name    string
	Age     uint
	PostNum uint
	Posts   []Post
}

// 文章
type Post struct {
	gorm.Model
	Title        string
	Content      string
	CommentNum   uint
	CommentState string
	UserID       uint
	Comments     []Comment
}

// 评论
type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}

func main() {
	dsn := "root:admin@tcp(192.168.2.2:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 同时新增文章和评论
	//post := Post{
	//	UserID:  1,
	//	Title:   "劝学",
	//	Content: "君子曰：学不可以已。青，取之于蓝，而青于蓝；冰，水为之，而寒于水。",
	//	Comments: []Comment{
	//		{Content: "666"},
	//		{Content: "优秀"},
	//		{Content: "无敌"},
	//		{Content: "一般"},
	//	},
	//}
	//savePost(db, &post)

	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// 使用 Preload预加载
	user := User{}
	db.Preload("Posts.Comments").First(&user, 1)

	fmt.Printf("用户信息: ID=%d, Name=%s, Age=%d\n", user.ID, user.Name, user.Age)
	fmt.Println("用户发布的文章:")

	for _, post := range user.Posts {
		fmt.Printf("  文章: ID=%d, Title=%s, Content=%s\n", post.ID, post.Title, post.Content)
		fmt.Println("  评论:")

		for _, comment := range post.Comments {
			fmt.Printf("    - ID=%d, Content=%s\n", comment.ID, comment.Content)
		}
		fmt.Println()
	}

	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	post := Post{}
	// db.Order("CommentNum desc").First(&post)
	// Error 1054 (42S22): Unknown column 'CommentNum' in 'order clause'

	db.Order("comment_num desc").First(&post)
	//db.Debug()
	fmt.Println("评论数量最多的文章信息: ", post)
}

func savePost(db *gorm.DB, post *Post) {
	db.Create(post)
}
