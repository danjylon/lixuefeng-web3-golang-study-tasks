package routers

import (
	"blogSystem/web"
	"github.com/gin-gonic/gin"
)

func InitBlog(engine *gin.Engine) {
	//http://localhost:8888/user/v1
	//userRouter := engine.Group("/user", auth.TokenCheck, auth.AuthCheck)//只对/user路由鉴权
	blogRouter := engine.Group("/blog")
	blogRouter.POST("/", web.CreateBlog)
	blogRouter.PUT("/", web.UpdateBlog)
	blogRouter.GET("/", web.GetBlogs)
	blogRouter.GET("/myBlogs", web.GetOwnBlogs)
	blogRouter.GET("/:id", web.GetBlogById)
	blogRouter.DELETE("/:id", web.DeleteBlogById)

}
