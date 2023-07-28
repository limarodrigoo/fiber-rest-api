package repository

import (
	"context"
	"rest-api/helper"
	"rest-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const AddressCollection = "addresses"

type AddressRepository interface {
	Save(address models.Address) error
	Update(address *models.Address) error
	GetByAddress(hash string) (address *models.Address, err error)
	GetAll() (addresses []*models.Address, err error)
	Delete(id string) error
}

type addressRepository struct {
	c *mongo.Collection
}

func NewAddressRepository(conn mongo.Database) AddressRepository {
	return &addressRepository{conn.Collection(AddressCollection)}
}

func (a *addressRepository) Delete(id string) error {
	var filter = bson.D{{Key: "address", Value: id}}

	_, err := a.c.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func (a *addressRepository) GetAll() (addresses []*models.Address, err error) {
	filter := bson.D{}

	cursor, err := a.c.Find(context.TODO(), filter)

	if err != nil {
		helper.ErrorPanic(err)
	}

	var addressesFound []*models.Address
	err = cursor.All(context.TODO(), &addressesFound)

	for cursor.Next(context.TODO()) {
		var address *models.Address
		cursor.Decode(&address)
		addressesFound = append(addressesFound, address)
	}

	return addressesFound, err
}

func (a *addressRepository) GetByAddress(addressHash string) (address *models.Address, err error) {
	var filter = bson.D{{Key: "address", Value: addressHash}}

	var addressFound *models.Address

	err = a.c.FindOne(context.TODO(), filter).Decode(&addressFound)

	return addressFound, err
}

func (a *addressRepository) Save(address models.Address) error {
	_, err := a.c.InsertOne(context.TODO(), address)

	if err != nil {
		helper.ErrorPanic(err)
	}

	return nil
}

func (a *addressRepository) Update(address *models.Address) error {
	panic("unimplemented")
}
