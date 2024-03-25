package usecase

import (
	"github.com/redooz/podekex-hexagonal-architecture/domain/api"
	"github.com/redooz/podekex-hexagonal-architecture/domain/constants"
	domainError "github.com/redooz/podekex-hexagonal-architecture/domain/error"
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
	"github.com/redooz/podekex-hexagonal-architecture/domain/spi"
)

type PokemonUseCase struct {
	pokemonPersistencePort spi.IPokemonPersistencePort
	typeServicePort        api.ITypeServicePort
}

func NewPokemonUseCase(pokemonPersistencePort spi.IPokemonPersistencePort, typeServicePort api.ITypeServicePort) *PokemonUseCase {
	return &PokemonUseCase{pokemonPersistencePort: pokemonPersistencePort, typeServicePort: typeServicePort}
}

func (p PokemonUseCase) SavePokemon(pokemon *model.Pokemon) error {
	err := p.pokemonPersistencePort.SavePokemonToDB(pokemon)
	if err != nil {
		return err
	}
	return nil
}

func (p PokemonUseCase) GetAllPokemons() ([]*model.Pokemon, error) {
	pokemons, err := p.pokemonPersistencePort.GetAllPokemonsFromDB()
	if err != nil {
		return nil, err
	}

	if len(pokemons) == 0 {
		return nil, &domainError.NoDataFound{Message: constants.PokemonNoDataFoundErrorMessage}
	}

	return pokemons, nil
}

func (p PokemonUseCase) GetPokemonByNumber(pokemonNumber int) (*model.Pokemon, error) {
	pokemon, err := p.pokemonPersistencePort.GetPokemonByNumberFromDB(pokemonNumber)
	if err != nil {
		return nil, err
	}

	if pokemon == nil {
		return nil, &domainError.NoDataFound{Message: constants.PokemonNoDataFoundErrorMessage}
	}

	return pokemon, nil
}

func (p PokemonUseCase) UpdatePokemon(pokemon *model.Pokemon, number int) error {
	// Check if the pokemon exists
	_, err := p.GetPokemonByNumber(number)

	if err != nil {
		return err
	}

	// Check if the type exists
	_, err = p.typeServicePort.GetTypeById(pokemon.PokemonType.ID)

	if err != nil {
		return err
	}

	err = p.pokemonPersistencePort.UpdatePokemonFromDB(pokemon, number)

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

	err = p.pokemonPersistencePort.DeletePokemonFromDB(pokemonNumber)

	if err != nil {
		return err
	}

	return nil
}
