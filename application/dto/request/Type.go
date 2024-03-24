package request

type Type struct {
	FirstType  string `json:"first_type"`
	SecondType string `json:"second_type"`
}

func NewType(firstType string, secondType string) *Type {
	return &Type{FirstType: firstType, SecondType: secondType}
}
