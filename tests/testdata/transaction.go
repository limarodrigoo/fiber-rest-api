package testdata

import (
	"rest-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTransactionMongoData() models.Transaction {
	transactionTestData := models.Transaction{
		ID:          primitive.NewObjectID(),
		Hash:        "e8fb9e06e8bcc16b695c5b2e8e39f9afafd80beb7dee0b1cb34caa375b1e3d0e",
		Ver:         1,
		VinSz:       1,
		VoutSz:      2,
		Size:        207,
		Weight:      720,
		Fee:         0,
		RelayedBy:   "0.0.0.0",
		LockTime:    0,
		TxIndex:     500981282829641,
		DoubleSpend: false,
		Time:        1690518134,
		BlockIndex:  800566,
		BlockHeight: 800566,
		Inputs: []struct {
			Sequence int    "json:\"sequence\" bson:\"sequence\""
			Witness  string "json:\"witness\" bson:\"witness\""
			Script   string "json:\"script\" bson:\"script\""
			Index    int    "json:\"index\" bson:\"index\""
			PrevOut  struct {
				Addr              string "json:\"addr\" bson:\"addr\""
				N                 int    "json:\"n\" bson:\"n\""
				Script            string "json:\"script\" bson:\"script\""
				SpendingOutpoints []struct {
					TxIndex int "json:\"tx_index\" bson:\"tx_index\""
					N       int "json:\"n\" bson:\"n\""
				} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
				Spent   bool "json:\"spent\" bson:\"spent\""
				TxIndex int  "json:\"tx_index\" bson:\"tx_index\""
				Type    int  "json:\"type\" bson:\"type\""
				Value   int  "json:\"value\" bson:\"value\""
			} "json:\"prev_out\" bson:\"prev_out\""
		}{
			{
				Sequence: 4294967295,
				Witness:  "01200000000000000000000000000000000000000000000000000000000000000000",
				Script:   "0336370c0b2f4d61726120506f6f6c2f1a01000200000003000000181fa5b2050000008f010000",
				Index:    0,
				PrevOut: struct {
					Addr              string "json:\"addr\" bson:\"addr\""
					N                 int    "json:\"n\" bson:\"n\""
					Script            string "json:\"script\" bson:\"script\""
					SpendingOutpoints []struct {
						TxIndex int "json:\"tx_index\" bson:\"tx_index\""
						N       int "json:\"n\" bson:\"n\""
					} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
					Spent   bool "json:\"spent\" bson:\"spent\""
					TxIndex int  "json:\"tx_index\" bson:\"tx_index\""
					Type    int  "json:\"type\" bson:\"type\""
					Value   int  "json:\"value\" bson:\"value\""
				}{
					Addr:   "",
					N:      4294967295,
					Script: "",
					SpendingOutpoints: []struct {
						TxIndex int "json:\"tx_index\" bson:\"tx_index\""
						N       int "json:\"n\" bson:\"n\""
					}{
						{
							TxIndex: 500981282829641,
							N:       0,
						},
					},
					Spent:   true,
					TxIndex: 0,
					Type:    0,
					Value:   0,
				},
			},
		},
		Outs: []struct {
			Type              int         "json:\"type\" bson:\"type\""
			Spent             bool        "json:\"spent\" bson:\"spent\""
			Value             interface{} "json:\"value\" bson:\"value\""
			SpendingOutpoints []struct {
				TxIndex int "json:\"tx_index\" bson:\"tx_index\""
				N       int "json:\"n\" bson:\"n\""
			} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
			N       int    "json:\"n\" bson:\"n\""
			TxIndex int    "json:\"tx_index\" bson:\"tx_index\""
			Script  string "json:\"script\" bson:\"script\""
			Addr    string "json:\"addr\" bson:\"addr\""
		}{
			{
				Type:  0,
				Spent: false,
				SpendingOutpoints: []struct {
					TxIndex int "json:\"tx_index\" bson:\"tx_index\""
					N       int "json:\"n\" bson:\"n\""
				}{
					{TxIndex: 0, N: 0},
				},
				Value:   636721454,
				N:       0,
				TxIndex: 500981282829641,
				Script:  "76a9142fc701e2049ee4957b07134b6c1d771dd5a96b2188ac",
				Addr:    "15MdAHnkxt9TMC2Rj595hsg8Hnv693pPBB",
			},
		},
	}
	return transactionTestData
}

