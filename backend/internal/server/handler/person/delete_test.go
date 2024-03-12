package person_test

import (
	"context"
	"testing"

	"github.com/sklyar/ad-booking/backend/internal/service"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/api/gen/booking"
	"github.com/sklyar/ad-booking/backend/api/gen/booking/bookingconnect"
	"github.com/sklyar/ad-booking/backend/internal/test/apptest"
	"github.com/stretchr/testify/require"
)

func TestHandler_Delete(t *testing.T) {
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

	client := apptest.NewGRPCClient(suite, bookingconnect.NewContactPersonServiceClient)

	req := &booking.DeletePersonRequest{Id: person.ID}

	// Act.
	_, err = client.Delete(ctx, connect.NewRequest(req))
	require.NoError(t, err)

	// Assert.
	_, err = personService.Get(ctx, person.ID)
	require.Error(t, err, "expected person to be deleted")
}
