package models

type Address struct {
	Hash          string        `bson:"hash160" json:"hash160"`
	Address       string        `bson:"address" json:"address"`
	TotalReceived int           `bson:"total_received" json:"total_receibed"`
	NTx           int           `bson:"n_tx" json:"n_tx"`
	NUnredeemed   int           `bson:"n_unredeemed" json:"n_unredeemed"`
	TotalSent     int           `bson:"total_sent" json:"total_sent"`
	FinalBalance  int           `bson:"final_balance" json:"final_balance"`
	Tx            []Transaction `bson:"txs" json:"txs"`
}
