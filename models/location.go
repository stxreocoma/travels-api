package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	ID       uint32 `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance uint32 `json:"distance"`
}
