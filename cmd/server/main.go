package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
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

	c := cron.New(cron.WithSeconds())

	// This schedules the job to run on the first day of every month at midnight
	_, err := c.AddFunc("0 0 0 1 * ?", func() {
		log.Println("Cron job triggered!")
		err := services.TransactionService.DeleteExpiredTransactions(context.Background())
		if err != nil {
			log.Printf("error deleting expired transactions: %s", err)
		}
		log.Println("successfully deleted expired transactions")
	})
	if err != nil {
		log.Fatalf("error setting up cron job: %s", err)
	}
	c.Start()

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
