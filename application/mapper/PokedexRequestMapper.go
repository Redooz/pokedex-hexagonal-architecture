package mapper

import (
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/response"
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
)

func PokedexRequestToModel(request *request.Pokedex) *model.Pokemon {
	return &model.Pokemon{
		Number: request.Number,
		Name:   request.Name,
		PokemonType: model.Type{
			ID: request.TypeId,
		},
	}
}

func TypeModelToResponse(model *model.Type) response.Type {
	return response.Type{
		ID:         model.ID,
		FirstType:  model.FirstType,
		SecondType: model.SecondType,
	}
}

func PokedexModelToResponse(model *model.Pokemon) *response.Pokedex {
	return &response.Pokedex{
		Number: model.Number,
		Name:   model.Name,
		Type:   TypeModelToResponse(&model.PokemonType),
	}
}

func SlicePokemonModelToSliceResponse(pokemons []*model.Pokemon) []*response.Pokedex {
	var pokedexResponse []*response.Pokedex
	for _, pokemon := range pokemons {
		pokedexResponse = append(pokedexResponse, PokedexModelToResponse(pokemon))
	}
	return pokedexResponse
}
