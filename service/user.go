package service

import (
	"gin-gorm/dao"
	"gin-gorm/model"
)

func Login(username string, password string) (*model.User, error) {
	var user model.User
	err := dao.DB.Model(&user).Where("username = ? And password = ?", username, password).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
