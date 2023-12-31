package services

import (
	"log"
	"meffin-transactions-api/configs"
	"meffin-transactions-api/internal/repository"
)

type Services struct {
	TransactionService TransactionService
}

func InitServices(config configs.Config) *Services {
	transactionRepository, err := repository.NewTransactionRepository(repository.Config{
		DbUrl: config.DbUrl,
	})
	if err != nil {
		log.Printf("could not init transactionRepository: %v \n", err)
	}
	err = transactionRepository.AutoMigrate()
	if err != nil {
		log.Printf("could not auto migrate transactionRepository: %v \n", err)
	}

	return &Services{
		TransactionService: NewTransactionService(transactionRepository),
	}

}
