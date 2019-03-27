package health

import (
	"gin_init/pkg/database"
	"gin_init/pkg/logger"
	"gin_init/pkg/redis"

	"github.com/gin-gonic/gin"
)

type JsonApiOut struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HealthCheck(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			json := JsonApiOut{
				Code:    500,
				Message: "ok",
				Data:    err,
			}
			c.JSON(500, json)
		}
	}()

	// 检查mysql
	db := database.Conn()

	errors := db.GetErrors()
	if len(errors) > 0 {
		logger.Error("db connect error!", errors[0].Error())
		panic(errors[0])
	}

	err := db.DB().Ping()
	if err != nil {
		logger.Error("db connect error!", err.Error())
	}

	//检查redis
	rDB := redis.Conn()
	_, err = rDB.Ping().Result()
	if err != nil {
		logger.Error("redis connect fail!", err.Error())
		panic(err)
	}

	json := JsonApiOut{
		Code:    200,
		Message: "ok",
		Data:    nil,
	}
	c.JSON(200, &json)
}
