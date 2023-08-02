package service

import (
	"gin-gorm/dao"
	"gin-gorm/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]

func Ping(ctx *gin.Context) {
	var user model.User
	dao.DB.Where("id=?", 123).Find(&user)
	ctx.JSON(http.StatusOK, &user)
}
