package person

import (
	"context"

	"github.com/sklyar/ad-booking/backend/api/gen/booking"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/internal/service"
)

func (h Handler) Create(
	ctx context.Context,
	req *connect.Request[booking.CreatePersonRequest],
) (*connect.Response[booking.CreatePersonResponse], error) {
	person, err := h.service.Create(ctx, service.PersonCreate{
		Name: req.Msg.Name,
		VKID: req.Msg.VkId,
	})
	if err != nil {
		return nil, err
	}

	return &connect.Response[booking.CreatePersonResponse]{
		Msg: &booking.CreatePersonResponse{
			ContactPerson: convertPerson(person),
		},
	}, nil
}
