package controllers

import (
	"fmt"
	"rest-api/helper"
	"rest-api/models"
	"rest-api/repository"

	"github.com/gofiber/fiber/v2"
)

type BlockController interface {
	FindAllBlocks(ctx *fiber.Ctx) error
	FindBlocksByHash(ctx *fiber.Ctx) error
	DeleteBlock(ctx *fiber.Ctx) error
	UpdateBlock(ctx *fiber.Ctx) error
}

type blockController struct {
	blockRepo       repository.BlockRepository
	transactionRepo repository.TransactionRepository
}

func NewBlocksController(blockRepo repository.BlockRepository, transactionRepo repository.TransactionRepository) BlockController {
	return &blockController{blockRepo, transactionRepo}
}

func (b *blockController) DeleteBlock(ctx *fiber.Ctx) error {
	hash := ctx.Params("hash")
	err := b.blockRepo.Delete(hash)
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

func (b *blockController) FindAllBlocks(ctx *fiber.Ctx) error {
	result, err := b.blockRepo.GetAll()

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

func (b *blockController) FindBlocksByHash(ctx *fiber.Ctx) error {
	var err error
	var result *models.Block
	hash := ctx.Params("hash")
	result, err = b.blockRepo.GetByHash(hash)

	if err != nil {
		result, err := helper.GetBlockJSON(hash)
		if err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		b.blockRepo.Save(*result)
		b.transactionRepo.BulkSaveTransaction(result.Tx)

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": result,
		})
	}

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

func (b *blockController) UpdateBlock(ctx *fiber.Ctx) error {
	var body models.Block
	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	err = b.blockRepo.Update(body)
	if err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"updated": true,
	})
}
