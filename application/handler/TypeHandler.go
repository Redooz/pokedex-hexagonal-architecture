package handler

import (
	"errors"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/response"
	"github.com/redooz/podekex-hexagonal-architecture/application/mapper"
	"github.com/redooz/podekex-hexagonal-architecture/domain/api"
	domainError "github.com/redooz/podekex-hexagonal-architecture/domain/error"
	"net/http"
)

type TypeHandler struct {
	typeServicePort api.ITypeServicePort
}

func NewTypeHandler(typeServicePort api.ITypeServicePort) *TypeHandler {
	return &TypeHandler{typeServicePort: typeServicePort}
}

func (t TypeHandler) CreatePokemonType(pokemonType *request.Type) (httpStatus int, err error) {
	err = t.typeServicePort.SaveType(mapper.TypeRequestToModel(pokemonType))

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (t TypeHandler) GetAllPokemonTypes() (response []*response.Type, httpStatus int, err error) {
	types, err := t.typeServicePort.GetAllTypes()

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var noDataFound *domainError.NoDataFound
	if errors.As(err, &noDataFound) {
		return nil, http.StatusNotFound, nil
	}

	return mapper.SliceTypeModelToSliceResponse(types), http.StatusOK, nil
}

func (t TypeHandler) GetPokemonTypeByID(typeID int) (response *response.Type, httpStatus int, err error) {
	pokemonType, err := t.typeServicePort.GetTypeById(typeID)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var noDataFound *domainError.NoDataFound
	if errors.As(err, &noDataFound) {
		return nil, http.StatusNotFound, nil
	}

	return mapper.TypeModelToResponse(pokemonType), http.StatusOK, nil
}

func (t TypeHandler) UpdatePokemonType(pokemonType *request.Type) (httpStatus int, err error) {
	err = t.typeServicePort.UpdateType(mapper.TypeRequestToModel(pokemonType))

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

func (t TypeHandler) DeletePokemonType(typeID int) (httpStatus int, err error) {
	err = t.typeServicePort.DeleteType(typeID)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
