package person

import (
	"context"
	"fmt"

	"github.com/sklyar/ad-booking/backend/internal/entity"
	"github.com/sklyar/ad-booking/backend/internal/repository"
	"github.com/sklyar/ad-booking/backend/internal/service"
)

// Service is a contact person service.
type Service struct {
	repo repository.ContactPerson
}

var _ service.Person = (*Service)(nil)

// New creates a new contact person service.
func New(repo repository.ContactPerson) *Service {
	return &Service{repo: repo}
}

// Create creates a new contact person.
func (s *Service) Create(ctx context.Context, data service.PersonCreate) (*entity.ContactPerson, error) {
	person := &entity.ContactPerson{
		Name: data.Name,
		VKID: data.VKID,
	}

	if err := s.repo.Create(ctx, person); err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return person, nil
}

// Update updates a contact person.
// It returns the updated contact person if found.
// If the contact person is not found, an apperror.UnknownPerson error is returned.
// Any other error should be considered as an internal error.
func (s *Service) Update(ctx context.Context, data service.PersonUpdate) (*entity.ContactPerson, error) {
	person, err := s.Get(ctx, data.ID)
	if err != nil {
		return nil, fmt.Errorf("get person: %w", err)
	}

	var updated bool
	if data.Name != nil && *data.Name != person.Name {
		person.Name = *data.Name
		updated = true
	}
	if data.VKID != nil && *data.VKID != person.VKID {
		person.VKID = *data.VKID
		updated = true
	}
	if !updated {
		return person, nil
	}

	if err := s.repo.Update(ctx, person); err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	return person, nil
}

// Delete deletes a contact person by his id.
// It returns the deleted contact person if found.
// If the contact person is not found, an apperror.UnknownPerson error is returned.
// Any other error should be considered as an internal error.
func (s *Service) Delete(ctx context.Context, id uint64) (*entity.ContactPerson, error) {
	person, err := s.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get person: %w", err)
	}

	if err := s.repo.Delete(ctx, person); err != nil {
		return nil, fmt.Errorf("delete: %w", err)
	}

	return person, nil
}

// Get returns a contact person by id.
func (s *Service) Get(ctx context.Context, id uint64) (*entity.ContactPerson, error) {
	person, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetch person: %w", err)
	}
	if person == nil {
		// TODO: Wrap error with a custom error type.
		return nil, fmt.Errorf("person not found")
	}

	return person, nil
}

// Filter returns a list of contact persons that match the filter.
// Any error should be considered as an internal error.
func (s *Service) Filter(ctx context.Context, filter service.PersonFilter) ([]entity.ContactPerson, error) {
	if err := filter.Validate(); err != nil {
		return nil, fmt.Errorf("validate filter: %w", err)
	}

	persons, err := s.repo.Filter(ctx, repository.ContactPersonFilter{
		Pagination: repository.Pagination{
			LastID: filter.Pagination.LastID,
			Limit:  filter.Pagination.Limit,
		},
		Sorting: repository.Sorting{
			Field:     filter.Sorting.Field,
			Direction: repository.OrderDirection(filter.Sorting.Direction),
		},
		Name: filter.Name,
		VKID: filter.VKID,
	})
	if err != nil {
		return nil, fmt.Errorf("fetch persons: %w", err)
	}

	return persons, nil
}
