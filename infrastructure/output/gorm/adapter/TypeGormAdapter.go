package adapter

import (
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/entitymapper"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/repository"
)

type TypePersistenceAdapter struct {
	Repository repository.Type
}

func NewTypePersistenceAdapter(repository repository.Type) *TypePersistenceAdapter {
	return &TypePersistenceAdapter{Repository: repository}
}

func (t TypePersistenceAdapter) SaveTypeToDB(pokemonType *model.Type) error {
	err := t.Repository.Save(entitymapper.TypeModelToEntity(pokemonType))

	if err != nil {
		return err
	}

	return nil
}

func (t TypePersistenceAdapter) GetAllTypesFromDB() ([]*model.Type, error) {
	types, err := t.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return entitymapper.SliceTypeEntityToSliceModel(types), nil
}

func (t TypePersistenceAdapter) GetTypeByIdFromDB(typeId int) (*model.Type, error) {
	pokemonType, err := t.Repository.GetById(typeId)

	if err != nil {
		return nil, err
	}

	return entitymapper.TypeEntityToModel(pokemonType), nil
}

func (t TypePersistenceAdapter) UpdateTypeFromDB(pokemonType *model.Type, typeId int) error {
	// Check if the type exists
	_, err := t.GetTypeByIdFromDB(typeId)

	if err != nil {
		return err
	}

	err = t.Repository.Update(entitymapper.TypeModelToEntity(pokemonType), typeId)

	if err != nil {
		return err
	}

	return nil
}

func (t TypePersistenceAdapter) DeleteTypeFromDB(typeId int) error {
	err := t.Repository.Delete(typeId)

	if err != nil {
		return err
	}

	return nil
}
