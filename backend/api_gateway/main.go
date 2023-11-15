package main

import (
	"api.accommodation.com/config"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	gin.SetMode(gin.ReleaseMode)
	clients := client.InitClients(&cfg)

	r := router.RouterInit(clients)
	r.Run(cfg.Address)
}
