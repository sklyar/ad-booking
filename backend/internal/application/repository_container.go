package application

import (
	"github.com/sklyar/ad-booking/backend/internal/infrastructure/db"
	"github.com/sklyar/ad-booking/backend/internal/repository"
	"github.com/sklyar/ad-booking/backend/internal/repository/person"
)

type RepositoryContainer struct {
	ContactPersonRepository repository.ContactPerson
}

func newRepositoryContainer(dbHandler db.Handler) *RepositoryContainer {
	return &RepositoryContainer{
		ContactPersonRepository: person.New(dbHandler),
	}
}
