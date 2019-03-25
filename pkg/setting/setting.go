package setting

import (
	"gin_init/pkg/logger"
	"os"
	"time"

	"github.com/go-ini/ini"
)

// App app 相关配置
type App struct {
	AppName        string
	Debug          bool
	MaxPageSizeNum int
}

//Server 系统配置
type Server struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Mysql 数据库配置
type Mysql struct {
	User     string
	Password string
	Host     string
}

// Redis redis配置
type Redis struct {
	Host         string
	Password     string
	Database     int
	PoolSize     int
	MinIdleConns int
}

// Jaeger Jaeger
type Jaeger struct {
	Host   string
	Server string
}

var (
	AppSetting      = &App{}    // AppSetting app配置读取
	ServerSetting   = &Server{} // ServerSetting app配置读取
	DatabaseSetting = &Mysql{}  // DatabaseSetting database
	RedisSetting    = &Redis{}  // RedisSetting redis配置读取
	JaegerSetting   = &Jaeger{} // JaegerSetting redis配置读取
	cfg             *ini.File
)

func init() {
	var err error

	configEnv := os.Getenv("GINENV")
	if configEnv != "production" && configEnv != "testing" {
		configEnv = "development"
	}

	pwdDir, err := os.Getwd()
	if err != nil {
		logger.Fatalf("Getwd err: %v", err)
		os.Exit(1)
	}

	if configEnv == "testing" {
		cfg, err = ini.Load(pwdDir + "/../../conf/" + configEnv + "/app.ini")
	} else {
		cfg, err = ini.Load(pwdDir + "/conf/" + configEnv + "/app.ini")
	}

	if err != nil {
		logger.Fatalf("Fail to parse 'conf/app.ini': %v", err)
		os.Exit(1)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("mysql", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("jaeger", JaegerSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logger.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
		os.Exit(1)
	}
}
