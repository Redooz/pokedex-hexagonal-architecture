package response

type Pokedex struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Name   string `json:"name"`
	Type   Type   `json:"type"`
}
