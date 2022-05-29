package config

import "os"

type Config struct {
	Port           string
	AgentAppDBHost string
	AgentAppDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           os.Getenv("AGENT_APP_PORT"),
		AgentAppDBHost: os.Getenv("AGENT_APP_DB_HOST"),
		AgentAppDBPort: os.Getenv("AGENT_APP_DB_PORT"),
	}
}
