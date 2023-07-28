package main

import (
	"fmt"
	"log"
	"rest-api/config"
	"rest-api/controllers"
	"rest-api/db"
	"rest-api/repository"
	"rest-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Println("Running API ...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load enviroment variables", err)
	}

	db.ConnectDatabase(&loadConfig)

	DB := *db.MI.DB

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	transactionsRepo := repository.NewTransactionsRepository(DB)
	blocksRepo := repository.NewBlocksRepository(DB)
	addressRepo := repository.NewAddressRepository(DB)

	transactionsController := controllers.NewTransactionsController(transactionsRepo, blocksRepo, addressRepo)
	blocksController := controllers.NewBlocksController(blocksRepo, transactionsRepo)

	transactionsRoutes := routes.NewTransactionRoutes(transactionsController)
	blocksRoutes := routes.NewBlocksRoutes(blocksController)

	transactionsRoutes.Install(app)
	blocksRoutes.Install(app)

	log.Fatal(app.Listen(":3000"))

}
