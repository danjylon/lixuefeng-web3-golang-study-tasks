package utils

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(context *gin.Context) Claims {
	value, exists := context.Get("user")
	if !exists {
		panic(errors.New("未能获取用户信息"))
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code": http.StatusInternalServerError,
		//	"msg":  "未能获取用户信息",
		//})
	}
	var claims Claims
	err := json.Unmarshal([]byte(value.(string)), &claims)
	if err != nil {
		panic(errors.New(err.Error()))
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code": http.StatusInternalServerError,
		//	"msg":  err.Error(),
		//})
	}
	return claims
}
