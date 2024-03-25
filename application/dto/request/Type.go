package request

type Type struct {
	FirstType  string `json:"first_type" validate:"required, min=3, max=50"`
	SecondType string `json:"second_type" validate:"required, min=3, max=50"`
}

func NewType(firstType string, secondType string) *Type {
	return &Type{FirstType: firstType, SecondType: secondType}
}
