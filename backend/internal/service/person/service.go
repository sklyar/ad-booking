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

// Get returns a contact person by id.
func (s *Service) Get(ctx context.Context, id uint64) (*entity.ContactPerson, error) {
	person, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetch person: %w", err)
	}

	return person, nil
}

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
