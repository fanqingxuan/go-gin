package config

import (
	"go-gin/internal/environment"

	"github.com/golang-module/carbon/v2"
)

func InitEnvironment() {

	carbon.SetDefault(carbon.Default{
		Layout:   instance.App.TimeFormat,
		Timezone: instance.App.TimeZone,
	})
	environment.SetEnvMode(instance.App.Mode)
	environment.SetTimeZone(instance.App.TimeZone)
}
