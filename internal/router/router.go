package router

import (
	"meffin-transactions-api/internal/handlers"
	"meffin-transactions-api/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, services *services.Services) {

	router.Use(cors.Default())

	h := handlers.NewHandler(services.TransactionService, services.CategoryService)

	router.GET("/health", handlers.Health)
	router.GET("/users/:user_id/transactions", h.GetUserTransactions)

	router.POST("/transactions", h.CreateTransaction)
	router.PUT("/transactions", h.UpdateTransaction)
	router.DELETE("/transactions/:transaction_id", h.DeleteTransaction)
	router.DELETE("/transactions/expired", h.DeleteExpiredTransactions)

	router.GET("/users/:user_id/categories", h.GetUserCategories)

	router.POST("/categories", h.CreateCategory)
	router.PUT("/categories", h.UpdateCategory)
	router.DELETE("/categories/:category_id", h.DeleteCategory)

	router.GET("/webhooks/delete-expired-transactions", h.DeleteExpiredTransactions)

	router.NoRoute(handlers.NoRoute)
}
