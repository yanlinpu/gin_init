package v1_test

import (
	_ "gin_init/pkg/setting"
	tt "gin_init/pkg/t"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiGet(t *testing.T) {
	url := "/api/v1/get-test?page=1&pageSize=20"
	params := make(map[string]interface{})
	headers := make(map[string]string)
	//发起请求
	rr, httpCode := tt.DoMethod("GET", url, params, headers)
	assert.Equal(t, httpCode, 200, "HttpCode错误")
	assert.Equal(t, rr.Code, 200, "自定义code错误")
	assert.Equal(t, rr.Msg, "success", "提示信息错误")
}

func TestApiPost(t *testing.T) {
	url := "/api/v1/post-test"
	params := make(map[string]interface{})
	headers := make(map[string]string)
	params["page"] = 1
	params["pageSize"] = 20
	//发起请求
	rr, httpCode := tt.DoMethod("POST", url, params, headers)
	assert.Equal(t, httpCode, 200, "HttpCode错误")
	assert.Equal(t, rr.Code, 200, "自定义code错误")
	assert.Equal(t, rr.Msg, "success", "提示信息错误")
}
