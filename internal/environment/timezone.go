package environment

import "os"

func SetTimeZone(val string) {
	err := os.Setenv("TZ", val)
	if err != nil {
		panic("设置环境变量失败:" + err.Error())
	}
}
