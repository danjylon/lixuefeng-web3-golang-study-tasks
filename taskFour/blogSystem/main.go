package main

import (
	. "blogSystem/config"
	. "blogSystem/models"
	. "blogSystem/routers"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

/*
*
作业要求概述
本次作业要求你使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。
项目初始化
创建一个新的 Go 项目，使用 go mod init 初始化项目依赖管理。
安装必要的库，如 Gin 框架、GORM 以及数据库驱动（如 MySQL 或 SQLite）。
数据库设计与模型定义
设计数据库表结构，至少包含以下几个表：
users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
使用 GORM 定义对应的 Go 模型结构体。

用户认证与授权
实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。 /web/user.go
使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。 使用redis，/web/user.go->Login, /middlewares/auth.go->TokenCheck, /utils/jwt.go
文章管理功能
实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
实现文章的更新功能，只有文章的作者才能更新自己的文章。
实现文章的删除功能，只有文章的作者才能删除自己的文章。
评论功能
实现评论的创建功能，已认证的用户可以对文章发表评论。
实现评论的读取功能，支持获取某篇文章的所有评论列表。
错误处理与日志记录
对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。
*/
func main() {
	// 初始化配置和数据库
	if err := InitDatabase(); err != nil {
		panic(fmt.Sprintf("初始化失败: %v", err))
	}
	InitTables()
	engine := gin.Default()
	//设置session，secret是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	//设置session中间件，session名称为mysession，cookie名称也是mysession，通过engine引用该中间件
	handlerFunc := sessions.Sessions("mysession", store)
	// engine.Use(f)，只要是f(context *gin.Context)，参数为(context *gin.Context)的函数，既可以传入engine.Use方法，先Use哪个，哪个中间件就先执行
	engine.Use(handlerFunc)
	// 配置html页面
	//engine.LoadHTMLGlob("./templates/**/*")
	// 配置静态文件，参数1：url的访问路径，参数2：文件存放目录
	engine.Static("/static", "./static")
	// 初始化路由
	InitRouters(engine)
	port := GetServerPort()
	log.Println("port:", port)
	err := engine.Run(port)
	if err != nil {
		return
	}
}
