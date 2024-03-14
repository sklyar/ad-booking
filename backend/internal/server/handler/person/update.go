package person

import (
	"context"

	"github.com/sklyar/ad-booking/backend/api/gen/booking"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/internal/service"
)

func (h Handler) Update(
	ctx context.Context,
	req *connect.Request[booking.UpdatePersonRequest],
) (*connect.Response[booking.UpdatePersonResponse], error) {
	person, err := h.service.Update(ctx, service.PersonUpdate{
		ID:   req.Msg.Id,
		Name: req.Msg.Name,
		VKID: req.Msg.VkId,
	})
	if err != nil {
		return nil, err
	}

	return &connect.Response[booking.UpdatePersonResponse]{
		Msg: &booking.UpdatePersonResponse{
			ContactPerson: convertPerson(person),
		},
	}, nil
}
