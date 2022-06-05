package main

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/startup"
	cfg "github.com/Nebojsa1999/XMLProjekat/backend/connection_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
