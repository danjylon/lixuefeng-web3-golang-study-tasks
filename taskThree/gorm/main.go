package main

import (
	"fmt"
	"gorm.io/gorm"
	. "gorm_demo/dao"
	. "gorm_demo/model"
)

/*
*
## 进阶gorm
### 题目1：模型定义
- 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
  - 要求 ：
  - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
  - 编写Go代码，使用Gorm创建这些模型对应的数据库表。

### 题目2：关联查询
- 基于上述博客系统的模型定义。
  - 要求 ：
  - 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
  - 编写Go代码，使用Gorm查询评论数量最多的文章信息。

### 题目3：钩子函数
- 继续使用博客系统的模型。
  - 要求 ：
  - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
  - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func main() {
	user := new(User)
	fmt.Println(user)
	//db := config.GetDB()
	//fmt.Println(db)
	//InitTables()
	//users := []User{
	//	{Name: "张三", Phone: "+8613111111111"},
	//	{Name: "李四", Phone: "+8613222222222"},
	//	{Name: "王五", Phone: "+8613333333333"},
	//	{Name: "赵六", Phone: "+8613444444444"},
	//}
	//CreateUser(users)
	//posts := []Post{
	//	{Title: "模型定义", Content: "假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。\n  - 要求 ：\n  - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。\n  - 编写Go代码，使用Gorm创建这些模型对应的数据库表", UserID: 1},
	//	{Title: "关联查询", Content: "基于上述博客系统的模型定义。\n  - 要求 ：\n  - 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。\n  - 编写Go代码，使用Gorm查询评论数量最多的文章信息。", UserID: 1},
	//	{Title: "钩子函数", Content: "继续使用博客系统的模型。\n  - 要求 ：\n  - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。\n  - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 \"无评论\"。", UserID: 1},
	//	{Title: "使用SQL扩展库进行查询", Content: "假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。\n要求 ：\n编写Go代码，使用Sqlx查询 employees 表中所有部门为 \"技术部\" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。\n编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。", UserID: 2},
	//	{Title: "实现类型安全映射", Content: "假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。\n要求 ：\n定义一个 Book 结构体，包含与 books 表对应的字段。\n编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。", UserID: 2},
	//	{Title: "删除有序数组中的重复项", Content: "删除有序数组中的重复项给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。\n可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。", UserID: 3},
	//	{Title: "最长公共前缀", Content: "考察：字符串处理、循环嵌套\n题目：查找字符串数组中的最长公共前缀", UserID: 4},
	//}
	//CreatePost(posts)
	//comments := []Comment{
	//	{Content: "good", PostID: 8},
	//	{Content: "very good", PostID: 9},
	//	{Content: "一般", PostID: 10},
	//	{Content: "excellent", PostID: 8},
	//	{Content: "not bad", PostID: 8},
	//	{Content: "感谢楼主分享", PostID: 9},
	//	{Content: "再接再厉", PostID: 10},
	//	{Content: "谢谢分享", PostID: 11},
	//}
	//CreatComments(comments)

	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	posts := GetPosts(1)
	for _, post := range posts {
		fmt.Printf("post: %v\n", post)
	}

	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	post := GetMostComentsPost()
	fmt.Printf("post: %v\n", post)

	//删除评论
	DeleteComment(Comment{Model: gorm.Model{ID: 13}})
}
