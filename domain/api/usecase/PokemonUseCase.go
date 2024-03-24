package usecase

import (
	"github.com/redooz/podekex-hexagonal-architecture/domain/constants"
	domainError "github.com/redooz/podekex-hexagonal-architecture/domain/error"
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
	"github.com/redooz/podekex-hexagonal-architecture/domain/spi"
)

type PokemonUseCase struct {
	pokemonPersistencePort spi.IPokemonPersistencePort
}

func NewPokemonUseCase(pokemonPersistencePort spi.IPokemonPersistencePort) *PokemonUseCase {
	return &PokemonUseCase{pokemonPersistencePort: pokemonPersistencePort}
}

func (p PokemonUseCase) SavePokemon(pokemon *model.Pokemon) error {
	err := p.pokemonPersistencePort.SavePokemon(pokemon)
	if err != nil {
		return err
	}
	return nil
}

func (p PokemonUseCase) GetAllPokemons() ([]*model.Pokemon, error) {
	pokemons, err := p.pokemonPersistencePort.GetAllPokemons()
	if err != nil {
		return nil, err
	}

	if len(pokemons) == 0 {
		return nil, &domainError.NoDataFound{Message: constants.PokemonNoDataFoundErrorMessage}
	}

	return pokemons, nil
}

func (p PokemonUseCase) GetPokemonByNumber(pokemonNumber int) (*model.Pokemon, error) {
	pokemon, err := p.pokemonPersistencePort.GetPokemonByNumber(pokemonNumber)
	if err != nil {
		return nil, err
	}

	if pokemon == nil {
		return nil, &domainError.NoDataFound{Message: constants.PokemonNoDataFoundErrorMessage}
	}

	return pokemon, nil
}

func (p PokemonUseCase) UpdatePokemon(pokemon *model.Pokemon) error {
	// Check if the pokemon exists
	_, err := p.GetPokemonByNumber(pokemon.Number)

	if err != nil {
		return err
	}

	err = p.pokemonPersistencePort.UpdatePokemon(pokemon)

	if err != nil {
		return err
	}

	return nil
}

func (p PokemonUseCase) DeletePokemon(pokemonNumber int) error {
	// Check if the pokemon exists
	_, err := p.GetPokemonByNumber(pokemonNumber)

	if err != nil {
		return err
	}

	err = p.pokemonPersistencePort.DeletePokemon(pokemonNumber)

	if err != nil {
		return err
	}

	return nil
}
