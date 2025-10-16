package routers

import (
	"blogSystem/web"
	"github.com/gin-gonic/gin"
)

func InitComment(engine *gin.Engine) {
	//http://localhost:8888/user/v1
	//userRouter := engine.Group("/user", auth.TokenCheck, auth.AuthCheck)//只对/user路由鉴权
	blogRouter := engine.Group("/comment")
	blogRouter.POST("/", web.CreateComment)
	blogRouter.GET("/:postID", web.GetPostComments)

}
