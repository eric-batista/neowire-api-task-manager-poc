package models

import "time"

type Base struct {
	ID        string    `json:"id" validate:"uuid"`
	CreatedAt time.Time `json:"createdAt" validate:"-"`
	UpdatedAt time.Time `json:"updatedAt" validate:"-"`
}
