package config

import "os"

type Config struct {
	Port        string
	UserHost    string
	UserPort    string
	PostingHost string
	PostingPort string
	JobHost     string
	JobPort     string
}

func NewConfig() *Config {
	return &Config{
		Port:        os.Getenv("GATEWAY_PORT"),
		UserHost:    os.Getenv("USER_SERVICE_HOST"),
		UserPort:    os.Getenv("USER_SERVICE_PORT"),
		PostingHost: os.Getenv("POSTING_SERVICE_HOST"),
		PostingPort: os.Getenv("POSTING_SERVICE_PORT"),
		JobHost:     os.Getenv("JOB_SERVICE_HOST"),
		JobPort:     os.Getenv("JOB_SERVICE_PORT"),
	}
}
