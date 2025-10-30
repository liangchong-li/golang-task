package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"

	"task4/config"
	"task4/internal/model"
	"task4/internal/util"
)

type PostController struct {
}

type CreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type UpdateRequest struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Create 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
func (p *PostController) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.BadRequest(c, err.Error())
		return
	}

	value, exists := c.Get("user_id")
	if !exists {
		util.Unauthorized(c, "用户未认证")
		return
	}

	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  value.(uint),
	}

	tx := config.DB.Create(&post)
	if tx.Error != nil {
		util.InternalServerError(c, "新增文章失败")
		return
	}
	util.Success(c, "创建文章成功")
}

// List 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
func (p *PostController) List(c *gin.Context) {

	var posts []model.Post

	tx := config.DB.Find(&posts)
	if tx.Error != nil {
		util.NotFound(c, "查询文章列表失败")
		return
	}

	util.Success(c, posts)
}

func (p *PostController) Info(c *gin.Context) {
	id := c.Param("id")

	var post model.Post
	tx := config.DB.Find(&post, id)
	if tx.Error != nil {
		util.NotFound(c, "查询文章详情失败")
		return
	}

	util.Success(c, post)
}

// Update 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func (p *PostController) Update(c *gin.Context) {
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.BadRequest(c, err.Error())
		return
	}

	// 操作之前，先查询。检查文章的作者是否为自己
	if !checkPostPermission(c, req.ID) {
		return
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
		util.InternalServerError(c, "更新文章失败")
		return
	}
	util.Success(c, "更新文章成功!")
}

// Delete 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func (p *PostController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.BadRequest(c, "无效的文章ID")
		return
	}

	// 操作之前，先查询。检查文章的作者是否为自己
	if !checkPostPermission(c, uint(id)) {
		return
	}

	var post model.Post
	tx := config.DB.Delete(&post, id)
	if tx.Error != nil {
		util.InternalServerError(c, "删除文章失败")
		return
	}
	util.Success(c, "删除文章成功!")
}

func checkPostPermission(c *gin.Context, postID uint) bool {
	var post model.Post
	tx := config.DB.First(&post, postID)
	if tx.Error != nil {
		util.NotFound(c, "未找到文章信息")
		return false
	}

	userID, exists := c.Get("user_id")
	if !exists {
		util.Unauthorized(c, "用户未认证")
		return false
	}

	if post.UserID != userID.(uint) {
		util.Unauthorized(c, "您不是本文章的作者,无法操作!")
		return false
	}

	return true
}
