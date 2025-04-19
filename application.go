package main

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
	"log"
)

func main() {

	cfg, err := config.SetupEnv()
	if err != nil {
		log.Printf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)

}
