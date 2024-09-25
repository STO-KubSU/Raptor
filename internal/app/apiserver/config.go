package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	FrontendUrl string `toml:"frontend_url"`
	LogLevel    string `toml:"log_level"`
}

// Конфигурация сервера
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
