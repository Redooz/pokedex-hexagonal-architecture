package response

type Type struct {
	ID         int    `json:"id"`
	FirstType  string `json:"first_type"`
	SecondType string `json:"second_type"`
}

func NewType(ID int, firstType string, secondType string) *Type {
	return &Type{ID: ID, FirstType: firstType, SecondType: secondType}
}
