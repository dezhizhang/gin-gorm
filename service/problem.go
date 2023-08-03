package service

import (
	"gin-gorm/dao"
	"gin-gorm/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProblemList(c *gin.Context) {
	var user []model.User
	//page, _ := strconv.Atoi(c.DefaultQuery("page", global.Page))
	//size, _ := strconv.Atoi(c.DefaultQuery("size", global.Size))

	err := dao.DB.Model(&model.User{}).Find(&user).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
	})

}
