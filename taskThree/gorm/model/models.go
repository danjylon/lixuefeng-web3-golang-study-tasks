package model

import (
	"gorm.io/gorm"
	"gorm_demo/config"
	"log"
)

/*
*
### 题目1：模型定义
- 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
  - 要求 ：
  - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
  - 编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
// 使用init函数，在项目启动时初始化并创建表
func InitTables() {
	//通过结构体生成表
	err := config.GetDB().Migrator().AutoMigrate(User{})
	if err != nil {
		log.Fatal(err)
		return
	}
	err = config.GetDB().Migrator().AutoMigrate(Post{})
	if err != nil {
		log.Fatal(err)
		return
	}
	err = config.GetDB().Migrator().AutoMigrate(Comment{})
	if err != nil {
		log.Fatal(err)
		return
	}
}

type User struct {
	gorm.Model        //包含ID uint `gorm:"primarykey"`，CreatedAt time.Time，UpdatedAt time.Time，DeletedAt DeletedAt `gorm:"index"`
	Name       string `gorm:"type:varchar(30);not null" json:"name" form:"name" binding:"required"`
	Phone      string `gorm:"type:varchar(20); unique" json:"phone" form:"phone" binding:"required,e164"`
	PostNum    int    `gorm:"default:0" json:"postNum" form:"postNum" binding:"omitempty,number"`
}

type Post struct {
	gorm.Model           //包含ID uint `gorm:"primarykey"`，CreatedAt time.Time，UpdatedAt time.Time，DeletedAt DeletedAt `gorm:"index"`
	Title         string `gorm:"type:varchar(50);not null" json:"title" form:"title" binding:"required"`
	Content       string `gorm:"type:text;unique;not null" json:"content" form:"content" binding:"required"`
	CommentStatus bool   `gorm:"default:false" json:"comment_status" form:"comment_status" binding:"required"`
	// 用户，多对一
	UserID uint `json:"userID" form:"userID" binding:"required"`
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" binding:"omitempty"` //`gorm:"references:Tno;constraint:OnDelete:CASCADE;"
	// 查询course时，gorm会自动查询选择该课程的students
	Comments []Comment `binding:"omitempty"`
}

type Comment struct {
	gorm.Model        //包含ID uint `gorm:"primarykey"`，CreatedAt time.Time，UpdatedAt time.Time，DeletedAt DeletedAt `gorm:"index"`
	Content    string `gorm:"type:varchar(255);not null" json:"content" form:"content" binding:"required"`
	// 多对一
	PostID uint `json:"postID" form:"postID" binding:"required"`
	Post   Post `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;" binding:"omitempty"`
}