func GetTransactionData() models.Transaction {
	transactionData := models.Transaction{
		Hash:        "e8fb9e06e8bcc16b695c5b2e8e39f9afafd80beb7dee0b1cb34caa375b1e3d0e",
		Ver:         1,
		VinSz:       1,
		VoutSz:      2,
		Size:        207,
		Weight:      720,
		Fee:         0,
		RelayedBy:   "0.0.0.0",
		LockTime:    0,
		TxIndex:     500981282829641,
		DoubleSpend: false,
		Time:        1690518134,
		BlockIndex:  800566,
		BlockHeight: 800566,
		Inputs: []struct {
			Sequence int    "json:\"sequence\" bson:\"sequence\""
			Witness  string "json:\"witness\" bson:\"witness\""
			Script   string "json:\"script\" bson:\"script\""
			Index    int    "json:\"index\" bson:\"index\""
			PrevOut  struct {
				Addr              string "json:\"addr\" bson:\"addr\""
				N                 int    "json:\"n\" bson:\"n\""
				Script            string "json:\"script\" bson:\"script\""
				SpendingOutpoints []struct {
					TxIndex int "json:\"tx_index\" bson:\"tx_index\""
					N       int "json:\"n\" bson:\"n\""
				} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
				Spent   bool "json:\"spent\" bson:\"spent\""
				TxIndex int  "json:\"tx_index\" bson:\"tx_index\""
				Type    int  "json:\"type\" bson:\"type\""
				Value   int  "json:\"value\" bson:\"value\""
			} "json:\"prev_out\" bson:\"prev_out\""
		}{
			{
				Sequence: 4294967295,
				Witness:  "01200000000000000000000000000000000000000000000000000000000000000000",
				Script:   "0336370c0b2f4d61726120506f6f6c2f1a01000200000003000000181fa5b2050000008f010000",
				Index:    0,
				PrevOut: struct {
					Addr              string "json:\"addr\" bson:\"addr\""
					N                 int    "json:\"n\" bson:\"n\""
					Script            string "json:\"script\" bson:\"script\""
					SpendingOutpoints []struct {
						TxIndex int "json:\"tx_index\" bson:\"tx_index\""
						N       int "json:\"n\" bson:\"n\""
					} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
					Spent   bool "json:\"spent\" bson:\"spent\""
					TxIndex int  "json:\"tx_index\" bson:\"tx_index\""
					Type    int  "json:\"type\" bson:\"type\""
					Value   int  "json:\"value\" bson:\"value\""
				}{
					Addr:   "",
					N:      4294967295,
					Script: "",
					SpendingOutpoints: []struct {
						TxIndex int "json:\"tx_index\" bson:\"tx_index\""
						N       int "json:\"n\" bson:\"n\""
					}{
						{
							TxIndex: 500981282829641,
							N:       0,
						},
					},
					Spent:   true,
					TxIndex: 0,
					Type:    0,
					Value:   0,
				},
			},
		},
		Outs: []struct {
			Type              int         "json:\"type\" bson:\"type\""
			Spent             bool        "json:\"spent\" bson:\"spent\""
			Value             interface{} "json:\"value\" bson:\"value\""
			SpendingOutpoints []struct {
				TxIndex int "json:\"tx_index\" bson:\"tx_index\""
				N       int "json:\"n\" bson:\"n\""
			} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
			N       int    "json:\"n\" bson:\"n\""
			TxIndex int    "json:\"tx_index\" bson:\"tx_index\""
			Script  string "json:\"script\" bson:\"script\""
			Addr    string "json:\"addr\" bson:\"addr\""
		}{
			{
				Type:  0,
				Spent: false,
				SpendingOutpoints: []struct {
					TxIndex int "json:\"tx_index\" bson:\"tx_index\""
					N       int "json:\"n\" bson:\"n\""
				}{
					{TxIndex: 0, N: 0},
				},
				Value:   636721454,
				N:       0,
				TxIndex: 500981282829641,
				Script:  "76a9142fc701e2049ee4957b07134b6c1d771dd5a96b2188ac",
				Addr:    "15MdAHnkxt9TMC2Rj595hsg8Hnv693pPBB",
			},
		},
	}
	return transactionData
}

