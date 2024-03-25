package api

import "github.com/redooz/podekex-hexagonal-architecture/domain/model"

type ITypeServicePort interface {
	SaveType(pokemonType *model.Type) error
	GetAllTypes() ([]*model.Type, error)
	GetTypeById(typeId int) (*model.Type, error)
	UpdateType(pokemonType *model.Type, typeId int) error
	DeleteType(typeId int) error
}
