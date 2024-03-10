package person

import (
	"context"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/api/gen/contactperson"
	"github.com/sklyar/ad-booking/backend/api/gen/contactperson/contactpersonconnect"
	"github.com/sklyar/ad-booking/backend/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	service service.Person
}

var _ contactpersonconnect.ServiceClient = (*Handler)(nil)

func New(service service.Person) *Handler {
	return &Handler{service: service}
}

func (h Handler) Create(
	ctx context.Context,
	req *connect.Request[contactperson.CreateRequest],
) (*connect.Response[contactperson.CreateResponse], error) {
	person, err := h.service.Create(ctx, service.PersonCreate{
		Name: req.Msg.Name,
		VKID: req.Msg.VkId,
	})
	if err != nil {
		return nil, err
	}

	return &connect.Response[contactperson.CreateResponse]{
		Msg: &contactperson.CreateResponse{
			ContactPerson: &contactperson.ContactPerson{
				Id:        person.ID,
				Name:      person.Name,
				VkId:      person.VKID,
				CreatedAt: timestamppb.New(person.CreatedAt),
				UpdatedAt: timestamppb.New(person.UpdatedAt),
			},
		},
	}, nil
}
