package service

import (
	"context"
	"github.com/sklyar/ad-booking/backend/internal/entity"
)

// Person is a contact person service.
type Person interface {
	Create(ctx context.Context, data PersonCreate) (*entity.ContactPerson, error)
}

// PersonCreate is a dto for creating a new contact person.
type PersonCreate struct {
	Name string
	VKID string
}
