package dao

import (
	. "blogSystem/config"
	. "blogSystem/models"
)

func CreateUser(user User) error {
	return GetDB().Create(&user).Error
}

func GetUserByUsername(username string) (User, error) {
	var user User
	if err := GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
