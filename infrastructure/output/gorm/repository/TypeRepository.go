package repository

import (
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/entity"
)

type TypeRepository struct {
	Database *gorm.Database
}

func NewTypeRepository(database *gorm.Database) *TypeRepository {
	return &TypeRepository{Database: database}
}

func (t TypeRepository) Save(pokemonType *entity.Type) error {
	result := t.Database.DB.Create(pokemonType)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TypeRepository) GetAll() ([]*entity.Type, error) {
	var pokemonTypes []*entity.Type

	result := t.Database.DB.Find(&pokemonTypes)

	if result.Error != nil {
		return nil, result.Error
	}

	return pokemonTypes, nil
}

func (t TypeRepository) GetById(typeId int) (*entity.Type, error) {
	var pokemonType entity.Type

	result := t.Database.DB.Where(&entity.Type{ID: typeId}).First(&pokemonType)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pokemonType, nil
}

func (t TypeRepository) Update(pokemonType *entity.Type, typeId int) error {
	// Check if the type exists
	_, err := t.GetById(typeId)

	if err != nil {
		return err
	}

	result := t.Database.DB.Where(&entity.Type{ID: typeId}).Updates(pokemonType)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TypeRepository) Delete(typeId int) error {
	// Check if the type exists
	_, err := t.GetById(typeId)

	if err != nil {
		return err
	}

	result := t.Database.DB.Where(&entity.Type{ID: typeId}).Delete(&entity.Type{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
