package controllers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"rest-api/controllers"
	"rest-api/models"
	"rest-api/routes"
	"rest-api/tests/mocks"
	"rest-api/tests/testdata"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type ResponseData struct {
	Data models.Transaction `json:"data"`
}

func Test_transactionController(t *testing.T) {

	t.Run(
		"test_sucessfullreturn", func(t *testing.T) {
			control := gomock.NewController(t)
			defer control.Finish()

			hash := "e8fb9e06e8bcc16b695c5b2e8e39f9afafd80beb7dee0b1cb34caa375b1e3d0e"

			blocksRepo := mocks.NewMockBlockRepository(control)
			addressRepo := mocks.NewMockAddressRepository(control)
			transactionData := testdata.GetPointerTransactionData()
			transactionRepo := mocks.NewMockTransactionRepository(control)
			transactionController := controllers.NewTransactionsController(transactionRepo, blocksRepo, addressRepo)
			transactionRouter := routes.NewTransactionRoutes(transactionController)
			transactionRepo.EXPECT().GetTransactionByHash(hash).Return(transactionData, nil)

			app := fiber.New()
			transactionRouter.Install(app)

			route := fmt.Sprintf("/transaction/%s", hash)

			req := httptest.NewRequest("GET", route, nil)
			resp, _ := app.Test(req, 1)

			body, _ := ioutil.ReadAll(resp.Body)
			var jsonData ResponseData
			json.Unmarshal([]byte(string(body)), &jsonData)
			assert.Equalf(t, transactionData.BlockIndex, jsonData.Data.BlockIndex, "it_returns_the_correct_block_index")
			assert.Equalf(t, http.StatusOK, resp.StatusCode, "it_returns_expected_status_code")
		})

}
