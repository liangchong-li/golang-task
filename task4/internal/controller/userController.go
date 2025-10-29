package controller

import (
	"log"

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

// 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
// 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
// Register 用户注册，TODO : 密码加密
func (u *UserController) Register(c *gin.Context) {
	// 绑定参数
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.BadRequest(c, err.Error())
		log.Fatal("参数异常!", err)
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	tx := config.DB.Create(&user)
	if tx.Error != nil {
		c.JSON(500, gin.H{
			"message": "插入失败!",
		})
		log.Fatal("插入失败!", tx.Error)
	}
	util.Success(c, "注册成功")
}

// TODO：认证、授权、JWT
func (u *UserController) Login(c *gin.Context) {
	// 根据用户名查询
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.BadRequest(c, err.Error())
		// TODO：该函数执行后，程序会结束
		log.Fatal("参数异常!", err)
	}

	var user model.User
	tx := config.DB.Where("username = ?", req.Username).First(&user)
	if tx.Error != nil {
		c.JSON(500, gin.H{
			"message": "查询失败!",
		})
		log.Fatal("查询失败!", tx.Error)
	}

	util.Success(c, user)
	//engine := gin.Default()
	/*engine.GET("/ping", func() {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Println("启动成功")
	err := engine.Run()
	if err != nil {
		log.Fatal("启动失败", err)
		return
	}*/
}
