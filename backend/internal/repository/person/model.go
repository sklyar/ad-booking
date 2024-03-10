package person

import (
	"time"

	"github.com/sklyar/ad-booking/backend/internal/entity"
)

type model struct {
	ID        uint64
	Name      string
	VKID      string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func newModel(v *entity.ContactPerson) model {
	return model{
		ID:        v.ID,
		Name:      v.Name,
		VKID:      v.VKID,
		CreatedAt: v.CreatedAt,
		UpdateAt:  v.UpdatedAt,
	}
}

func (m model) Convert() *entity.ContactPerson {
	return &entity.ContactPerson{
		ID:        m.ID,
		Name:      m.Name,
		VKID:      m.VKID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdateAt,
	}
}
