package repository

import "rest-api/models"

type BlockRepository interface {
	Save(block models.Block) error
	Update(block *models.Block) error
	GetByHash(hash string) (block *models.Block, err error)
	GetAll() (block []*models.Block, err error)
	Delete(id string) error
}
