package repository

import (
	"context"
	"rest-api/helper"
	"rest-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const TransactionCollection = "transactions"

type TransactionRepository interface {
	SaveTransaction(transaction models.Transaction) error
	BulkSaveTransaction(transactions []models.Transaction) error
	UpdateTransaction(transaction *models.Transaction) error
	GetTransactionByHash(hash string) (transaction *models.Transaction, err error)
	GetTransactionByBlock(block int) (transaction []*models.Transaction, err error)
	GetAllTransactions() (transaction []*models.Transaction, err error)
	DeleteTransaction(id string) error
}

type transactionRepository struct {
	c *mongo.Collection
}

func NewTransactionsRepository(conn mongo.Database) TransactionRepository {
	return &transactionRepository{conn.Collection(TransactionCollection)}
}

func (t *transactionRepository) DeleteTransaction(id string) error {
	var filter = bson.D{{Key: "hash", Value: id}}

	_, err := t.c.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (t *transactionRepository) GetAllTransactions() (transaction []*models.Transaction, err error) {
	filter := bson.D{}

	cursor, err := t.c.Find(context.TODO(), filter)

	if err != nil {
		helper.ErrorPanic(err)
	}

	var transactions []*models.Transaction
	err = cursor.All(context.TODO(), &transactions)

	for cursor.Next(context.TODO()) {
		var transaction *models.Transaction
		cursor.Decode(&transaction)
		transactions = append(transactions, transaction)
	}

	return transactions, err
}

func (t *transactionRepository) GetTransactionByHash(hash string) (transaction *models.Transaction, err error) {

	var filter = bson.D{{Key: "hash", Value: hash}}

	var transactionFound *models.Transaction

	err = t.c.FindOne(context.TODO(), filter).Decode(&transactionFound)

	return transactionFound, err
}

func (t *transactionRepository) GetTransactionByBlock(blockIndex int) (transaction []*models.Transaction, err error) {

	var filter = bson.D{{Key: "block_index", Value: blockIndex}}

	cursor, err := t.c.Find(context.TODO(), filter)

	if err != nil {
		helper.ErrorPanic(err)
	}

	var transactions []*models.Transaction
	err = cursor.All(context.TODO(), &transactions)

	for cursor.Next(context.TODO()) {
		var transaction *models.Transaction
		cursor.Decode(&transaction)
		transactions = append(transactions, transaction)
	}

	return transactions, err
}

func (t *transactionRepository) SaveTransaction(transaction models.Transaction) error {

	t.c.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "hash", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	_, err := t.c.InsertOne(context.TODO(), transaction)

	if err != nil {
		helper.ErrorPanic(err)
	}

	return nil
}

func (t *transactionRepository) BulkSaveTransaction(transactions []models.Transaction) error {
	t.c.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "hash", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	opts := options.BulkWrite().SetOrdered(false)

	models := []mongo.WriteModel{}

	for _, transaction := range transactions {
		models = append(models, mongo.NewInsertOneModel().SetDocument(transaction))
	}

	t.c.BulkWrite(context.Background(), models, opts)

	return nil
}

func (t *transactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	panic("unimplemented")
}
