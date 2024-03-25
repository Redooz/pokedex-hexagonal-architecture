package spi

import "github.com/redooz/podekex-hexagonal-architecture/domain/model"

type ITypePersistencePort interface {
	SaveTypeToDB(pokemonType *model.Type) error
	GetAllTypesFromDB() ([]*model.Type, error)
	GetTypeByIdFromDB(typeId int) (*model.Type, error)
	UpdateTypeFromDB(pokemonType *model.Type, typeId int) error
	DeleteTypeFromDB(typeId int) error
}
