package repository

import (
	"log"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DbUrl string
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(config Config) (IRepository, error) {
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
	return &Repository{db: database}, nil

}

func (r *Repository) AutoMigrate() error {
	return r.db.AutoMigrate(&RowTransaction{}, &RowCategory{})
}
