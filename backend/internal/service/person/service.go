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
