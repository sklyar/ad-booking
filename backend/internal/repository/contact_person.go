package repository

import (
	"context"
	"fmt"
	"github.com/sklyar/ad-booking/backend/internal/entity"
	"time"
)

// ContactPerson is an interface for contact person repository.
type ContactPerson interface {
	Create(ctx context.Context, person *entity.ContactPerson) error
}

type ContactPersonFilter struct {
	Name      *string
	VKID      *string
	CreatedAt *time.Time

	Pagination Pagination
	OrderBy    OrderBy
}

func (f ContactPersonFilter) Validate(allowedFields []string) error {
	if err := f.Pagination.Validate(); err != nil {
		return fmt.Errorf("pagination: %w", err)
	}

	if err := f.OrderBy.Validate(allowedFields); err != nil {
		return fmt.Errorf("order by: %w", err)
	}

	return nil
}
