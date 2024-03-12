package person_test

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sklyar/ad-booking/backend/internal/service"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/api/gen/booking"
	"github.com/sklyar/ad-booking/backend/api/gen/booking/bookingconnect"
	"github.com/sklyar/ad-booking/backend/internal/test"
	"github.com/sklyar/ad-booking/backend/internal/test/apptest"
	"github.com/stretchr/testify/require"
)

func TestHandler_Get(t *testing.T) {
	t.Parallel()

	// Arrange.
	ctx := context.Background()
	suite := apptest.NewSuite(t, ctx)

	personService := suite.App.ServiceContainer.PersonService
	person, err := personService.Create(ctx, service.PersonCreate{
		Name: "person-1",
		VKID: "vk-id-1",
	})
	require.NoError(t, err)

	exp := &booking.GetPersonResponse{
		ContactPerson: &booking.ContactPerson{
			Id:        person.ID,
			Name:      person.Name,
			VkId:      person.VKID,
			CreatedAt: timestamppb.New(person.CreatedAt),
			UpdatedAt: timestamppb.New(person.UpdatedAt),
		},
	}

	req := new(booking.GetPersonRequest)
	req.Id = person.ID

	client := apptest.NewGRPCClient(suite, bookingconnect.NewContactPersonServiceClient)

	// Act.
	resp, err := client.Get(ctx, connect.NewRequest(req))
	require.NoError(t, err)

	// Assert.
	test.AssertProtoEq(t, exp, resp.Msg)
}
