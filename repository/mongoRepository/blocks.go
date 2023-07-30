package mongorepository

import (
	"context"
	"rest-api/helper"
	"rest-api/models"
	"rest-api/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const BlocksCollection = "blocks"

type blockRepository struct {
	c *mongo.Collection
}

func NewBlocksRepository(conn mongo.Database) repository.BlockRepository {
	return &blockRepository{conn.Collection(BlocksCollection)}
}

func (b *blockRepository) Delete(id string) error {
	var filter = bson.D{{Key: "hash", Value: id}}

	_, err := b.c.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (b *blockRepository) GetAll() (block []*models.Block, err error) {
	filter := bson.D{}

	cursor, err := b.c.Find(context.TODO(), filter)

	if err != nil {
		helper.ErrorPanic(err)
	}

	var blocks []*models.Block
	err = cursor.All(context.TODO(), &blocks)

	for cursor.Next(context.TODO()) {
		var block *models.Block
		cursor.Decode(&block)
		blocks = append(blocks, block)
	}

	return blocks, err
}

func (b *blockRepository) GetByHash(hash string) (block *models.Block, err error) {

	var filter = bson.D{{Key: "hash", Value: hash}}

	var blockFound *models.Block

	err = b.c.FindOne(context.TODO(), filter).Decode(&blockFound)

	return blockFound, err
}

func (b *blockRepository) Save(block models.Block) error {
	_, err := b.c.InsertOne(context.TODO(), block)

	if err != nil {
		helper.ErrorPanic(err)
	}

	return nil
}

func (b *blockRepository) Update(blockUpdate models.Block) error {
	filter := bson.D{{Key: "hash", Value: blockUpdate.Hash}}
	update := bson.M{
		"$set": bson.M{
			"ver":         blockUpdate.Ver,
			"prev_block":  blockUpdate.PrevBlock,
			"mrkl_root":   blockUpdate.MrklRoot,
			"time":        blockUpdate.Time,
			"bits":        blockUpdate.Bits,
			"next_block":  blockUpdate.NextBlock,
			"fee":         blockUpdate.Fee,
			"nonce":       blockUpdate.Nonce,
			"n_tx":        blockUpdate.Ntx,
			"size":        blockUpdate.Size,
			"block_index": blockUpdate.BlockIndex,
			"main_chain":  blockUpdate.MainChain,
			"height":      blockUpdate.Height,
			"weight":      blockUpdate.Weight,
		},
	}
	_, err := b.c.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return err
		}
		helper.ErrorPanic(err)
	}

	return nil
}
