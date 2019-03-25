package e

import (
	"fmt"
	"gin_init/pkg/setting"
)

//MsgFlags 映射数组
var MsgFlags = map[int]string{
	SUCCESS:          "success",
	VALIDATIONERROR:  "params error",
	NOTFOUND:         "not find",
	ERROR:            "fail",
	PAGESIZEMAXERROR: fmt.Sprintf("每页不能超过%d个", setting.AppSetting.MaxPageSizeNum),
}

//GetMsg 错误码映射返回
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
