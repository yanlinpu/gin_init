package help

import (
	"gin_init/pkg/logger"
	"strconv"
)

func StringToInt(param string) int {
	vaule, err := strconv.Atoi(param)
	if err != nil {
		logger.Info("StringToInt error :", err.Error())
		return 0
	}
	return vaule
}
