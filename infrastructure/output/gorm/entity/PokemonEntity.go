package entity

import "gorm.io/gorm"

type PokemonEntity struct {
	ID     int    `json:"id" gorm:"primary_key;autoIncrement;not null;unique"`
	Number int    `json:"number" gorm:"not null;unique"`
	Name   string `json:"name" gorm:"not null;unique"`
	Type   TypeEntity
	gorm.Model
}
