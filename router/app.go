package router

import (
	"gin-gorm/service"
	"github.com/gin-gonic/gin"
)

// Router 路由
func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", service.Ping)
	return r
}
