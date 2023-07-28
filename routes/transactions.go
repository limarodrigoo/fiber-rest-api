package routes

import (
	"rest-api/controllers"

	"github.com/gofiber/fiber/v2"
)

type transactionRoutes struct {
	transactionsController controllers.TransactionController
}

func NewTransactionRoutes(transactionController controllers.TransactionController) Routes {
	return &transactionRoutes{transactionController}
}

// Install implements Routes.
func (t *transactionRoutes) Install(app *fiber.App) {
	app.Get("/transactions", t.transactionsController.FindAllTransactions)
	app.Get("/transaction/:hash", t.transactionsController.FindTransactionByHash)
	app.Get("/transactions/by-block/:hash", t.transactionsController.FindTransactionByBlock)
	app.Get("/transactions/by-address/:address", t.transactionsController.FindTransactionByAddress)
	app.Delete("/transaction/:hash", t.transactionsController.DeleteTransaction)
}
