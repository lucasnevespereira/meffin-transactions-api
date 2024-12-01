package repository

import (
	"context"
	"time"
)

type IRepository interface {
	AutoMigrate() error
	Create(ctx context.Context, rowServer *RowTransaction) (*RowTransaction, error)
	GetTransactionsByUserID(ctx context.Context, userID string) ([]*RowTransaction, error)
	DeleteTransaction(ctx context.Context, transactionID int64) error
	DeleteExpiredTransactions(ctx context.Context) error
	UpdateTransaction(ctx context.Context, updatedTransaction *RowTransaction) (*RowTransaction, error)
	CreateCategory(ctx context.Context, rowCategory *RowCategory) (*RowCategory, error)
	GetCategoriesByUserID(ctx context.Context, userID string) ([]*RowCategory, error)
	DeleteCategory(ctx context.Context, categoryID int64) error
	UpdateCategory(ctx context.Context, updatedCategory *RowCategory) (*RowCategory, error)
}

type RowTransaction struct {
	ID          int64     `db:"id"`
	UserID      string    `db:"user_id"`
	Type        string    `db:"type"`
	Description string    `db:"description"`
	Amount      float64   `db:"amount"`
	IsFixed     bool      `db:"is_fixed"`
	DayOfMonth  int64     `db:"day_of_month"`
	EndDate     string    `db:"end_date"`
	Category    string    `db:"category"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (RowTransaction) TableName() string {
	return "transactions"
}

type RowCategory struct {
	ID     int64  `db:"id"`
	UserID string `db:"user_id"`
	Name   string `db:"name"`
	Type   string `db:"type"`
	Color  string `db:"color"`
}

func (RowCategory) TableName() string {
	return "categories"
}