func GetPointerTransactionData() *models.Transaction {
	transactionData := &models.Transaction{
		Hash:        "e8fb9e06e8bcc16b695c5b2e8e39f9afafd80beb7dee0b1cb34caa375b1e3d0e",
		Ver:         1,
		VinSz:       1,
		VoutSz:      2,
		Size:        207,
		Weight:      720,
		Fee:         0,
		RelayedBy:   "0.0.0.0",
		LockTime:    0,
		TxIndex:     500981282829641,
		DoubleSpend: false,
		Time:        1690518134,
		BlockIndex:  800566,
		BlockHeight: 800566,
		Inputs: []struct {
			Sequence int    "json:\"sequence\" bson:\"sequence\""
			Witness  string "json:\"witness\" bson:\"witness\""
			Script   string "json:\"script\" bson:\"script\""
			Index    int    "json:\"index\" bson:\"index\""
			PrevOut  struct {
				Addr              string "json:\"addr\" bson:\"addr\""
				N                 int    "json:\"n\" bson:\"n\""
				Script            string "json:\"script\" bson:\"script\""
				SpendingOutpoints []struct {
					TxIndex int "json:\"tx_index\" bson:\"tx_index\""
					N       int "json:\"n\" bson:\"n\""
				} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
				Spent   bool "json:\"spent\" bson:\"spent\""
				TxIndex int  "json:\"tx_index\" bson:\"tx_index\""
				Type    int  "json:\"type\" bson:\"type\""
				Value   int  "json:\"value\" bson:\"value\""
			} "json:\"prev_out\" bson:\"prev_out\""
		}{
			{
				Sequence: 4294967295,
				Witness:  "01200000000000000000000000000000000000000000000000000000000000000000",
				Script:   "0336370c0b2f4d61726120506f6f6c2f1a01000200000003000000181fa5b2050000008f010000",
				Index:    0,
				PrevOut: struct {
					Addr              string "json:\"addr\" bson:\"addr\""
					N                 int    "json:\"n\" bson:\"n\""
					Script            string "json:\"script\" bson:\"script\""
					SpendingOutpoints []struct {
						TxIndex int "json:\"tx_index\" bson:\"tx_index\""
						N       int "json:\"n\" bson:\"n\""
					} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
					Spent   bool "json:\"spent\" bson:\"spent\""
					TxIndex int  "json:\"tx_index\" bson:\"tx_index\""
					Type    int  "json:\"type\" bson:\"type\""
					Value   int  "json:\"value\" bson:\"value\""
				}{
					Addr:   "",
					N:      4294967295,
					Script: "",
					SpendingOutpoints: []struct {
						TxIndex int "json:\"tx_index\" bson:\"tx_index\""
						N       int "json:\"n\" bson:\"n\""
					}{
						{
							TxIndex: 500981282829641,
							N:       0,
						},
					},
					Spent:   true,
					TxIndex: 0,
					Type:    0,
					Value:   0,
				},
			},
		},
		Outs: []struct {
			Type              int         "json:\"type\" bson:\"type\""
			Spent             bool        "json:\"spent\" bson:\"spent\""
			Value             interface{} "json:\"value\" bson:\"value\""
			SpendingOutpoints []struct {
				TxIndex int "json:\"tx_index\" bson:\"tx_index\""
				N       int "json:\"n\" bson:\"n\""
			} "json:\"spending_outpoints\" bson:\"spending_outpoints\""
			N       int    "json:\"n\" bson:\"n\""
			TxIndex int    "json:\"tx_index\" bson:\"tx_index\""
			Script  string "json:\"script\" bson:\"script\""
			Addr    string "json:\"addr\" bson:\"addr\""
		}{
			{
				Type:  0,
				Spent: false,
				SpendingOutpoints: []struct {
					TxIndex int "json:\"tx_index\" bson:\"tx_index\""
					N       int "json:\"n\" bson:\"n\""
				}{
					{TxIndex: 0, N: 0},
				},
				Value:   636721454,
				N:       0,
				TxIndex: 500981282829641,
				Script:  "76a9142fc701e2049ee4957b07134b6c1d771dd5a96b2188ac",
				Addr:    "15MdAHnkxt9TMC2Rj595hsg8Hnv693pPBB",
			},
		},
	}
	return transactionData
}
