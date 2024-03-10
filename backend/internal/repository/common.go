package repository

import (
	"errors"
	"slices"
)

var (
	ErrInvalidOrderByField = errors.New("invalid order by field")
)

type OrderDirection string

type Pagination struct {
	Limit  uint64
	Offset uint64
}

type OrderBy struct {
	Field     string
	Direction OrderDirection
}

func (o OrderBy) Validate(allowedFields []string) error {
	if !slices.Contains(allowedFields, o.Field) {
		return ErrInvalidOrderByField
	}

	return nil
}
