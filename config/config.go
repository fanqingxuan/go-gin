package config

type Config struct {
	App   `yaml:"app"`
	Redis `yaml:"redis"`
	// Mysql sqlx.Config   `yaml:"Mysql"`
	Log `yaml:"log"`
}
