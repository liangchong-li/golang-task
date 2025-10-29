package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"task4/config"
	"task4/internal/model"
)

// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
type PostController struct {
}

type PostRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userID"`
}

func (p *PostController) Create(c *gin.Context) {
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "参数异常!",
		})
		log.Fatal("参数异常!", err)
	}

	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	tx := config.DB.Create(&post)
	if tx.Error != nil {
		c.JSON(500, gin.H{
			"message": "插入失败!",
		})
		log.Fatal("插入失败!", tx.Error)
	}
	c.JSON(200, gin.H{
		"message": "创建文章成功!",
	})
}

func (p *PostController) List(c *gin.Context) {

	var posts []model.Post

	tx := config.DB.Find(&posts)
	if tx.Error != nil {
		c.JSON(500, gin.H{
			"message": "查询失败!",
		})
		log.Fatal("查询失败!", tx.Error)
	}
	c.JSON(200, posts)
}

func (p *PostController) Info(c *gin.Context) {
	id := c.Param("id")

	var post model.Post
	tx := config.DB.Find(&post, id)
	if tx.Error != nil {
		c.JSON(500, gin.H{
			"message": "查询失败!",
		})
		log.Fatal("查询失败!", tx.Error)
	}
	c.JSON(200, post)
}

func (p *PostController) Update(c *gin.Context) {
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "参数异常!",
		})
		log.Fatal("参数异常!", err)
	}

	post := model.Post{
		Model: gorm.Model{
			ID: req.ID,
		},
		Title:   req.Title,
		Content: req.Content,
	}

	tx := config.DB.Model(&post).Updates(post)
	if tx.Error != nil {
		c.JSON(500, gin.H{
			"message": "更新失败!",
		})
		log.Fatal("更新失败!", tx.Error)
	}
	c.JSON(200, gin.H{
		"message": "更新文章成功!",
	})
}

func (p *PostController) Delete(c *gin.Context) {
	id := c.Param("id")

	var post model.Post
	tx := config.DB.Delete(&post, id)
	if tx.Error != nil {
		c.JSON(500, gin.H{
			"message": "删除失败!",
		})
		log.Fatal("删除失败!", tx.Error)
	}
	c.JSON(200, post)
}
