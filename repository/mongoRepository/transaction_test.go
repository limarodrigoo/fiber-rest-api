package mongorepository

import (
	"fmt"
	"os"
	"rest-api/models"
	"rest-api/tests/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestTransactionFindByHashRepository(t *testing.T) {
	databaseName := "cryptos_test"
	err := os.Setenv("DB_NAME", databaseName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestdb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestdb.Close()

	mtestdb.Run("when_sending_a_valid_hash_returns_success", func(mt *mtest.T) {

		transactionData := testdata.GetTransactionMongoData()
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, "transactions"),
			mtest.FirstBatch,
			convertEntityToBson(transactionData),
		))
		databaseMock := mt.Client.Database(databaseName)

		repo := NewTransactionsRepository(*databaseMock)

		transaction, err := repo.GetTransactionByHash(transactionData.Hash)

		assert.Nil(t, err)
		assert.EqualValues(t, transactionData.Hash, transaction.Hash)
		assert.EqualValues(t, transactionData.TxIndex, transaction.TxIndex)
		assert.EqualValues(t, transactionData.BlockIndex, transaction.BlockIndex)
	})

	mtestdb.Run("when_returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, "transactions"),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(databaseName)

		repo := NewTransactionsRepository(*databaseMock)

		transaction, err := repo.GetTransactionByHash("test")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "mongo: no documents in result")
		assert.Nil(t, transaction)
	})

	mtestdb.Run("when_mongo_db_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "error", Value: 0},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewTransactionsRepository(*databaseMock)

		transaction, err := repo.GetTransactionByHash("test")

		assert.Nil(t, transaction)
		assert.NotNil(t, err)
	})
}

func TestDeleteTransactionRepository(t *testing.T) {
	databaseName := "cryptos_test"
	err := os.Setenv("DB_NAME", databaseName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestdb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestdb.Close()

	mtestdb.Run("when_delete_is_successfull", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewTransactionsRepository(*databaseMock)

		err := repo.DeleteTransaction("test")

		assert.Nil(t, err)
	})

	mtestdb.Run("when_delete_is_not_successfull", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: false},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewTransactionsRepository(*databaseMock)

		err := repo.DeleteTransaction("test")

		assert.NotNil(t, err)
	})
}

func convertEntityToBson(transaction models.Transaction) bson.D {
	return bson.D{
		{Key: "_id", Value: transaction.ID},
		{Key: "hash", Value: transaction.Hash},
		{Key: "tx_index", Value: transaction.TxIndex},
		{Key: "block_index", Value: transaction.BlockIndex},
	}
}
