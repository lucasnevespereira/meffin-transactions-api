package services

import (
	"log"
	"meffin-transactions-api/configs"
	"meffin-transactions-api/internal/repository"
)

type Services struct {
	TransactionService TransactionService
	CategoryService    ICategoryService
}

func InitServices(config configs.Config) *Services {
	repo, err := repository.NewRepository(repository.Config{
		DbUrl: config.DbUrl,
	})
	if err != nil {
		log.Printf("could not init transactionRepository: %v \n", err)
	}
	err = repo.AutoMigrate()
	if err != nil {
		log.Printf("could not auto migrate transactionRepository: %v \n", err)
	}

	return &Services{
		CategoryService:    NewCategoryService(repo),
		TransactionService: NewTransactionService(repo),
	}

}
