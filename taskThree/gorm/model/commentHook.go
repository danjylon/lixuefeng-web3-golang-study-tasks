package model

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

/*
 * 针对Comment的钩子函数，可以针对Comment的crud，对Comment的字段进行处理
 */
func (t *Comment) BeforeCreate(tx *gorm.DB) error {
	log.Println("BeforeCreate ...")
	return nil
}
func (t *Comment) AfterCreate(tx *gorm.DB) error {
	log.Println("AfterCreate ...")
	return nil
}
func (t *Comment) BeforeSave(tx *gorm.DB) error {
	log.Println("BeforeSave ...")
	return nil
}
func (t *Comment) AfterSave(tx *gorm.DB) error {
	log.Println("AfterSave ...")
	post := Post{Model: gorm.Model{ID: t.PostID}}
	if err := tx.Model(&Post{}).Where(&post).Take(&post).Error; err != nil {
		log.Println("err:", err)
		return err
	}
	fmt.Printf("post:%v\n", post)
	if !post.CommentStatus {
		if err := tx.Model(&Post{}).Where("id = ?", post.ID).Update("comment_status", true).Error; err != nil {
			log.Println("err:", err)
			return err
		}
	}
	return nil
}
func (t *Comment) BeforeUpdate(tx *gorm.DB) error {
	log.Println("BeforeUpdate ...")
	return nil
}
func (t *Comment) AfterUpdate(tx *gorm.DB) error {
	log.Println("AfterUpdate ...")
	return nil
}
func (t *Comment) BeforeDelete(tx *gorm.DB) error {
	log.Println("BeforeDelete ...")
	//删除前先根据id查一下要删除的comment的完整信息
	if err := tx.Model(&Comment{}).Where("id = ?", t.ID).Take(t).Error; err != nil {
		log.Println("err:", err)
		return err
	}
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (t *Comment) AfterDelete(tx *gorm.DB) error {
	log.Println("AfterDelete ...")
	//fmt.Printf("comment: %#v\n", t)
	var commentsCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", t.PostID).Count(&commentsCount).Error; err != nil {
		log.Println("err:", err)
		return err
	}
	fmt.Printf("commentsCount: %v\n", commentsCount)
	//如果该文章剩余的评论数为0，将状态改为false
	if commentsCount == 0 {
		if err := tx.Model(&Post{}).Where("id = ?", t.PostID).Update("comment_status", false).Error; err != nil {
			log.Println("err:", err)
			return err
		}
	}
	return nil
}
func (t *Comment) AfterFind(tx *gorm.DB) error {
	log.Println("AfterFind ...")
	return nil
}
