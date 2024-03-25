package request

type Pokedex struct {
	Number int    `json:"number" validate:"required"`
	Name   string `json:"name" validate:"required,min=3,max=50"`
	TypeId int    `json:"type_id" validate:"required"`
}

func NewPokedex(number int, name string, typeId int) *Pokedex {
	return &Pokedex{Number: number, Name: name, TypeId: typeId}
}
