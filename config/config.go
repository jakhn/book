package config

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
}

func Load() Config {

	var cfg Config

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "jakh"
	cfg.PostgresDatabase = "my_db"
	cfg.PostgresPassword = "00"
	cfg.PostgresPort = "5432"

	return cfg
}
