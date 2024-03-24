package handler

import (
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/response"
)

type IPokedexHandler interface {
	SavePokemonToPokedex(pokedex *request.Pokedex) error
	GetAllPokemonFromPokedex() ([]*response.Pokedex, error)
	GetPokemonFromPokedexByNumber(pokedexNumber int) (*response.Pokedex, error)
	UpdatePokemonInPokedex(pokedex *request.Pokedex) error
	DeletePokemonFromPokedex(pokedexNumber int) error
}
