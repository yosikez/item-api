package model

import "time"

type Item struct {
	Id        uint
	Name      string
	Code      string
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
