package controller

import (
	"github.com/gin-gonic/gin"

	"task4/config"
	"task4/internal/model"
	"task4/internal/util"
)

type CommentController struct {
}

type CommentRequest struct {
	Content string `json:"content"`
	UserID  uint   `json:"userID"`
	PostID  uint   `json:"postID"`
}

// Commit 实现评论的创建功能，已认证的用户可以对文章发表评论。
func (cm *CommentController) Commit(c *gin.Context) {
	var req CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.BadRequest(c, err.Error())
		return
	}

	comment := model.Comment{
		Content: req.Content,
		UserID:  req.UserID,
		PostID:  req.PostID,
	}

	tx := config.DB.Create(&comment)
	if tx.Error != nil {
		util.InternalServerError(c, "新增文章评论失败")
		return
	}
	util.Success(c, "提交评论成功!")

}

// Read 实现评论的读取功能，支持获取某篇文章的所有评论列表。
func (cm *CommentController) Read(c *gin.Context) {
	id := c.Param("id")

	var comments []model.Comment
	tx := config.DB.Where("post_id = ?", id).Find(&comments)
	//tx := config.DB.Debug().Where("post_id = ?", id).Find(&comment)
	if tx.Error != nil {
		util.NotFound(c, "未找到文章评论")
		return
	}
	util.Success(c, comments)
}
