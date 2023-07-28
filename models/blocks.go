package models

type Block struct {
	Hash       string        `bson:"hash" json:"hash"`
	Ver        int           `bson:"ver" json:"ver"`
	PrevBlock  string        `bson:"prev_block" json:"prev_block"`
	MrklRoot   string        `bson:"mrkl_root" json:"mrkl_root"`
	Time       int64         `bson:"time" json:"time"`
	Bits       int           `bson:"bits" json:"bits"`
	NextBlock  []string      `bson:"next_block" json:"next_block"`
	Fee        int           `bson:"fee" json:"fee"`
	Nonce      int           `bson:"nonce" json:"nonce"`
	Ntx        int           `bson:"n_tx" json:"n_tx"`
	Size       int           `bson:"size" json:"size"`
	BlockIndex int           `bson:"block_index" json:"block_index"`
	MainChain  bool          `bson:"main_chain" json:"main_chain"`
	Height     int           `bson:"height" json:"height"`
	Weight     int           `bson:"weight" json:"weight"`
	Tx         []Transaction `bson:"tx" json:"tx"`
}
