package handlers

import (
	"github.com/gin-gonic/gin"
	"meffin-transactions-api/internal/models"
	"meffin-transactions-api/internal/services"
	"net/http"
)

type Handler struct {
	transactionService services.TransactionService
}

func NewHandler(transactionService services.TransactionService) *Handler {
	return &Handler{
		transactionService: transactionService,
	}
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}

func NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Not Found"})
}

func (h *Handler) GetUserTransactions(c *gin.Context) {
	userID := c.Param("user_id")

	transactions, err := h.transactionService.GetUserTransactions(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions for the user."})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTransaction, err := h.transactionService.Create(c, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, createdTransaction)
}
