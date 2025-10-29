package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"task4/internal/controller"
)

func SetupRoutes() {
	engine := gin.New()

	// 创建控制器实例
	userController := &controller.UserController{}
	postController := &controller.PostController{}
	commentController := &controller.CommentController{}

	// 路由分组
	v1 := engine.Group("v1/api")
	user := v1.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	post := v1.Group("/post")
	{
		post.POST("/create", postController.Create)
		post.PUT("/update", postController.Update)
		post.GET("/list", postController.List)
		post.GET("/:id", postController.Info)
		post.DELETE("/:id", postController.Delete)
	}

	comment := v1.Group("/comment")
	{
		comment.POST("/commit", commentController.Commit)
		comment.GET("/:id", commentController.Read)
	}

	err := engine.Run()
	if err != nil {
		log.Fatal("路由启动异常", err)
		return
	}
}
