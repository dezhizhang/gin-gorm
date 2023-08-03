package router

import (
	"gin-gorm/controller"
	docs "gin-gorm/docs"
	"gin-gorm/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router 路由
func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", service.Ping)
		v1.GET("/user", service.GetProblemList)
		v1.POST("/user/login", controller.Login)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
