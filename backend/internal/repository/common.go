package repository

import (
	"errors"
	"slices"
)

var (
	ErrInvalidOrderByField = errors.New("invalid order by field")
)

// OrderDirection is a type for sorting direction.
type OrderDirection string

const (
	// OrderDirectionAsc is an ascending sorting direction.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc is a descending sorting direction.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Pagination is a dto for pagination.
type Pagination struct {
	LastID uint64
	Limit  int64
}

// Sorting is a dto for sorting.
type Sorting struct {
	Field     string
	Direction OrderDirection
}

// Validate checks if the sorting is valid.
func (o Sorting) Validate(allowedFields []string) error {
	if !slices.Contains(allowedFields, o.Field) {
		return ErrInvalidOrderByField
	}

	return nil
}
