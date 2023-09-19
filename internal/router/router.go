package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"meffin-transactions-api/internal/handlers"
	"meffin-transactions-api/internal/services"
)

func Setup(router *gin.Engine, services *services.Services) {

	router.Use(cors.Default())

	h := handlers.NewHandler(services.TransactionService)

	router.GET("/health", handlers.Health)
	router.GET("/users/:user_id/transactions", h.GetUserTransactions)

	router.POST("/transactions", h.CreateTransaction)
	router.PUT("/transactions", h.UpdateTransaction)
	router.DELETE("/transactions/:transaction_id", h.DeleteTransaction)
	router.DELETE("/transactions/expired", h.DeleteExpiredTransactions)

	router.NoRoute(handlers.NoRoute)
}
