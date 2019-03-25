package redis

import (
	"gin_init/pkg/logger"
	"gin_init/pkg/setting"

	"github.com/go-redis/redis"
)

var redisConn *redis.Client

// Setup redis 连接
func Setup() {
	connect()
}

// connect connect redis
func connect() {
	redisConn = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		DB:       setting.RedisSetting.Database, // use default DB
		PoolSize: setting.RedisSetting.PoolSize,
	})

	_, err := redisConn.Ping().Result()
	if err != nil {
		logger.Error("redis connect fail!")
	}
}

//Conn reids Conn
func Conn() *redis.Client {
	return redisConn
}

//GetZ 有序队列元素
func GetZ(Score float64, Member interface{}) redis.Z {
	var z redis.Z
	z.Score = Score
	z.Member = Member
	return z
}
