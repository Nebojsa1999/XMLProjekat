package config

import "os"

type Config struct {
	Port          string
	PostingDBHost string
	PostingDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("POSTING_SERVICE_PORT"),
		PostingDBHost: os.Getenv("POSTING_DB_HOST"),
		PostingDBPort: os.Getenv("POSTING_DB_PORT"),
	}
}
