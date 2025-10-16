package web

import (
	"blogSystem/dao"
	. "blogSystem/models"
	"blogSystem/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateComment(context *gin.Context) {
	info := utils.GetUserInfo(context)
	fmt.Println("info:", info)
	var comment Comment
	err := context.ShouldBindJSON(&comment)
	if err != nil {
		panic(err)
	}
	//根据postID查询文章是否存在
	_, err = dao.GetBlogById(comment.PostID)
	if err != nil {
		panic(errors.New("文章不存在"))
	}
	comment.UserID = info.UserID
	fmt.Println("comment:", comment)
	err = dao.CreateComment(comment)
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, Success("评论成功"))
}

func GetPostComments(context *gin.Context) {
	postId := context.Param("postID")
	if postId == "" {
		panic(errors.New("请提交文章id"))
	}
	fmt.Println("postId:", postId)
	id, err := strconv.Atoi(postId)
	if err != nil {
		panic(err)
	}
	//根据postID查询文章是否存在
	_, err = dao.GetBlogById(uint(id))
	if err != nil {
		panic(errors.New("文章不存在"))
	}
	pagination := InitPagination(context)
	comments, err := dao.GetPostComments(uint(id), &pagination)
	if err != nil {
		panic(errors.New("查询错误"))
	}
	context.JSON(http.StatusOK, Success(comments))
}
