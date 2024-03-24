package handler

import (
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/response"
)

type ITypeHandler interface {
	CreatePokemonType(pokemonType *request.Type) (httpStatus int, err error)
	GetAllPokemonTypes() (response []*response.Type, httpStatus int, err error)
	GetPokemonTypeByID(typeID int) (response *response.Type, httpStatus int, err error)
	UpdatePokemonType(pokemonType *request.Type) (httpStatus int, err error)
	DeletePokemonType(typeID int) (httpStatus int, err error)
}
