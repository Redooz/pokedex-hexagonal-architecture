package request

type Pokedex struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	TypeId int    `json:"type_id"`
}

func NewPokedex(number int, name string, typeId int) *Pokedex {
	return &Pokedex{Number: number, Name: name, TypeId: typeId}
}
