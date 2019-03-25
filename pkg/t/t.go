package t

import (
	"bytes"
	"encoding/json"
	"gin_init/pkg/database"
	"gin_init/pkg/logger"
	"gin_init/pkg/redis"
	_ "gin_init/pkg/setting"
	"gin_init/routers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func init() {
	database.Setup()
	redis.Setup()
	gin.SetMode(gin.ReleaseMode)
	// 清空数据
	// db := database.Conn()
	// db.Exec("TRUNCATE TABLE user;")
	// userInsertSql := `
	// 	insert into users (name) values
	// 	('c1'),
	// 	('c2')
	// `
	// db.Exec(userInsertSql)
}

func DoMethod(method, uri string, params map[string]interface{}, headers map[string]string) (rr Rest, code int) {
	var req *http.Request
	if method == "GET" {
		req = httptest.NewRequest("GET", uri, nil)
	} else {
		// 将参数转化为json比特流
		jsonByte, _ := json.Marshal(params)
		// 构造post请求，json数据以请求body的形式传递
		req = httptest.NewRequest(method, uri, bytes.NewReader(jsonByte))
	}

	// 初始化响应
	w := httptest.NewRecorder()
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 调用相应的handler接口
	router := routers.InitRouter()
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	body, _ := ioutil.ReadAll(result.Body)
	err := json.Unmarshal([]byte(body), &rr)
	if err != nil {
		logger.Error("JSON转换失败")
	}
	code = w.Code
	return
}
