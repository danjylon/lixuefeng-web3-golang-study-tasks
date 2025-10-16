package dao

import (
	. "blogSystem/config"
	. "blogSystem/models"
	"errors"
	"fmt"
)

func CreateBlog(post Post) (err error) {
	if err = GetDB().Create(&post).Error; err != nil {
		return
	}
	return nil
}

func GetBlogs(pagination *Pagination, filters map[string]interface{}) (paginatedResult *PaginatedResult, err error) {
	var total int64
	var posts []Post
	db := GetDB()
	if title, ok := filters["title"].(string); ok {
		fmt.Println("title:", title)
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if err := db.Model(&Post{}).Count(&total).Error; err != nil {
		return nil, err
	}
	fmt.Println("total:", total)
	// 3. 分页查询
	offset := (pagination.Page - 1) * pagination.Size
	if err := db.Offset(offset).Limit(pagination.Size).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}

	return NewPaginatedResult(posts, total, pagination.Page, pagination.Size), nil
}

func GetOwnBlogs(userID uint, pagination *Pagination, filters map[string]interface{}) (paginatedResult *PaginatedResult, err error) {
	var total int64
	var posts []Post
	db := GetDB().Where("user_id = ?", userID)
	if title, ok := filters["title"].(string); ok {
		fmt.Println("title:", title)
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if err := db.Model(&Post{}).Count(&total).Error; err != nil {
		return nil, err
	}
	fmt.Println("total:", total)
	// 3. 分页查询
	offset := (pagination.Page - 1) * pagination.Size
	if err := db.Offset(offset).Limit(pagination.Size).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}

	return NewPaginatedResult(posts, total, pagination.Page, pagination.Size), nil
}

func GetBlogById(id uint) (post Post, err error) {
	if err = GetDB().Where("id = ?", id).First(&post).Error; err != nil {
		return post, errors.New("未查询到该文章")
	}
	return
}

func UpdateBlog(post Post) error {
	if err := GetDB().Where("id= ?", post.ID).Updates(&post).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBlogById(post Post) error {
	if err := GetDB().Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
