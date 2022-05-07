package main

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/startup"
	cfg "github.com/Nebojsa1999/XMLProjekat/backend/posting_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
