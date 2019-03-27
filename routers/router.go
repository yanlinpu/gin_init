package routers

import (
	healthCtrl "gin_init/api/health"
	apiV1 "gin_init/api/v1"
	_ "gin_init/docs"
	"gin_init/middleware/ps"
	"gin_init/pkg/setting"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	if setting.AppSetting.Debug {
		r.Use(gin.Logger())
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// check := r.Group("health")
	// {
	// 	check.GET("check", healthCtrl.HealthCheck)
	// }
	r.GET("health/check", healthCtrl.HealthCheck)
	gV1 := r.Group("api/v1")
	{
		gV1.GET("get-test", ps.PageSizeCheck(), apiV1.GetParams)
		gV1.POST("post-test", ps.PageSizeCheck(), apiV1.PostParams)
		// gV1.Use(ps.PageSizeCheck()).GET("get-test", apiV1.GetParams)
		// gV1.Use(ps.PageSizeCheck()).POST("post-test", apiV1.PostParams)
	}

	return r
}
