package middlewares

import (
	. "blogSystem/config"
	"blogSystem/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// 定义需要排除鉴权的接口路径
var excludePaths = []string{
	"/index",
	"/user/login",
	"/user/register",
}

// 权限校验
func AuthCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		//放行
		context.Next()
	}
}

func TokenCheck(context *gin.Context) {
	// 检查当前请求路径是否在排除列表中
	for _, path := range excludePaths {
		if context.Request.URL.Path == path ||
			(path == "/api/v1/public/*" && strings.HasPrefix(context.Request.URL.Path, "/api/v1/public/")) {
			context.Next()
			return
		}
	}
	//Authorization:=context.GetHeader("Authorization")
	//从请求头中获取token
	token := context.Request.Header.Get("Authorization")
	//log.Println("Authorization: ", token)
	token = strings.ReplaceAll(token, "Bearer ", "")
	//log.Println("token: ", token)
	//jwt解析token获得用户信息
	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Println("ParseToken:", err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token验证失败",
		})
		return
	}
	//组装redis的key，从redis中查询用户信息
	redisKey := claims.UserKey + ":" + claims.Username
	userInfo, err := GetRDB().Get(GetContext(), redisKey).Result()
	if err != nil {
		log.Println("rediserr:", err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token验证失败",
		})
		return
	}
	context.Set("user", userInfo)
	//log.Println("redis:userInfo: ", userInfo)
	//如果token有效，过期时间剩余20分钟以内时更新token有效时间
	err = utils.VerifyToken(claims)
	if err != nil {
		log.Println("VerifyToken:", err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "token验证失败",
		})
		return
	}
	//放行
	context.Next()
}
