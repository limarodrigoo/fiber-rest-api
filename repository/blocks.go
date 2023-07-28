package repository

import (
	"context"
	"rest-api/helper"
	"rest-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const BlocksCollection = "blocks"

type BlockRepository interface {
	Save(block models.Block) error
	Update(block *models.Block) error
	GetByHash(hash string) (block *models.Block, err error)
	GetAll() (block []*models.Block, err error)
	Delete(id string) error
}

type blockRepository struct {
	c *mongo.Collection
}

func NewBlocksRepository(conn mongo.Database) BlockRepository {
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

func (b *blockRepository) Update(block *models.Block) error {
	panic("unimplemented")
}
