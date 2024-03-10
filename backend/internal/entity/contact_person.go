package entity

import "time"

type ContactPerson struct {
	ID        uint64
	Name      string
	VKID      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
