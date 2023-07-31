package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Hash        string             `json:"hash" bson:"hash"`
	Ver         int                `json:"ver" bson:"ver"`
	VinSz       int                `json:"vin_sz" bson:"vin_sz"`
	VoutSz      int                `json:"vout_sz" bson:"vout_sz"`
	Size        int                `json:"size" bson:"size"`
	Weight      int                `json:"weight" bson:"weight"`
	Fee         int                `json:"fee" bson:"fee"`
	RelayedBy   string             `json:"relayed_by" bson:"relayed_by"`
	LockTime    int                `json:"lock_time" bson:"lock_time"`
	TxIndex     int                `json:"tx_index" bson:"tx_index"`
	DoubleSpend bool               `json:"double_spend" bson:"double_spend"`
	Time        int64              `json:"time" bson:"time"`
	BlockIndex  int64              `json:"block_index" bson:"block_index"`
	BlockHeight int64              `json:"block_height" bson:"block_height"`
	Inputs      []struct {
		Sequence int    `json:"sequence" bson:"sequence"`
		Witness  string `json:"witness" bson:"witness"`
		Script   string `json:"script" bson:"script"`
		Index    int    `json:"index" bson:"index"`
		PrevOut  struct {
			Addr              string `json:"addr" bson:"addr"`
			N                 int    `json:"n" bson:"n"`
			Script            string `json:"script" bson:"script"`
			SpendingOutpoints []struct {
				TxIndex int `json:"tx_index" bson:"tx_index"`
				N       int `json:"n" bson:"n"`
			} `json:"spending_outpoints" bson:"spending_outpoints"`
			Spent   bool `json:"spent" bson:"spent"`
			TxIndex int  `json:"tx_index" bson:"tx_index"`
			Type    int  `json:"type" bson:"type"`
			Value   int  `json:"value" bson:"value"`
		} `json:"prev_out" bson:"prev_out"`
	} `json:"inputs" bson:"inputs"`
	Outs []struct {
		Type              int         `json:"type" bson:"type"`
		Spent             bool        `json:"spent" bson:"spent"`
		Value             interface{} `json:"value" bson:"value"`
		SpendingOutpoints []struct {
			TxIndex int `json:"tx_index" bson:"tx_index"`
			N       int `json:"n" bson:"n"`
		} `json:"spending_outpoints" bson:"spending_outpoints"`
		N       int    `json:"n" bson:"n"`
		TxIndex int    `json:"tx_index" bson:"tx_index"`
		Script  string `json:"script" bson:"script"`
		Addr    string `json:"addr" bson:"addr"`
	} `json:"out" bson:"out"`
}
