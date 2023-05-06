package main

import (
	"api.accommodation.com/config"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/router"
)

func main() {
	cfg := config.GetConfig()

	clients := client.InitClients(&cfg)

	r := router.RouterInit(clients)

	r.Run(cfg.Address)
}
