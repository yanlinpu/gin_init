package help

import (
	"gin_init/pkg/e"

	"github.com/gin-gonic/gin"
)

// Gin 上下文
type Gin struct {
	C *gin.Context
}

//Response 通一返回方法
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
	return
}
