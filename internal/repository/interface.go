package repository

import (
	"context"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, rowServer *RowTransaction) (*RowTransaction, error)
	GetTransactionsByUserID(ctx context.Context, userID string) ([]*RowTransaction, error)
	DeleteTransaction(ctx context.Context, transactionID uint) error
	UpdateTransaction(ctx context.Context, updatedTransaction *RowTransaction) (*RowTransaction, error)
}

type RowTransaction struct {
	gorm.Model
	UserID      string  `json:"userId"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	IsFixed     bool    `json:"is_fixed"`
	DayOfMonth  int     `json:"day_of_month"`
	EndDate     string  `json:"endDate"`
	Category    string  `json:"category"`
}

func (RowTransaction) TableName() string {
	return "transactions"
}
