package person

import (
	"context"
	"github.com/sklyar/ad-booking/backend/internal/test"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	ctx := context.Background()
	app := test.NewSuite(t, ctx)

}
