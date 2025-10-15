package dao

import (
	. "gorm_demo/config"
	. "gorm_demo/model"
)

func CreatComments(comments []Comment) {
	GetDB().CreateInBatches(comments, len(comments))
}

func DeleteComment(comment Comment) {
	GetDB().Delete(&comment)
}
