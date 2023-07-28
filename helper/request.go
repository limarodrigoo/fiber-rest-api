package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rest-api/models"
)

func GetTransactionJSON(hash string) (*models.Transaction, error) {
	const baseUrl = "https://blockchain.info/rawtx"
	url := fmt.Sprintf("%s/%s", baseUrl, hash)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data models.Transaction
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func GetBlockJSON(hash string) (*models.Block, error) {
	const baseUrl = "https://blockchain.info/rawblock"
	url := fmt.Sprintf("%s/%s", baseUrl, hash)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data models.Block
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func GetAddressJSON(hash string) (*models.Address, error) {
	const baseUrl = "https://blockchain.info/rawaddr"
	url := fmt.Sprintf("%s/%s", baseUrl, hash)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data models.Address
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
