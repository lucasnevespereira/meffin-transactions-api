package repository

import (
	"context"
	"time"
)

type TransactionRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, rowServer *RowTransaction) (*RowTransaction, error)
	GetTransactionsByUserID(ctx context.Context, userID string) ([]*RowTransaction, error)
	DeleteTransaction(ctx context.Context, transactionID string) error
	UpdateTransaction(ctx context.Context, updatedTransaction *RowTransaction) (*RowTransaction, error)
}

type RowTransaction struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	IsFixed     bool      `json:"is_fixed"`
	DayOfMonth  int       `json:"day_of_month"`
	EndDate     string    `json:"endDate"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (RowTransaction) TableName() string {
	return "transactions"
}
