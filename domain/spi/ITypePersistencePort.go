package spi

import "github.com/redooz/podekex-hexagonal-architecture/domain/model"

type ITypePersistencePort interface {
	SaveType(pokemonType *model.Type) error
	GetAllTypes() ([]*model.Type, error)
	GetTypeById(typeId int) (*model.Type, error)
	UpdateType(pokemonType *model.Type) error
	DeleteType(typeId int) error
}
