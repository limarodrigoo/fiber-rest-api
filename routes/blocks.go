package routes

import (
	"rest-api/controllers"

	"github.com/gofiber/fiber/v2"
)

type blocksRoutes struct {
	blocksController controllers.BlockController
}

func NewBlocksRoutes(blocksController controllers.BlockController) Routes {
	return &blocksRoutes{blocksController}
}

// Install implements Routes.
func (b *blocksRoutes) Install(app *fiber.App) {
	app.Get("/blocks", b.blocksController.FindAllBlocks)
	app.Get("/block/:hash", b.blocksController.FindBlocksByHash)
	app.Delete("/block/:hash", b.blocksController.DeleteBlock)
}
