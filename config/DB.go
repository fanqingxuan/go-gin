package config

type DB struct {
	DSN          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"max-open-conn"`
	MaxIdleConns int    `yaml:"max-idle-conn"`
}
