package product

import "time"

type Model struct {
	ID uint
	Name string
	Observations string
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Models []*Model



