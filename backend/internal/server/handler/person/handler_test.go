package person_test

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/api/gen/contactperson"
	"github.com/sklyar/ad-booking/backend/api/gen/contactperson/contactpersonconnect"
	"github.com/sklyar/ad-booking/backend/internal/test"
	"github.com/sklyar/ad-booking/backend/internal/test/apptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	ctx := context.Background()
	suite := apptest.NewSuite(t, ctx)

	req := new(contactperson.CreateRequest)
	test.ReadProtoFixture(t, "create/req.json", req)

	client := apptest.NewGRPCClient(suite, contactpersonconnect.NewServiceClient)
	resp, err := client.Create(ctx, connect.NewRequest(req))
	require.NoError(t, err)

	assert.Equal(t, req.Name, resp.Msg.ContactPerson.Name)
	assert.Equal(t, req.VkId, resp.Msg.ContactPerson.VkId)
}
