package person

import (
	"context"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/api/gen/booking"
)

func (h Handler) Get(
	ctx context.Context,
	req *connect.Request[booking.GetPersonRequest],
) (*connect.Response[booking.GetPersonResponse], error) {
	person, err := h.service.Get(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return &connect.Response[booking.GetPersonResponse]{
		Msg: &booking.GetPersonResponse{
			ContactPerson: convertPerson(person),
		},
	}, nil
}
