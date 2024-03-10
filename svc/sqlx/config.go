package sqlx

type Config struct {
	DataSource   string `yaml:"DataSource"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
}
