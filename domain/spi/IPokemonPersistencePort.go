package spi

import "github.com/redooz/podekex-hexagonal-architecture/domain/model"

type IPokemonPersistencePort interface {
	SavePokemon(pokemon *model.Pokemon) error
	GetAllPokemons() ([]*model.Pokemon, error)
	GetPokemonByNumber(pokemonNumber int) (*model.Pokemon, error)
	UpdatePokemon(pokemon *model.Pokemon) error
	DeletePokemon(pokemonNumber int) error
}
