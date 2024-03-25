package handler

import (
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/request"
	"github.com/redooz/podekex-hexagonal-architecture/application/dto/response"
)

type IPokedexHandler interface {
	SavePokemonToPokedex(pokedex *request.Pokedex) (httpStatus int, err error)
	GetAllPokemonFromPokedex() (response []*response.Pokedex, httpStatus int, err error)
	GetPokemonFromPokedexByNumber(pokedexNumber int) (response *response.Pokedex, httpStatus int, err error)
	UpdatePokemonInPokedex(pokedex *request.Pokedex, number int) (httpStatus int, err error)
	DeletePokemonFromPokedex(pokedexNumber int) (httpStatus int, err error)
}
