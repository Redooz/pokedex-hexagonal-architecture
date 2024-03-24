package handler

import (
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/response"
	"github.com/redooz/podekex-hexagonal-architecture/application/mapper"
	"github.com/redooz/podekex-hexagonal-architecture/domain/api"
)

type PokedexHandler struct {
	pokemonServicePort api.IPokemonServicePort
	typeServicePort    api.ITypeServicePort
}

func (p PokedexHandler) SavePokemonToPokedex(pokedex *request.Pokedex) error {
	err := p.pokemonServicePort.SavePokemon(mapper.PokedexRequestToModel(pokedex))
	if err != nil {
		return err
	}

	return nil
}

func (p PokedexHandler) GetAllPokemonFromPokedex() ([]*response.Pokedex, error) {
	pokedex, err := p.pokemonServicePort.GetAllPokemons()

	if err != nil {
		return nil, err
	}

	return mapper.SlicePokemonModelToSliceResponse(pokedex), nil
}

func (p PokedexHandler) GetPokemonFromPokedexByNumber(pokedexNumber int) (*response.Pokedex, error) {
	//TODO implement me
	panic("implement me")
}

func (p PokedexHandler) UpdatePokemonInPokedex(pokedex *request.Pokedex) error {
	//TODO implement me
	panic("implement me")
}

func (p PokedexHandler) DeletePokemonFromPokedex(pokedexNumber int) error {
	//TODO implement me
	panic("implement me")
}
