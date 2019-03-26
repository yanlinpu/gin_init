package main

import (
	"fmt"
	"gin_init/pkg/database"
	"gin_init/pkg/redis"
	"gin_init/pkg/setting"
	"gin_init/routers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "gin_init/pkg/tracing"
	// opentracing "github.com/opentracing/opentracing-go"
)

func init() {
	database.Setup()
	redis.Setup()
	// tracer, closer := tracing.Init(setting.JaegerSetting.Server)
	// defer closer.Close()
	// opentracing.SetGlobalTracer(tracer)
}

// @title 服务
// @version 1.0
// @description 服务，以下为所有api接口。
// @host www.baidu.com
// @BasePath /v1
func main() {
	if !setting.AppSetting.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	port := setting.ServerSetting.Port
	routersInit := routers.InitRouter()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      routersInit,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
