package handlers

import (
	"github.com/gin-gonic/gin"
	"meffin-transactions-api/internal/models"
	"meffin-transactions-api/internal/services"
	"net/http"
	"strconv"
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

	var incomes []*models.Transaction
	var expenses []*models.Transaction

	for _, transaction := range transactions {
		if transaction.Type == models.Income {
			incomes = append(incomes, transaction)
		} else if transaction.Type == models.Expense {
			expenses = append(expenses, transaction)
		}
	}

	response := &models.UserTransactionsResponse{
		UserID:   userID,
		Incomes:  incomes,
		Expenses: expenses,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var request models.CreateTransactionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTransaction, err := h.transactionService.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, createdTransaction)
}

func (h *Handler) DeleteTransaction(c *gin.Context) {
	transactionID := c.Param("transaction_id")

	parsedTransactionID, err := strconv.ParseUint(transactionID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	err = h.transactionService.DeleteTransaction(c, uint(parsedTransactionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	var request models.Transaction

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTransaction, err := h.transactionService.UpdateTransaction(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, updatedTransaction)
}
