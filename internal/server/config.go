package server

type Config struct {
	BindAddr string
}

func NewConfig(bindAddr string) *Config {
	return &Config{
		BindAddr: bindAddr,
	}
}