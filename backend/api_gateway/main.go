package main

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/startup"
	"github.com/Nebojsa1999/XMLProjekat/backend/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
