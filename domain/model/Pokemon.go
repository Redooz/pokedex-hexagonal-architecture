package model

type Pokemon struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Name   string `json:"name"`
	TypeId int    `json:"type_id"`
}

func NewPokemon(ID int, number int, name string, typeId int) *Pokemon {
	return &Pokemon{ID: ID, Number: number, Name: name, TypeId: typeId}
}
