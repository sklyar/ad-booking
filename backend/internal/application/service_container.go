package application

import (
	"github.com/sklyar/ad-booking/backend/internal/service"
	"github.com/sklyar/ad-booking/backend/internal/service/person"
)

// ServiceContainer is a container for services.
type ServiceContainer struct {
	PersonService service.Person
}

// NewServiceContainer creates a new service container.
func newServiceContainer(repositoryContainer *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		PersonService: person.New(repositoryContainer.ContactPersonRepository),
	}
}
