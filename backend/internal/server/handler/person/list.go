package person

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sklyar/ad-booking/backend/internal/server/handler/common"

	"github.com/sklyar/ad-booking/backend/internal/entity"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/api/gen/booking"
	"github.com/sklyar/ad-booking/backend/internal/service"
)

func (h Handler) List(
	ctx context.Context,
	req *connect.Request[booking.ListPersonRequest],
) (*connect.Response[booking.ListPersonResponse], error) {
	persons, err := h.service.Filter(ctx, service.PersonFilter{
		Pagination: common.ConvertPagination(req.Msg.Pagination),
		Sorting:    common.ConvertSorting(req.Msg.Sorting),
		Name:       req.Msg.Name,
		VKID:       req.Msg.VkId,
	})
	if err != nil {
		return nil, err
	}

	return &connect.Response[booking.ListPersonResponse]{
		Msg: &booking.ListPersonResponse{
			Data: convertPersons(persons),
		},
	}, nil

}

func convertPersons(persons []entity.ContactPerson) []*booking.ContactPerson {
	res := make([]*booking.ContactPerson, len(persons))
	for i, p := range persons {
		res[i] = convertPerson(&p)
	}
	return res
}

func convertPerson(person *entity.ContactPerson) *booking.ContactPerson {
	return &booking.ContactPerson{
		Id:        person.ID,
		Name:      person.Name,
		VkId:      person.VKID,
		CreatedAt: timestamppb.New(person.CreatedAt),
		UpdatedAt: timestamppb.New(person.UpdatedAt),
	}
}
