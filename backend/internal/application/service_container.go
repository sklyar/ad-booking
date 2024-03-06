package application

import (
	"github.com/sklyar/ad-booking/backend/internal/service"
)

type ServiceContainer struct {
	ContactPersonService service.Person
}

func newServiceContainer(repositoryContainer *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		ContactPersonService: contactperson.New(repositoryContainer.ContactPersonRepository),
	}
}
