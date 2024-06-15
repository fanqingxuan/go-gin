package config

type App struct {
	Port     string `yaml:"port"`
	Debug    bool   `yaml:"debug"`
	TimeZone string `yaml:"timezone"`
}
