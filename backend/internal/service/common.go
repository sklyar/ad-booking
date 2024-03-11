package service

import "errors"

var (
	ErrInvalidOrderDirection = errors.New("invalid order direction")
	ErrRequiredLimit         = errors.New("required limit")
)

// OrderDirection is a type for sorting direction.
type OrderDirection string

// OrderDirection values.
const (
	OrderDirectionUnspecified OrderDirection = ""
	OrderDirectionAsc         OrderDirection = "ASC"
	OrderDirectionDesc        OrderDirection = "DESC"
)

// Validate checks if the order direction is valid.
func (d OrderDirection) validate() bool {
	switch d {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

// Pagination is a dto for pagination.
type Pagination struct {
	LastID uint64
	Limit  int64
}

// Validate checks if the pagination is valid.
func (p Pagination) Validate() error {
	if p.Limit == 0 {
		return ErrRequiredLimit
	}

	return nil
}

// Sorting is a dto for sorting.
type Sorting struct {
	Field     string
	Direction OrderDirection
}

// Validate checks if the sorting is valid.
func (o Sorting) Validate() error {
	if !o.Direction.validate() {
		return ErrInvalidOrderDirection
	}

	return nil
}
