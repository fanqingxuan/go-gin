package redisx

type Config struct {
	Addr     string `yaml:"Addr"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}
