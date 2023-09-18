package services

import (
	"context"
	"fmt"
	"meffin-transactions-api/internal/models"
	"meffin-transactions-api/internal/repository"
	"strconv"
	"time"
)

type TransactionService interface {
	Create(ctx context.Context, request models.CreateTransactionRequest) (*models.Transaction, error)
	GetUserTransactions(ctx context.Context, userId string) ([]*models.Transaction, error)
	DeleteTransaction(ctx context.Context, transactionID string) error
	UpdateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
}

type TransactionServiceImpl struct {
	repository repository.TransactionRepository
}

// Enforces implementation of interface at compile time
var _ TransactionService = (*TransactionServiceImpl)(nil)

func NewTransactionService(transactionRepository repository.TransactionRepository) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		repository: transactionRepository,
	}
}

func (s *TransactionServiceImpl) Create(ctx context.Context, request models.CreateTransactionRequest) (*models.Transaction, error) {
	createdTransaction, err := s.repository.Create(ctx, &repository.RowTransaction{
		UserID:      request.UserID,
		Type:        request.Type,
		Description: request.Description,
		Amount:      request.Amount,
		IsFixed:     request.IsFixed,
		DayOfMonth:  int64(request.DayOfMonth),
		EndDate:     request.EndDate,
		Category:    request.Category,
	})
	if err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          strconv.FormatInt(createdTransaction.ID, 10),
		UserID:      createdTransaction.UserID,
		Type:        createdTransaction.Type,
		Description: createdTransaction.Description,
		Amount:      createdTransaction.Amount,
		IsFixed:     createdTransaction.IsFixed,
		DayOfMonth:  int(createdTransaction.DayOfMonth),
		EndDate:     createdTransaction.EndDate,
		Category:    createdTransaction.Category,
	}, nil
}

func (s *TransactionServiceImpl) GetUserTransactions(ctx context.Context, userID string) ([]*models.Transaction, error) {
	rowTransactions, err := s.repository.GetTransactionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return toTransactions(rowTransactions), nil
}

func (s *TransactionServiceImpl) DeleteTransaction(ctx context.Context, transactionID string) error {
	id, err := strconv.ParseInt(transactionID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid transaction ID: %v", err)
	}
	return s.repository.DeleteTransaction(ctx, id)
}

func (s *TransactionServiceImpl) UpdateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	transactionID, err := strconv.ParseInt(transaction.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid transaction ID: %v", err)
	}
	updatedTransaction, err := s.repository.UpdateTransaction(ctx, &repository.RowTransaction{
		ID:          transactionID,
		UserID:      transaction.UserID,
		Type:        transaction.Type,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		IsFixed:     transaction.IsFixed,
		DayOfMonth:  int64(transaction.DayOfMonth),
		EndDate:     transaction.EndDate,
		Category:    transaction.Category,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          strconv.FormatInt(updatedTransaction.ID, 10),
		UserID:      updatedTransaction.UserID,
		Type:        updatedTransaction.Type,
		Description: updatedTransaction.Description,
		Amount:      updatedTransaction.Amount,
		IsFixed:     updatedTransaction.IsFixed,
		DayOfMonth:  int(updatedTransaction.DayOfMonth),
		EndDate:     updatedTransaction.EndDate,
		Category:    updatedTransaction.Category,
	}, nil
}

func toTransactions(rowTransactions []*repository.RowTransaction) []*models.Transaction {
	var transactions []*models.Transaction

	for _, rowTransaction := range rowTransactions {
		transactions = append(transactions, &models.Transaction{
			ID:          strconv.FormatInt(rowTransaction.ID, 10),
			UserID:      rowTransaction.UserID,
			Type:        rowTransaction.Type,
			Description: rowTransaction.Description,
			Amount:      rowTransaction.Amount,
			IsFixed:     rowTransaction.IsFixed,
			DayOfMonth:  int(rowTransaction.DayOfMonth),
			EndDate:     rowTransaction.EndDate,
			Category:    rowTransaction.Category,
		})
	}

	return transactions
}
