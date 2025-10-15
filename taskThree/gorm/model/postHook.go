package model

import (
	"gorm.io/gorm"
	"log"
)

/*
 * 针对Post的钩子函数，可以针对Post的crud，对Post的字段进行处理
 */
//Create是走BeforeSave、BeforeCreate、AfterCreate、AfterSave
func (t *Post) BeforeCreate(tx *gorm.DB) error {
	log.Println("BeforeCreate ...")
	return nil
}
func (t *Post) AfterCreate(tx *gorm.DB) error {
	log.Println("AfterCreate ...")
	return nil
}

// Save时，只走BeforeSave、AfterSave
func (t *Post) BeforeSave(tx *gorm.DB) error {
	log.Println("BeforeSave ...")
	return nil
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (t *Post) AfterSave(tx *gorm.DB) error {
	log.Println("AfterSave ...")
	userID := t.UserID
	if err := tx.Model(&User{}).Where("id = ?", userID).Update("post_num", gorm.Expr("post_num+1")).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (t *Post) BeforeUpdate(tx *gorm.DB) error {
	log.Println("BeforeUpdate ...")
	return nil
}
func (t *Post) AfterUpdate(tx *gorm.DB) error {
	log.Println("AfterUpdate ...")
	return nil
}
func (t *Post) BeforeDelete(tx *gorm.DB) error {
	log.Println("BeforeDelete ...")
	return nil
}

// 为 Post 模型添加一个钩子函数，在文章删除时自动更新用户的文章数量统计字段。
func (t *Post) AfterDelete(tx *gorm.DB) error {
	log.Println("AfterDelete ...")
	userID := t.UserID
	if err := tx.Model(&User{}).Where("id = ?", userID).Update("post_num", gorm.Expr("post_num-1")).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (t *Post) AfterFind(tx *gorm.DB) error {
	log.Println("AfterFind ...")
	return nil
}
