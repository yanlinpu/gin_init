package ps

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_init/pkg/e"
	"gin_init/pkg/help"
	"gin_init/pkg/logger"
	"gin_init/pkg/setting"
	// "gin_init/pkg/util"
)

type pp struct {
	PageSize int `json:"pageSize"`
}

func PageSizeCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			response = help.Gin{C: c}
			// c1       = c.Copy()
			pageSize int
		)

		if c.Request.Method == "GET" {
			pageSize = help.StringToInt(c.Query("pageSize"))
		} else {
			var (
				bodyBytes []byte
				pageS     pp
			)
			// json解析
			if c.Request.Body != nil {
				// bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
				bodyBytes, _ = c.GetRawData()
				err := json.Unmarshal(bodyBytes, &pageS)
				if err != nil {
					logger.Error(err)
				} else {
					pageSize = pageS.PageSize
				}
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // 关键点
			}
		}
		if pageSize > setting.AppSetting.MaxPageSizeNum {
			response.Response(http.StatusOK, e.PAGESIZEMAXERROR, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
