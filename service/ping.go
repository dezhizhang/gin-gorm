package service

import (
	"gin-gorm/dao"
	"gin-gorm/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(ctx *gin.Context) {
	var user model.User
	dao.DB.Where("id=?", 123).Find(&user)
	ctx.JSON(http.StatusOK, &user)
}
