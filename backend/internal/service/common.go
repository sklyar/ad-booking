package service

import "errors"

var (
	ErrInvalidOrderDirection = errors.New("invalid order direction")
	ErrInvalidOrderByField   = errors.New("invalid order by field")
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

func (o OrderBy) Validate(allowedFields []string) error {
	if !o.Direction.validate() {
		return ErrInvalidOrderDirection
	}

	var found bool
	for _, field := range allowedFields {
		if field == o.Field {
			found = true
			break
		}
	}
	if !found {
		return ErrInvalidOrderByField
	}

	return nil
}
