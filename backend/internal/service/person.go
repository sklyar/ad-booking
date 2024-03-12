package service

import (
	"context"
	"fmt"

	"github.com/sklyar/ad-booking/backend/internal/entity"
)

// Person is a contact person service.
type Person interface {
	// Create creates a new contact person.
	// Any error should be considered as an internal error.
	Create(ctx context.Context, data PersonCreate) (*entity.ContactPerson, error)

	// Delete deletes a contact person by his id.
	// It returns the deleted contact person if found.
	// If the contact person is not found, an apperror.UnknownPerson error is returned.
	// Any other error should be considered as an internal error.
	Delete(ctx context.Context, id uint64) (*entity.ContactPerson, error)

	// Get returns a contact person by id.
	// If the contact person is not found, an apperror.UnknownPerson error is returned.
	// Any other error should be considered as an internal error.
	Get(ctx context.Context, id uint64) (*entity.ContactPerson, error)

	// Filter returns a list of contact persons that match the filter.
	// Any error should be considered as an internal error.
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
