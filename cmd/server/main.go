package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"meffin-transactions-api/configs"
	"meffin-transactions-api/internal/router"
	"meffin-transactions-api/internal/services"
)

func main() {
	r := gin.Default()
	config := configs.Load()
	services := services.InitServices(config)
	router.Setup(r, services)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
