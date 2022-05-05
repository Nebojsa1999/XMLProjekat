package main

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/startup"
	cfg "github.com/Nebojsa1999/XMLProjekat/backend/user_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
