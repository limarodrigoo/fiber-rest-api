package controllers

import (
	"fmt"
	"rest-api/helper"
	"rest-api/models"
	"rest-api/repository"

	"github.com/gofiber/fiber/v2"
)

type TransactionController interface {
	FindAllTransactions(ctx *fiber.Ctx) error
	FindTransactionByHash(ctx *fiber.Ctx) error
	FindTransactionByBlock(ctx *fiber.Ctx) error
	FindTransactionByAddress(ctx *fiber.Ctx) error
	DeleteTransaction(ctx *fiber.Ctx) error
	AddTransaction(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type transactionController struct {
	transactionRepo repository.TransactionRepository
	blocksRepo      repository.BlockRepository
	addresRepo      repository.AddressRepository
}

func NewTransactionsController(transactionRepo repository.TransactionRepository, blocksRepo repository.BlockRepository, addresRepo repository.AddressRepository) TransactionController {
	return &transactionController{transactionRepo, blocksRepo, addresRepo}
}

func (t *transactionController) DeleteTransaction(ctx *fiber.Ctx) error {
	hash := ctx.Params("hash")
	err := t.transactionRepo.DeleteTransaction(hash)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"deleted": true,
	})
}

func (t *transactionController) FindAllTransactions(ctx *fiber.Ctx) error {
	result, err := t.transactionRepo.GetAllTransactions()

	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}

func (t *transactionController) FindTransactionByHash(ctx *fiber.Ctx) error {
	var err error
	var result *models.Transaction
	hash := ctx.Params("hash")
	result, err = t.transactionRepo.GetTransactionByHash(hash)

	if err != nil {
		result, err := helper.GetTransactionJSON(hash)
		if err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t.transactionRepo.SaveTransaction(*result)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": result,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}

func (t *transactionController) FindTransactionByBlock(ctx *fiber.Ctx) error {
	blockHash := ctx.Params("hash")
	block, err := t.blocksRepo.GetByHash(blockHash)

	if err != nil {
		result, err := helper.GetBlockJSON(blockHash)
		if err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t.blocksRepo.Save(*result)
		t.transactionRepo.BulkSaveTransaction(result.Tx)
		return ctx.JSON(fiber.Map{
			"data": result.Tx,
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": block.Tx,
		})
	}
}

func (t *transactionController) FindTransactionByAddress(ctx *fiber.Ctx) error {
	addressHash := ctx.Params("address")
	address, err := t.addresRepo.GetByAddress(addressHash)

	if err != nil {
		result, err := helper.GetAddressJSON(addressHash)

		if err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		t.addresRepo.Save(*result)
		t.transactionRepo.BulkSaveTransaction(result.Tx)

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": result.Tx,
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": address.Tx,
		})
	}
}

func (t *transactionController) AddTransaction(ctx *fiber.Ctx) error {
	var newTransaction models.Transaction
	err := ctx.BodyParser(&newTransaction)
	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	err = t.transactionRepo.SaveTransaction(newTransaction)
	if err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"created": true,
	})
}

func (t *transactionController) Update(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
