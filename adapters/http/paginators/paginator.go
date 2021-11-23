package paginators

import (
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
)

func NewPaginator(baseURL string) paginators.Pager {
	defaultPerPage := utilities.GetEnvUint("DEFAULT_PAGINATOR_PER_PAGE", 10)
	defaultPage := utilities.GetEnvUint("DEFAULT_PAGINATOR_PAGE", 1)

	return paginators.Pager{
		BaseURL: baseURL,
		PerPage: defaultPerPage,
		Page:    defaultPage,
	}
}
