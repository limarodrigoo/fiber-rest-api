package repository

import "rest-api/models"

type AddressRepository interface {
	Save(address models.Address) error
	Update(address *models.Address) error
	GetByAddress(hash string) (address *models.Address, err error)
	GetAll() (addresses []*models.Address, err error)
	Delete(id string) error
}
