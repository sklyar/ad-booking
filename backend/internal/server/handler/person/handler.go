package person

import (
	"github.com/sklyar/ad-booking/backend/api/gen/booking/bookingconnect"
	"github.com/sklyar/ad-booking/backend/internal/service"
)

type Handler struct {
	service service.Person
}

var _ bookingconnect.ContactPersonServiceHandler = (*Handler)(nil)

func New(service service.Person) *Handler {
	return &Handler{service: service}
}
