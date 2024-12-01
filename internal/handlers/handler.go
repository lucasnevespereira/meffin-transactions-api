package handlers

import (
	"log"
	"meffin-transactions-api/internal/models"
	"meffin-transactions-api/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	categoryService    services.ICategoryService
	transactionService services.TransactionService
}

func NewHandler(transactionService services.TransactionService, categoryService services.ICategoryService) *Handler {
	return &Handler{
		categoryService:    categoryService,
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

	if response.Incomes == nil && response.Expenses == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		return
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

	err := h.transactionService.DeleteTransaction(c, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

func (h *Handler) DeleteExpiredTransactions(c *gin.Context) {

	log.Printf("Webhook: Deleting expired transactions, date: %s", time.Now().Format(time.RFC3339))

	err := h.transactionService.DeleteExpiredTransactions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expired transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted expired transactions"})
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

func (h *Handler) CreateCategory(c *gin.Context) {
	var request models.CreateCategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := h.categoryService.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

func (h *Handler) GetUserCategories(c *gin.Context) {
	userID := c.Param("user_id")

	categories, err := h.categoryService.GetUserCategories(c, userID)
	if categories == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No categories found for the user"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching categories for the user."})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("category_id")

	err := h.categoryService.DeleteCategory(c, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var request models.Category

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := h.categoryService.UpdateCategory(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}
