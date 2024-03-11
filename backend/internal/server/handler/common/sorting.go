package common

import (
	"github.com/sklyar/ad-booking/backend/api/gen/common"
	"github.com/sklyar/ad-booking/backend/internal/service"
)

func ConvertSorting(sorting *common.Sorting) service.Sorting {
	if sorting == nil {
		return service.Sorting{}
	}

	return service.Sorting{
		Field:     sorting.Field,
		Direction: convertDirection(sorting.Direction),
	}
}

func convertDirection(direction common.Sorting_Direction) service.OrderDirection {
	switch direction {
	case common.Sorting_DIRECTION_ASC:
		return service.OrderDirectionAsc
	case common.Sorting_DIRECTION_DESC:
		return service.OrderDirectionDesc
	}

	return service.OrderDirectionUnspecified
}
