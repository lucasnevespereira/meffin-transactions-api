package repository

import (
	"context"
	"time"
)

type TransactionRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, rowServer *RowTransaction) (*RowTransaction, error)
	GetTransactionsByUserID(ctx context.Context, userID string) ([]*RowTransaction, error)
	DeleteTransaction(ctx context.Context, transactionID int64) error
	UpdateTransaction(ctx context.Context, updatedTransaction *RowTransaction) (*RowTransaction, error)
}

type RowTransaction struct {
	ID          int64     `db:"id"`
	UserID      string    `db:"user_id"`
	Type        string    `db:"type"`
	Description string    `db:"description"`
	Amount      float64   `db:"amount"`
	IsFixed     bool      `db:"is_fixed"`
	DayOfMonth  int64     `db:"day_of_month"`
	EndDate     string    `db:"endDate"`
	Category    string    `db:"category"`
	CreatedAt   time.Time `db:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt"`
}

func (RowTransaction) TableName() string {
	return "transactions"
}
