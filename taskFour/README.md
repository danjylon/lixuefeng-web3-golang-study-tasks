# 说明
使用了redis7和mysql
# 1 用户认证与授权
## 1.1 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。 /web/user.go
### 1.1.1 注册
http://localhost:9999/user/register POST
{
"username":"rose",
"pwd":"123456",
"pwdConfirm":"123456",
"phone": "+8613444444444"
}
### 1.1.2 登录
http://localhost:9999/user/login POST
{
    "username":"tom", //jerry，jack
    "pwd":"123456"
}
## 1.2 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。 使用redis，/web/user.go->Login, /middlewares/auth.go->TokenCheck, /utils/jwt.go
# 2 文章管理功能
## 2.1 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
http://localhost:9999/blog POST Header中加入Authorization=Bearer jwt
### post1
{
    "title": "关联查询",
    "content": "基于上述博客系统的模型定义。\n  - 要求 ：\n  - 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。\n  - 编写Go代码，使用Gorm查询评论数量最多的文章信息。"
}
### post2
{
    "title":"模型定义",
    "content":"假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。\n  - 要求 ：\n  - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。\n  - 编写Go代码，使用Gorm创建这些模型对应的数据库表"
}
### post3
{
    "title": "最长公共前缀",
    "content": "考察：字符串处理、循环嵌套\n题目：查找字符串数组中的最长公共前缀"
}
### post4
{
    "title": "删除有序数组中的重复项",
    "content": "删除有序数组中的重复项给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。\n可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。"
}
### post5
{
    "title": "实现类型安全映射",
    "content": "假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。\n要求 ：\n定义一个 Book 结构体，包含与 books 表对应的字段。\n编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。"
}
### post6
{
    "title": "使用SQL扩展库进行查询",
    "content": "假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。\n要求 ：\n编写Go代码，使用Sqlx查询 employees 表中所有部门为 \"技术部\" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。\n编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。"
}
### post7
{
    "title": "钩子函数",
    "content": "继续使用博客系统的模型。\n  - 要求 ：\n  - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。\n  - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 \"无评论\"。"
}
## 2.2 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
### 2.2.1 获取所有文章列表
#### 2.2.1.1 查看自己的文章列表
http://localhost:9999/blog/myBlogs?tile=关联 GET Header中加入Authorization=Bearer jwt
#### 2.2.1.2 查看所有文章列表
http://localhost:9999/blog?title=关联 GET Header中加入Authorization=Bearer jwt
### 2.2.2 获取单个文章的详细信息
http://localhost:9999/blog/1 GET Header中加入Authorization=Bearer jwt
## 2.3 实现文章的更新功能，只有文章的作者才能更新自己的文章。
http://localhost:9999/blog PUT Header中加入Authorization=Bearer jwt
{
    "id": 1,
    "title": "模型定义2",
    "content": "假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。\n  - 要求 ：\n  - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。\n  - 编写Go代码，使用Gorm创建这些模型对应的数据库表"
}
## 2.4 实现文章的删除功能，只有文章的作者才能删除自己的文章。
http://localhost:9999/blog/1 DELETE Header中加入Authorization=Bearer jwt
# 3 评论功能
## 3.1 实现评论的创建功能，已认证的用户可以对文章发表评论。
http://localhost:9999/comment POST  Header中加入Authorization=Bearer jwt
{
    "content":"感谢楼主分享",
    "postID":3
}
## 3.2 实现评论的读取功能，支持获取某篇文章的所有评论列表。
http://localhost:9999/comment/:postID
http://localhost:9999/comment/3 GET Header中加入Authorization=Bearer jwt
# 4 错误处理与日志记录
## 4.1 对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
## 4.2 使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。