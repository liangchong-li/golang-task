package main

import (
	"github.com/bytedance/gopkg/util/logger"
	"os"
	"os/signal"
	"syscall"
	"task4/config"
	"task4/internal/route"
)

func main() {

	// 优化：定义config.go,加载配置、读取环境变量。环境变量优先
	// 在main或config.InitDataBase中使用配置
	defaultPort := "8080"

	config.InitDataBase()
	router := route.SetupRoutes()

	// 设置Gin模式
	//	gin.SetMode(gin.ReleaseMode)

	// 启动服务器
	go func() {
		if err := router.Run(":" + defaultPort); err != nil {
			logger.Fatal("服务器启动失败: %v", err)
		}
	}()

	logger.Info("服务器启动成功，监听端口: %s", defaultPort)

	// 等待中断信号优雅关闭
	// 终端信号
	// 1. SIGINT (Interrupt Signal) 用户主动中断程序，通常由 Ctrl+C 触发
	// 2. SIGTERM (Termination Signal) 通常由系统工具（如 kill 命令）发送，请求程序正常终止
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("正在关闭服务器...")

	// 关闭数据库连接
	config.CloseDB()
	logger.Info("服务器已关闭")
}
