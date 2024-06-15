package config

type Config struct {
	App `yaml:"App"`
	// Redis redisx.Config `yaml:"Redis"`
	// Mysql sqlx.Config   `yaml:"Mysql"`
}
