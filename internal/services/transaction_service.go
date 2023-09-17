package services

import (
	"context"
	"meffin-transactions-api/internal/models"
	"meffin-transactions-api/internal/repository"
)

type TransactionService interface {
	Create(ctx context.Context, request models.CreateTransactionRequest) (*models.Transaction, error)
	GetUserTransactions(ctx context.Context, userId string) ([]*models.Transaction, error)
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
		DayOfMonth:  request.DayOfMonth,
		EndDate:     request.EndDate,
		Category:    request.Category,
	})
	if err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          createdTransaction.ID,
		UserID:      createdTransaction.UserID,
		Type:        createdTransaction.Type,
		Description: createdTransaction.Description,
		Amount:      createdTransaction.Amount,
		IsFixed:     createdTransaction.IsFixed,
		DayOfMonth:  createdTransaction.DayOfMonth,
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

func toTransactions(rowTransactions []*repository.RowTransaction) []*models.Transaction {
	var transactions []*models.Transaction

	for _, rowTransaction := range rowTransactions {
		transactions = append(transactions, &models.Transaction{
			ID:          rowTransaction.ID,
			UserID:      rowTransaction.UserID,
			Type:        rowTransaction.Type,
			Description: rowTransaction.Description,
			Amount:      rowTransaction.Amount,
			IsFixed:     rowTransaction.IsFixed,
			DayOfMonth:  rowTransaction.DayOfMonth,
			EndDate:     rowTransaction.EndDate,
			Category:    rowTransaction.Category,
		})
	}

	return transactions
}
