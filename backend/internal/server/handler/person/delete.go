package person

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/api/gen/booking"
)

func (h Handler) Delete(
	ctx context.Context,
	req *connect.Request[booking.DeletePersonRequest],
) (*connect.Response[emptypb.Empty], error) {
	_, err := h.service.Delete(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return &connect.Response[emptypb.Empty]{}, nil
}
