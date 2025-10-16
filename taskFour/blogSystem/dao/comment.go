package dao

import (
	. "blogSystem/config"
	. "blogSystem/models"
	"fmt"
	"gorm.io/gorm"
)

func CreateComment(comment Comment) error {
	if err := GetDB().Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func GetPostComments(postId uint, pagination *Pagination) (paginatedResult *PaginatedResult, err error) {
	var total int64
	var comments []Comment

	if err := GetDB().Model(&Comment{}).Where("post_id = ?", postId).Count(&total).Error; err != nil {
		return nil, err
	}
	fmt.Println("total:", total)
	// 3. 分页查询
	offset := (pagination.Page - 1) * pagination.Size
	if err := GetDB().Offset(offset).Limit(pagination.Size).Order("created_at DESC").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username") // 明确指定需要的字段
		}).Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return NewPaginatedResult(comments, total, pagination.Page, pagination.Size), nil

}
