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

func CreateBlog(context *gin.Context) {
	info := utils.GetUserInfo(context)
	fmt.Println("info:", info)
	var post Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		panic(err)
	}
	post.UserID = info.UserID
	fmt.Println("post:", post)
	err = dao.CreateBlog(post)
	if err != nil {
		panic(err)
	}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "保存成功",
	//})
	context.JSON(http.StatusOK, Success("创建成功"))
}

func GetBlogs(context *gin.Context) {
	pagination := InitPagination(context)
	// 3. 构建过滤条件
	filters := make(map[string]interface{})
	if title := context.Query("title"); title != "" {
		filters["title"] = title
	}
	result, err := dao.GetBlogs(&pagination, filters)
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, Success(result))

}
func GetOwnBlogs(context *gin.Context) {
	info := utils.GetUserInfo(context)
	fmt.Println("info:", info)
	pagination := InitPagination(context)
	// 3. 构建过滤条件
	filters := make(map[string]interface{})
	if title := context.Query("title"); title != "" {
		filters["title"] = title
	}
	result, err := dao.GetOwnBlogs(info.UserID, &pagination, filters)
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, Success(result))

}

func GetBlogById(context *gin.Context) {
	postId := context.Param("id")
	if postId == "" {
		panic(errors.New("请提交文章id"))
	}
	fmt.Println("postId:", postId)
	id, err := strconv.Atoi(postId)
	if err != nil {
		panic(err)
	}
	post, err := dao.GetBlogById(uint(id))
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, Success(post))
}

func UpdateBlog(context *gin.Context) {
	info := utils.GetUserInfo(context)
	fmt.Println("info:", info)
	var post Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		panic(err)
	}
	fmt.Println("post:", post)
	if post.ID == 0 {
		panic(errors.New("请提交文章id"))
	}
	postQuery, err := dao.GetBlogById(post.ID)
	if err != nil {
		panic(err)
	}
	if postQuery.UserID != info.UserID {
		panic(errors.New("仅可修改自己的文章"))
	}
	err = dao.UpdateBlog(post)
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, Success("修改成功"))
}

func DeleteBlogById(context *gin.Context) {
	info := utils.GetUserInfo(context)
	fmt.Println("info:", info)
	postId := context.Param("id")
	if postId == "" {
		panic(errors.New("请提交文章id"))
	}
	fmt.Println("postId:", postId)
	id, err := strconv.Atoi(postId)
	if err != nil {
		panic(err)
	}
	postQuery, err := dao.GetBlogById(uint(id))
	if err != nil {
		panic(err)
	}
	if postQuery.UserID != info.UserID {
		panic(errors.New("仅可删除自己的文章"))
	}
	err = dao.DeleteBlogById(postQuery)
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, Success("删除成功"))
}
