package entity

import "gorm.io/gorm"

type TypeEntity struct {
	ID         int    `json:"id" gorm:"primary_key;autoIncrement;not null;unique"`
	FirstType  string `json:"first_type" gorm:"not null;unique"`
	SecondType string `json:"second_type" gorm:"not null;unique"`
	PokemonID  int
	gorm.Model
}
