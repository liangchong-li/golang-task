package controller

import (
	"github.com/gin-gonic/gin"

	"task4/config"
	"task4/internal/model"
	"task4/internal/util"
)

type UserController struct {
}

type RegisterRequest struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Register 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
func (u *UserController) Register(c *gin.Context) {
	// 绑定参数
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.BadRequest(c, err.Error())
		//log.Fatal("参数异常!", err)
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	password, err := user.HashedPassword()
	if err != nil {
		util.InternalServerError(c, "加密失败!")
		return
	}

	user.Password = password

	tx := config.DB.Create(&user)
	if tx.Error != nil {
		util.InternalServerError(c, "新增用户失败!")
		return
	}
	util.Success(c, "注册成功")
}

// Login 登录 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
func (u *UserController) Login(c *gin.Context) {
	// 根据用户名查询
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.BadRequest(c, err.Error())
		return
	}

	var user model.User
	tx := config.DB.Where("username = ?", req.Username).First(&user)
	// 查询不到记录时，也会抛出异常
	if tx.Error != nil {
		util.NotFound(c, "查询用户失败!")
		return
	}

	// 验证密码
	if err := user.CheckPassword(req.Password); err != nil {
		util.BadRequest(c, "密码错误!")
		return
	}

	// 返回token
	token, err := util.Generate(user.ID, user.Username)
	if err != nil {
		util.InternalServerError(c, "生成token失败")
		return
	}
	util.Success(c, gin.H{
		"token": token,
		"user":  user,
	})
}
