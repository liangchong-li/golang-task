# Golang Blog API
一个使用Go语言、Gin框架和GORM库开发的个人博客系统后端API。

## 功能特性
* 用户注册和登录
* JWT认证和授权
* 文章的CRUD操作
* 评论功能
* 统一的错误处理
* 日志记录
* 数据库自动迁移

## 技术栈
* Go 1.25.3
* Gin - Web框架
* GORM - ORM库
* mysql - 数据库
* JWT - 身份认证
* Logrus - 日志库
* bcrypt - 密码加密

## 项目结构
```
task4
├── cmd/
│   └── main.go                     # 程序入口
├── config/
│   └── dataBase.go                 # 数据库配置
├── internal/
│   ├── controllers/
│   │   ├── userController.go       # 用户控制器
│   │   ├── postController.go       # 文章控制器
│   │   └── commentController.go    # 评论控制器
│   ├── middleware/
│   │   ├── auth.go                 # JWT认证中间件
│   │   └── log.go                  # 日志中间件
│   ├── model/
│   │   ├── user.go                 # 用户模型
│   │   ├── post.go                 # 文章模型
│   │   └── comment.go              # 评论模型
│   ├── route/
│   │   └── route.go                # 路由配置
│   ├── util/
│   │   ├── jwt.go                  # JWT工具
│   │   └── response.go             # 响应工具
├── go.mod
├── go.sum
└── README.md
```

## 数据库设计
### Users表
* id (主键)
* username (用户名，唯一)
* email (邮箱，唯一)
* password (加密密码)
* created_at, updated_at, deleted_at
### Posts表
* id (主键)
* title (标题)
* content (内容)
* user_id (外键，关联users表)
* created_at, updated_at, deleted_at
### Comments表
* id (主键)
* content (内容)
* user_id (外键，关联users表)
* post_id (外键，关联posts表)
* created_at, updated_at, deleted_at

## API接口
### 用户接口
POST /api/v1/user/register - 用户注册
POST /api/v1/user/login - 用户登录
### 文章接口
GET /api/v1/post/list - 获取文章列表 (公开)
GET /api/v1/post/:id - 获取文章详情 (公开)
POST /api/v1/post/create - 创建文章 (需要认证)
PUT /api/v1/posts/update - 更新文章 (需要认证，仅作者)
DELETE /api/v1/posts/:id - 删除文章 (需要认证，仅作者)
### 评论接口
GET /api/v1/:post_id/comment - 获取文章评论 (公开)
POST /api/v1//:post_id/comment/list - 创建评论 (需要认证)