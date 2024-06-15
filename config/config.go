package config

type Config struct {
	App   `yaml:"app"`
	Redis `yaml:"redis"`
	DB    `yaml:"db"`
	Log   `yaml:"log"`
}
