package spi

import "github.com/redooz/podekex-hexagonal-architecture/domain/model"

type IPokemonPersistencePort interface {
	SavePokemonToDB(pokemon *model.Pokemon) error
	GetAllPokemonsFromDB() ([]*model.Pokemon, error)
	GetPokemonByNumberFromDB(pokemonNumber int) (*model.Pokemon, error)
	UpdatePokemonFromDB(pokemon *model.Pokemon, number int) error
	DeletePokemonFromDB(pokemonNumber int) error
}
