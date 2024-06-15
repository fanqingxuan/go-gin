package config

type Config struct {
	App `yaml:"app"`
	// Redis redisx.Config `yaml:"Redis"`
	// Mysql sqlx.Config   `yaml:"Mysql"`
	Log `yaml:"log"`
}
