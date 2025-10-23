package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//题目3：钩子函数
//继续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

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

	// 初始化几个user
	//initUser(db)
	// 初始化几个文章评论
	//comment := Comment{Content: "有深度", PostID: 3}
	//addComment(db, &comment)
	//comment = Comment{Content: "有趣", PostID: 3}
	//addComment(db, &comment)
	//comment = Comment{Content: "精彩", PostID: 3}
	//addComment(db, &comment)

	// 单表保存
	post := Post{UserID: 1, Title: "我爱中国", Content: "我爱你，爱到了空气\n无论你在哪里我感觉到你\n不曾惹人注意，却知道你已是我生命的必须\n爱上你，超越了自己\n让那岁月弥补着美丽"}
	savePost(db, &post)

	deleteComment(db, 3)

}

func savePost(db *gorm.DB, post *Post) {
	db.Create(post)
}

func addComment(db *gorm.DB, comment *Comment) {
	db.Create(comment)
}

func deleteComment(db *gorm.DB, id uint) {
	comment := Comment{}
	comment.ID = id
	// 要使用钩子，必须往方法中，写入足够的信息
	db.First(&comment, id)
	fmt.Println("删除前查询：", comment)
	db.Delete(&comment)
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	user := User{}
	user.ID = post.UserID
	tx.First(&user, post.UserID)
	tx.Model(user).Update("PostNum", user.PostNum+1)
	return
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	post := Post{}
	post.ID = comment.PostID
	fmt.Println("删除评论钩子，comment：", comment)
	tx.First(&post, comment.PostID)
	if post.CommentNum >= 1 {
		tx.Model(post).Update("CommentNum", post.CommentNum-1)
	}

	if post.CommentNum == 1 {
		fmt.Println("执行无评论修改")
		tx.Model(post).Update("CommentState", "无评论")
	}
	return
}

func initUser(db *gorm.DB) {
	users := []User{
		{Name: "张三", Age: 17},
		{Name: "李四", Age: 18},
		{Name: "王五", Age: 19},
	}
	db.Create(users)
}
