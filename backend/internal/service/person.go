package service

import (
	"context"
	"fmt"

	"github.com/sklyar/ad-booking/backend/internal/entity"
)

// Person is a contact person service.
type Person interface {
	Create(ctx context.Context, data PersonCreate) (*entity.ContactPerson, error)
	Get(ctx context.Context, id uint64) (*entity.ContactPerson, error)
	Filter(ctx context.Context, filter PersonFilter) ([]entity.ContactPerson, error)
}

// PersonCreate is a dto for creating a new contact person.
type PersonCreate struct {
	Name string
	VKID string
}

// PersonFilter is a dto for filtering contact persons.
type PersonFilter struct {
	Pagination Pagination
	Sorting    Sorting

	Name *string
	VKID *string
}

// Validate checks if the filter is valid.
func (f PersonFilter) Validate() error {
	if err := f.Pagination.Validate(); err != nil {
		return fmt.Errorf("pagination: %w", err)
	}

	if err := f.Sorting.Validate(); err != nil {
		return fmt.Errorf("sorting: %w", err)
	}

	return nil
}
