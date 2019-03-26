package help

import (
	"encoding/json"
	"myfavorite/system/logger"
)

func MapToJSONString(param map[string]interface{}) (value string) {
	jsonString, err := json.Marshal(param)
	if err != nil {
		logger.Error("MapToJSONString error: ", err.Error())
		return
	}
	value = string(jsonString)
	return
}
