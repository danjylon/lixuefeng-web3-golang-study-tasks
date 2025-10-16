package models

import "github.com/gin-gonic/gin"

// Pagination 分页请求参数
type Pagination struct {
	Page int `form:"page" json:"page" binding:"omitempty,min=1"`
	Size int `form:"size" json:"size" binding:"omitempty,min=1,max=100"`
}

// PaginatedResult 分页响应结果
type PaginatedResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Pages int         `json:"pages"` // 总页数
}

// NewPaginatedResult 创建分页响应
func NewPaginatedResult(list interface{}, total int64, page, size int) *PaginatedResult {
	pages := int((total + int64(size) - 1) / int64(size)) // 向上取整
	return &PaginatedResult{
		List:  list,
		Total: total,
		Page:  page,
		Size:  size,
		Pages: pages,
	}
}

// 1. 解析分页参数
func InitPagination(context *gin.Context) Pagination {
	var pagination Pagination
	if err := context.ShouldBindQuery(&pagination); err != nil {
		panic(err)
	}

	// 2. 设置默认值
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.Size == 0 {
		pagination.Size = 10
	}
	return pagination
}
