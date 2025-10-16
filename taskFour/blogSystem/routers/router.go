package routers

import (
	"blogSystem/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 全局处理跨域
func AccessControlAllowOrigin(context *gin.Context) {
	origin := context.GetHeader("Origin")
	context.Header("Access-Control-Allow-Origin", origin)
	context.Next()
}
func InitRouters(engine *gin.Engine) {
	//全局拦截鉴权
	engine.Use(middlewares.TokenCheck, middlewares.AuthCheck(), middlewares.RecoveryMiddleware(), AccessControlAllowOrigin)
	//初始化想要的路由
	InitUser(engine)
	InitBlog(engine)
	InitComment(engine)
	// 路由匹配失败自定义返回结果，默认返回404 not found
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "not found",
		})
		//context.HTML(http.StatusNotFound, "404.html", nil)
	})
}
