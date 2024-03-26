package usecase

import (
	"github.com/redooz/podekex-hexagonal-architecture/domain/constants"
	domainError "github.com/redooz/podekex-hexagonal-architecture/domain/error"
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
	"github.com/redooz/podekex-hexagonal-architecture/domain/spi"
)

type TypeUseCase struct {
	typePersistencePort spi.ITypePersistencePort
}

func NewTypeUseCase(typePersincePort spi.ITypePersistencePort) *TypeUseCase {
	return &TypeUseCase{typePersistencePort: typePersincePort}
}

func (t TypeUseCase) SaveType(pokemonType *model.Type) error {
	err := t.typePersistencePort.SaveTypeToDB(pokemonType)

	if err != nil {
		return err
	}

	return nil
}

func (t TypeUseCase) GetAllTypes() ([]*model.Type, error) {
	types, err := t.typePersistencePort.GetAllTypesFromDB()

	if err != nil {
		return nil, err
	}

	if len(types) == 0 {
		return nil, &domainError.NoDataFound{Message: constants.TypeNoDataFoundErrorMessage}
	}

	return types, nil
}

func (t TypeUseCase) GetTypeById(typeId int) (*model.Type, error) {
	pokemonType, err := t.typePersistencePort.GetTypeByIdFromDB(typeId)

	if err != nil {
		return nil, err
	}

	if pokemonType == nil {
		return nil, &domainError.NoDataFound{Message: constants.TypeNoDataFoundErrorMessage}
	}

	return pokemonType, nil
}

func (t TypeUseCase) UpdateType(pokemonType *model.Type, typeId int) error {
	// Check if the type exists
	_, err := t.GetTypeById(typeId)

	if err != nil {
		return err
	}

	err = t.typePersistencePort.UpdateTypeFromDB(pokemonType, typeId)

	if err != nil {
		return err
	}

	return nil
}

func (t TypeUseCase) DeleteType(typeId int) error {
	// Check if the type exists
	_, err := t.GetTypeById(typeId)

	if err != nil {
		return err
	}

	err = t.typePersistencePort.DeleteTypeFromDB(typeId)

	if err != nil {
		return err
	}

	return nil
}
