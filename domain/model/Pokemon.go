package model

type Pokemon struct {
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	Name        string `json:"name"`
	PokemonType Type   `json:"type"`
}

func NewPokemon(ID int, number int, name string, pokemonType Type) *Pokemon {
	return &Pokemon{ID: ID, Number: number, Name: name, PokemonType: pokemonType}
}
