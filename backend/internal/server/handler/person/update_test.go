package person_test

import (
	"context"
	"testing"

	"github.com/sklyar/ad-booking/backend/internal/test"

	"github.com/sklyar/ad-booking/backend/internal/service"

	"github.com/sklyar/ad-booking/backend/api/gen/booking"
	"github.com/sklyar/ad-booking/backend/api/gen/booking/bookingconnect"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/internal/test/apptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Update(t *testing.T) {
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

	req := &booking.UpdatePersonRequest{
		Id:   person.ID,
		Name: test.ToPtr("person-updated"),
		VkId: test.ToPtr("vk-id-updated"),
	}

	// Act.
	resp, err := client.Update(ctx, connect.NewRequest(req))
	require.NoError(t, err)

	// Assert.
	assert.Equal(t, *req.Name, resp.Msg.ContactPerson.Name)
	assert.Equal(t, *req.VkId, resp.Msg.ContactPerson.VkId)
}
