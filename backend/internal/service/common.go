package service

import "errors"

var (
	ErrInvalidOrderDirection = errors.New("invalid order direction")
	ErrRequiredLimit         = errors.New("required limit")
)

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

func (d OrderDirection) validate() bool {
	switch d {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

type Pagination struct {
	Limit  uint64
	Offset uint64
}

func (p Pagination) Validate() error {
	if p.Limit == 0 {
		return ErrRequiredLimit
	}

	return nil
}

type OrderBy struct {
	Field     string
	Direction OrderDirection
}

func (o OrderBy) Validate() error {
	if !o.Direction.validate() {
		return ErrInvalidOrderDirection
	}

	return nil
}
