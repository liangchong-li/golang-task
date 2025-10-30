package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"task4/internal/controller"
	"task4/internal/middleware"
)

func SetupRoutes() {
	engine := gin.New()

	// 创建控制器实例
	userController := &controller.UserController{}
	postController := &controller.PostController{}
	commentController := &controller.CommentController{}

	// 路由分组
	v1 := engine.Group("v1/api")

	// 用户登录、注册。无需认证
	user := v1.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	// 需要认证的路由
	authed := v1.Group("")
	// 注册权限认证中间件
	authed.Use(middleware.Auth())
	{
		// 文章接口，需要认证的
		post := authed.Group("/post")
		{
			post.POST("/create", postController.Create)
			post.PUT("/update", postController.Update)
			post.DELETE("/:id", postController.Delete)
		}

		// 评论接口，需要认证的
		comment := authed.Group("/:post_id/comment")
		{
			comment.POST("/commit", commentController.Commit)
		}
	}

	// 不需要认证的路由
	noAuth := v1.Group("")
	post := noAuth.Group("/post")
	{
		post.GET("/list", postController.List)
		post.GET("/:id", postController.Info)
	}

	comment := noAuth.Group("/:post_id/comment")
	{
		comment.GET("/list", commentController.Read)
	}

	err := engine.Run()
	if err != nil {
		log.Fatal("路由启动异常", err)
		return
	}
}
