package common

import (
	"github.com/sklyar/ad-booking/backend/api/gen/common"
	"github.com/sklyar/ad-booking/backend/internal/service"
)

func ConvertPagination(pagination *common.Pagination) service.Pagination {
	return service.Pagination{
		LastID: pagination.LastId,
		Limit:  pagination.Limit,
	}
}
