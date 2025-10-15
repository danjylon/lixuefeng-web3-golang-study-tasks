package dao

import (
	. "gorm_demo/config"
	. "gorm_demo/model"
)

func CreateUser(users []User) {
	GetDB().CreateInBatches(&users, len(users))
}
