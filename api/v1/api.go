package v1

import (
	error "gin_init/pkg/e"
	"gin_init/pkg/help"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
}

// GetParams ...
// @Summary Get请求测试
// @Description Get请求测试
// @Tags test
// @Accept  json
// @Produce  json
// @Param page body string true "page required"
// @Param pageSize body string true "pageSize required"
// @Success 200 {string} json "{"code":200,"data":{"page":1,"pageSize":20},"msg":"success"}"
// @Failure 400 {string} json "{"code":400,"data":"Key: 'Params.PageSize' Error:Field validation for 'PageSize' failed on the 'required' tag","msg":"params error"}"
// @Failure 405 {string} json "{"code":405,"data":null,"msg":"每页不能超过50个"}"
// @Router /api/v1/get-test [GET]
func GetParams(c *gin.Context) {
	var (
		params   Params
		response = help.Gin{C: c}
		data     = map[string]interface{}{}
	)
	err := c.BindQuery(&params)
	if err != nil {
		response.Response(http.StatusOK, error.VALIDATIONERROR, err.Error())
		return
	}
	data = map[string]interface{}{
		"page":     params.Page,
		"pageSize": params.PageSize,
	}
	response.Response(http.StatusOK, error.SUCCESS, data)
}

// PostParams ...
// @Summary POST请求测试
// @Description POST请求测试
// @Tags test
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"page":1,"pageSize":20},"msg":"success"}"
// @Failure 400 {string} json "{"code":400,"data":"Key: 'Params.PageSize' Error:Field validation for 'PageSize' failed on the 'required' tag","msg":"params error"}"
// @Failure 405 {string} json "{"code":405,"data":null,"msg":"每页不能超过50个"}"
// @Router /api/v1/post-test [POST]
func PostParams(c *gin.Context) {
	var (
		params   Params
		response = help.Gin{C: c}
		data     = map[string]interface{}{}
	)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.Response(http.StatusOK, error.VALIDATIONERROR, err.Error())
		return
	}
	data = map[string]interface{}{
		"page":     params.Page,
		"pageSize": params.PageSize,
	}
	response.Response(http.StatusOK, error.SUCCESS, data)
}
