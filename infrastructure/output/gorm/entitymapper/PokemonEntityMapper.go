package entitymapper

import (
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/entity"
)

func PokemonEntityToModel(entity *entity.Pokemon) *model.Pokemon {
	return &model.Pokemon{
		ID:     entity.ID,
		Number: entity.Number,
		Name:   entity.Name,
		PokemonType: model.Type{
			ID:         entity.Type.ID,
			FirstType:  entity.Type.FirstType,
			SecondType: entity.Type.SecondType,
		},
	}
}

func PokemonModelToEntity(model *model.Pokemon) *entity.Pokemon {
	return &entity.Pokemon{
		Number: model.Number,
		Name:   model.Name,
		Type: entity.Type{
			ID:         model.PokemonType.ID,
			FirstType:  model.PokemonType.FirstType,
			SecondType: model.PokemonType.SecondType,
		},
	}
}

func SlicePokemonEntityToSliceModel(entityPokemons []*entity.Pokemon) []*model.Pokemon {
	var modelPokemons []*model.Pokemon
	for _, entityPokemon := range entityPokemons {
		modelPokemons = append(modelPokemons, PokemonEntityToModel(entityPokemon))
	}
	return modelPokemons
}

func SlicePokemonModelToSliceEntity(modelPokemons []*model.Pokemon) []*entity.Pokemon {
	var entityPokemons []*entity.Pokemon
	for _, modelPokemon := range modelPokemons {
		entityPokemons = append(entityPokemons, PokemonModelToEntity(modelPokemon))
	}
	return entityPokemons
}
