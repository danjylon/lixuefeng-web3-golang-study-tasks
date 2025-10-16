package models

import (
	"blogSystem/config"
	"gorm.io/gorm"
	"log"
)

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
	gorm.Model
	Username string    `gorm:"type:varchar(30);not null;unique"  json:"username" form:"username" binding:"required,max=30"`
	Pwd      string    `gorm:"type:varchar(100);not null" json:"pwd" form:"pwd" binding:"required,max=30"`
	Phone    string    `gorm:"type:varchar(20);not null;unique" json:"phone" form:"phone" binding:"required,e164,max=20"`
	Email    *string   `gorm:"type:varchar(30);unique" json:"email" form:"email" binding:"omitempty,email,max=30"`
	PostNum  int       `gorm:"default:0" json:"postNum" form:"postNum" binding:"omitempty,number"`
	Posts    []Post    `binding:"omitempty"`
	Comments []Comment `binding:"omitempty"`
}
type Post struct {
	gorm.Model           //包含ID uint `gorm:"primarykey"`，CreatedAt time.Time，UpdatedAt time.Time，DeletedAt DeletedAt `gorm:"index"`
	Title         string `gorm:"type:varchar(50);not null" json:"title" form:"title" binding:"required"`
	Content       string `gorm:"type:text;not null" json:"content" form:"content" binding:"required"`
	CommentStatus bool   `gorm:"default:false" json:"comment_status" form:"comment_status" binding:"omitempty"`
	// 用户，多对一
	UserID uint `json:"userID" form:"userID" binding:"omitempty"`
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" binding:"omitempty"`
	// 查询course时，gorm会自动查询选择该课程的students
	Comments []Comment `binding:"omitempty"`
}

type Comment struct {
	gorm.Model        //包含ID uint `gorm:"primarykey"`，CreatedAt time.Time，UpdatedAt time.Time，DeletedAt DeletedAt `gorm:"index"`
	Content    string `gorm:"type:varchar(255);not null" json:"content" form:"content" binding:"required"`
	// 多对一
	PostID uint `json:"postID" form:"postID" binding:"required"`
	//Post   Post `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;" binding:"omitempty"`
	// 用户，多对一
	UserID uint `json:"userID" form:"userID" binding:"omitempty"`
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" binding:"omitempty"`
}
