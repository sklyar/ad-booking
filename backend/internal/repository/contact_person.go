package repository

import (
	"context"
	"fmt"

	"github.com/sklyar/ad-booking/backend/internal/entity"
)

// ContactPerson is an interface for contact person repository.
type ContactPerson interface {
	Create(ctx context.Context, person *entity.ContactPerson) error
	Delete(ctx context.Context, person *entity.ContactPerson) error
	Get(ctx context.Context, id uint64) (*entity.ContactPerson, error)
	Filter(ctx context.Context, filter ContactPersonFilter) ([]entity.ContactPerson, error)
}

// ContactPersonFilter is a filter for contact persons.
type ContactPersonFilter struct {
	Pagination Pagination
	Sorting    Sorting

	Name *string
	VKID *string
}

// Validate checks if the filter is valid.
func (f ContactPersonFilter) Validate(allowedFields []string) error {
	if err := f.Sorting.Validate(allowedFields); err != nil {
		return fmt.Errorf("order by: %w", err)
	}

	return nil
}
