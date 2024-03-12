package person_test

import (
	"context"
	"testing"

	"github.com/sklyar/ad-booking/backend/api/gen/booking"
	"github.com/sklyar/ad-booking/backend/api/gen/booking/bookingconnect"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/internal/test/apptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	suite := apptest.NewSuite(t, ctx)

	req := &booking.CreatePersonRequest{
		Name: "person-1",
		VkId: "vk-id1",
	}

	client := apptest.NewGRPCClient(suite, bookingconnect.NewContactPersonServiceClient)
	resp, err := client.Create(ctx, connect.NewRequest(req))
	require.NoError(t, err)

	assert.Equal(t, req.Name, resp.Msg.ContactPerson.Name)
	assert.Equal(t, req.VkId, resp.Msg.ContactPerson.VkId)
}
