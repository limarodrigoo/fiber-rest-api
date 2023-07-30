package repository

import "rest-api/models"

type TransactionRepository interface {
	SaveTransaction(transaction models.Transaction) error
	BulkSaveTransaction(transactions []models.Transaction) error
	UpdateTransaction(transaction *models.Transaction) error
	GetTransactionByHash(hash string) (transaction *models.Transaction, err error)
	GetTransactionByBlock(block int) (transaction []*models.Transaction, err error)
	GetAllTransactions() (transaction []*models.Transaction, err error)
	DeleteTransaction(id string) error
}
