package route

import (
	"github.com/gin-gonic/gin"
	"task4/internal/controller"
	"task4/internal/middleware"
)

func SetupRoutes() *gin.Engine {
	router := gin.New()

	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.ErrorHandleMiddleware())
	// 恢复任何恐慌（panic），确保服务器不会因为某个请求发生恐慌而崩溃。
	router.Use(gin.Recovery())

	// 创建控制器实例
	userController := &controller.UserController{}
	postController := &controller.PostController{}
	commentController := &controller.CommentController{}

	// 路由分组
	v1 := router.Group("v1/api")

	// 用户登录、注册。无需认证
	user := v1.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	// 需要认证的路由
	authed := v1.Group("")
	// 注册权限认证中间件
	authed.Use(middleware.AuthMiddleware())
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

	return router
}
