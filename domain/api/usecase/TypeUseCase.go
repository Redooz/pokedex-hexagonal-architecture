package usecase

import (
	"github.com/redooz/podekex-hexagonal-architecture/domain/constants"
	domainError "github.com/redooz/podekex-hexagonal-architecture/domain/error"
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
	"github.com/redooz/podekex-hexagonal-architecture/domain/spi"
)

type TypeUseCase struct {
	typePersincePort spi.ITypePersistencePort
}

func (t TypeUseCase) SaveType(pokemonType *model.Type) error {
	err := t.typePersincePort.SaveType(pokemonType)

	if err != nil {
		return err
	}

	return nil
}

func (t TypeUseCase) GetAllTypes() ([]*model.Type, error) {
	types, err := t.typePersincePort.GetAllTypes()

	if err != nil {
		return nil, err
	}

	if len(types) == 0 {
		return nil, &domainError.NoDataFound{Message: constants.TypeNoDataFoundErrorMessage}
	}

	return types, nil
}

func (t TypeUseCase) GetTypeById(typeId int) (*model.Type, error) {
	pokemonType, err := t.typePersincePort.GetTypeById(typeId)

	if err != nil {
		return nil, err
	}

	if pokemonType == nil {
		return nil, &domainError.NoDataFound{Message: constants.TypeNoDataFoundErrorMessage}
	}

	return pokemonType, nil
}

func (t TypeUseCase) UpdateType(pokemonType *model.Type) error {
	// Check if the type exists
	_, err := t.GetTypeById(pokemonType.ID)

	if err != nil {
		return err
	}

	err = t.typePersincePort.UpdateType(pokemonType)

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

	err = t.typePersincePort.DeleteType(typeId)

	if err != nil {
		return err
	}

	return nil
}
