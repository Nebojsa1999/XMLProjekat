package main

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/startup"
	cfg "github.com/Nebojsa1999/XMLProjekat/agent-app-backend/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
