package person_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sklyar/ad-booking/backend/internal/entity"

	"github.com/sklyar/ad-booking/backend/internal/service"

	"connectrpc.com/connect"
	"github.com/sklyar/ad-booking/backend/api/gen/booking"
	"github.com/sklyar/ad-booking/backend/api/gen/booking/bookingconnect"
	"github.com/sklyar/ad-booking/backend/internal/test"
	"github.com/sklyar/ad-booking/backend/internal/test/apptest"
	"github.com/stretchr/testify/require"
)

var testTime = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

func TestHandler_List(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	suite := apptest.NewSuite(t, ctx)

	// Create persons for test.
	persons := make([]*entity.ContactPerson, 5)
	personService := suite.App.ServiceContainer.PersonService
	for i := 0; i < 5; i++ {
		person, err := personService.Create(ctx, service.PersonCreate{
			Name: fmt.Sprintf("person-%d", i+1),
			VKID: fmt.Sprintf("vk-id-%d", i+1),
		})
		require.NoError(t, err)
		persons[i] = person
	}

	req := new(booking.ListPersonRequest)
	test.ReadProtoFixture(t, "list/req.json", req)

	expResp := new(booking.ListPersonResponse)
	test.ReadProtoFixture(t, "list/resp.json", expResp)

	client := apptest.NewGRPCClient(suite, bookingconnect.NewContactPersonServiceClient)
	resp, err := client.List(ctx, connect.NewRequest(req))
	require.NoError(t, err)

	for i := range resp.Msg.Data {
		resp.Msg.Data[i].CreatedAt = timestamppb.New(testTime)
		resp.Msg.Data[i].UpdatedAt = timestamppb.New(testTime)
	}

	test.AssertProtoEq(t, expResp, resp.Msg)
}
