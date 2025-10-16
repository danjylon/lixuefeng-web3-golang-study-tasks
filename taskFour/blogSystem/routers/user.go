package routers

import (
	"blogSystem/web"
	"github.com/gin-gonic/gin"
)

func InitUser(engine *gin.Engine) {
	//http://localhost:8888/user/v1
	//userRouter := engine.Group("/user", auth.TokenCheck, auth.AuthCheck)//只对/user路由鉴权
	userRouter := engine.Group("/user")
	//userRouter.GET("/login", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "login.html", nil)
	//})
	userRouter.POST("/login", web.Login)
	userRouter.POST("/register", web.Register)
	userRouter.GET("/:username", web.GetUserInfo)
}
