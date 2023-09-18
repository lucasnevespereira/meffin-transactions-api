package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	DbUrl string
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// Enforces implementation of interface at compile time
var _ TransactionRepository = (*TransactionRepositoryImpl)(nil)

func NewTransactionRepository(config Config) (*TransactionRepositoryImpl, error) {
	database, err := gorm.Open(postgres.Open(config.DbUrl))
	if err != nil {
		return nil, errors.Wrapf(err, "could not create postgres client")
	}

	internalDB, errInternalDB := database.DB()
	if errInternalDB != nil {
		return nil, errors.Wrapf(errInternalDB, "could not get internal db")
	}

	if errPing := internalDB.Ping(); errPing != nil {
		return nil, errors.Wrapf(errPing, "could not ping database")
	}

	log.Println("Transaction Repository started")
	return &TransactionRepositoryImpl{db: database}, nil

}

func (r *TransactionRepositoryImpl) AutoMigrate() error {
	return r.db.AutoMigrate(&RowTransaction{})
}

func (r *TransactionRepositoryImpl) Create(ctx context.Context, rowTransaction *RowTransaction) (*RowTransaction, error) {
	result := r.db.WithContext(ctx).Create(rowTransaction)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", result.Error)
	}

	return rowTransaction, nil
}

func (r *TransactionRepositoryImpl) GetTransactionsByUserID(ctx context.Context, userId string) ([]*RowTransaction, error) {
	var rowTransactions []*RowTransaction
	result := r.db.WithContext(ctx).Find(&rowTransactions).Where("userId = ?", userId)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get transactions: %v", result.Error)
	}

	return rowTransactions, nil
}

func (r *TransactionRepositoryImpl) DeleteTransaction(ctx context.Context, transactionID int64) error {
	result := r.db.WithContext(ctx).Delete(&RowTransaction{}, transactionID)
	if result.Error != nil {
		return fmt.Errorf("failed to delete transaction: %v", result.Error)
	}
	return nil
}

func (r *TransactionRepositoryImpl) UpdateTransaction(ctx context.Context, updatedTransaction *RowTransaction) (*RowTransaction, error) {
	result := r.db.WithContext(ctx).Model(updatedTransaction).Updates(updatedTransaction)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update transaction: %v", result.Error)
	}
	return updatedTransaction, nil
}
