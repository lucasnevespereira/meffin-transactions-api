package repository

import (
	"context"
	"fmt"
	"time"
)

func (r *Repository) Create(ctx context.Context, rowTransaction *RowTransaction) (*RowTransaction, error) {
	result := r.db.WithContext(ctx).Create(rowTransaction)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", result.Error)
	}

	return rowTransaction, nil
}

func (r *Repository) GetTransactionsByUserID(ctx context.Context, userId string) ([]*RowTransaction, error) {
	var rowTransactions []*RowTransaction
	result := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&rowTransactions)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get transactions: %v", result.Error)
	}

	return rowTransactions, nil
}

func (r *Repository) DeleteTransaction(ctx context.Context, transactionID int64) error {
	result := r.db.WithContext(ctx).Delete(&RowTransaction{}, transactionID)
	if result.Error != nil {
		return fmt.Errorf("failed to delete transaction: %v", result.Error)
	}
	return nil
}

func (r *Repository) UpdateTransaction(ctx context.Context, updatedTransaction *RowTransaction) (*RowTransaction, error) {
	result := r.db.WithContext(ctx).Model(updatedTransaction).Updates(updatedTransaction)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update transaction: %v", result.Error)
	}
	return updatedTransaction, nil
}

func (r *Repository) DeleteExpiredTransactions(ctx context.Context) error {
	currentDateString := time.Now().Format("2006-01-02")

	result := r.db.WithContext(ctx).
		Where("end_date <= ? AND end_date <> ''", currentDateString).
		Delete(&RowTransaction{})

	if result.Error != nil {
		return fmt.Errorf("failed to delete expired transactions: %v", result.Error)
	}

	return nil
}
