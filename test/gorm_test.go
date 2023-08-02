package test

import (
	"fmt"
	"gin-gorm/dao"
	"gin-gorm/model"
	"testing"
)

func TestGorm(t *testing.T) {
	var user model.User
	tx := dao.DB.Model(&model.User{}).Find(&user)
	if tx.RowsAffected == 1 {
		panic(tx.Error.Error())
	}
	fmt.Println(user.Name)
}
