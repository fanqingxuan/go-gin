package config

type Redis struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"` // no password set
	DB       int    `yaml:"db"`       // use default DB
}
