package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
//Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	gorm.Model
	Name  string
	Age   uint
	Posts []Post
}

// 文章
type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	Comments []Comment
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

	// 先创建被引用的表。Comment reference PostID, Post reference UserID
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal(err)
	}
}
