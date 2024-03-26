package adapter

import (
	"github.com/redooz/podekex-hexagonal-architecture/domain/model"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/entitymapper"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/repository"
)

type PokemonPersistenceAdapter struct {
	Repository *repository.PokemonRepository
}

func NewPokemonPersistenceAdapter(repository *repository.PokemonRepository) *PokemonPersistenceAdapter {
	return &PokemonPersistenceAdapter{Repository: repository}
}

func (p PokemonPersistenceAdapter) SavePokemonToDB(pokemon *model.Pokemon) error {
	err := p.Repository.Save(entitymapper.PokemonModelToEntity(pokemon))

	if err != nil {
		return err
	}

	return nil
}

func (p PokemonPersistenceAdapter) GetAllPokemonsFromDB() ([]*model.Pokemon, error) {
	pokemons, err := p.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return entitymapper.SlicePokemonEntityToSliceModel(pokemons), nil
}

func (p PokemonPersistenceAdapter) GetPokemonByNumberFromDB(pokemonNumber int) (*model.Pokemon, error) {
	pokemon, err := p.Repository.GetByNumber(pokemonNumber)

	if err != nil {
		return nil, err
	}

	return entitymapper.PokemonEntityToModel(pokemon), nil
}

func (p PokemonPersistenceAdapter) UpdatePokemonFromDB(pokemon *model.Pokemon, number int) error {
	err := p.Repository.Update(entitymapper.PokemonModelToEntity(pokemon), number)

	if err != nil {
		return err
	}

	return nil
}

func (p PokemonPersistenceAdapter) DeletePokemonFromDB(pokemonNumber int) error {
	err := p.Repository.Delete(pokemonNumber)

	if err != nil {
		return err
	}

	return nil
}
