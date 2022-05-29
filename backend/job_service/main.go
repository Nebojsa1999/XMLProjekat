package main

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
