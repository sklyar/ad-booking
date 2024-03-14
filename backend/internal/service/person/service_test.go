package person

import (
	"context"
	"testing"

	"github.com/sklyar/ad-booking/backend/internal/test"

	"github.com/sklyar/ad-booking/backend/internal/service"
	"github.com/sklyar/ad-booking/backend/internal/test/mock"
	"github.com/stretchr/testify/assert"

	"github.com/sklyar/ad-booking/backend/internal/entity"
)

func TestService_Update(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(r *mock.ContactPersonRepository)
		data    service.PersonUpdate
		want    *entity.ContactPerson
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "update contact person",
			setup: func(r *mock.ContactPersonRepository) {
				person := &entity.ContactPerson{ID: 1, Name: "name", VKID: "vkid"}
				r.EXPECT().Get(context.Background(), person.ID).Return(person, nil)

				updatedPerson := &entity.ContactPerson{ID: 1, Name: "new name", VKID: "new vkid"}
				r.EXPECT().Update(context.Background(), updatedPerson).Return(nil)
			},
			data:    service.PersonUpdate{ID: 1, Name: test.ToPtr("new name"), VKID: test.ToPtr("new vkid")},
			want:    &entity.ContactPerson{ID: 1, Name: "new name", VKID: "new vkid"},
			wantErr: assert.NoError,
		},
		{
			name: "no changes",
			setup: func(r *mock.ContactPersonRepository) {
				person := &entity.ContactPerson{ID: 1, Name: "name", VKID: "vkid"}
				r.EXPECT().Get(context.Background(), person.ID).Return(person, nil)
			},
			data:    service.PersonUpdate{ID: 1, Name: test.ToPtr("name"), VKID: test.ToPtr("vkid")},
			want:    &entity.ContactPerson{ID: 1, Name: "name", VKID: "vkid"},
			wantErr: assert.NoError,
		},
		{
			name: "get person error",
			setup: func(r *mock.ContactPersonRepository) {
				r.EXPECT().Get(context.Background(), uint64(1)).Return(nil, assert.AnError)
			},
			data: service.PersonUpdate{ID: 1, Name: test.ToPtr("name"), VKID: test.ToPtr("vkid")},
			want: nil,
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				return assert.ErrorIs(t, err, assert.AnError)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			contactPersonRepository := mock.NewContactPersonRepository(t)
			tt.setup(contactPersonRepository)

			s := New(contactPersonRepository)
			got, err := s.Update(context.Background(), tt.data)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
