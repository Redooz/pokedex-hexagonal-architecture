package repository

import (
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure/output/gorm/entity"
)

type Pokemon struct {
	Database *gorm.Database
}

func (p Pokemon) Save(pokemon *entity.Pokemon) error {
	result := p.Database.DB.Create(pokemon)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p Pokemon) GetAll() ([]*entity.Pokemon, error) {
	var pokemons []*entity.Pokemon

	result := p.Database.DB.Find(&pokemons)

	if result.Error != nil {
		return nil, result.Error
	}

	return pokemons, nil
}

func (p Pokemon) GetByNumber(pokemonNumber int) (*entity.Pokemon, error) {
	var pokemon entity.Pokemon

	result := p.Database.DB.Where(&entity.Pokemon{Number: pokemonNumber}).First(&pokemon)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pokemon, nil
}

func (p Pokemon) Update(pokemon *entity.Pokemon, number int) error {
	// Check if the pokemon exists
	_, err := p.GetByNumber(number)

	if err != nil {
		return err
	}

	result := p.Database.DB.Where(&entity.Pokemon{Number: number}).Updates(pokemon)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p Pokemon) Delete(pokemonNumber int) error {
	// Check if the pokemon exists
	_, err := p.GetByNumber(pokemonNumber)

	if err != nil {
		return err
	}

	result := p.Database.DB.Where(&entity.Pokemon{Number: pokemonNumber}).Delete(&entity.Pokemon{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
