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

type PokedexHandler struct {
	pokemonServicePort api.IPokemonServicePort
	typeServicePort    api.ITypeServicePort
}

func NewPokedexHandler(pokemonServicePort api.IPokemonServicePort, typeServicePort api.ITypeServicePort) *PokedexHandler {
	return &PokedexHandler{pokemonServicePort: pokemonServicePort, typeServicePort: typeServicePort}
}

func (p PokedexHandler) SavePokemonToPokedex(pokedex *request.Pokedex) (httpStatus int, err error) {
	err = p.pokemonServicePort.SavePokemon(mapper.PokedexRequestToModel(pokedex))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (p PokedexHandler) GetAllPokemonFromPokedex() (response []*response.Pokedex, httpStatus int, err error) {
	pokedex, err := p.pokemonServicePort.GetAllPokemons()

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var noDataFound *domainError.NoDataFound
	if errors.As(err, &noDataFound) {
		return nil, http.StatusNotFound, nil
	}

	return mapper.SlicePokemonModelToSliceResponse(pokedex), http.StatusOK, nil
}

func (p PokedexHandler) GetPokemonFromPokedexByNumber(pokedexNumber int) (response *response.Pokedex, httpStatus int, err error) {
	pokemon, err := p.pokemonServicePort.GetPokemonByNumber(pokedexNumber)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var noDataFound *domainError.NoDataFound
	if errors.As(err, &noDataFound) {
		return nil, http.StatusNotFound, nil
	}

	return mapper.PokedexModelToResponse(pokemon), http.StatusOK, nil
}

func (p PokedexHandler) UpdatePokemonInPokedex(pokedex *request.Pokedex) (httpStatus int, err error) {
	err = p.pokemonServicePort.UpdatePokemon(mapper.PokedexRequestToModel(pokedex))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

func (p PokedexHandler) DeletePokemonFromPokedex(pokedexNumber int) (httpStatus int, err error) {
	err = p.pokemonServicePort.DeletePokemon(pokedexNumber)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
