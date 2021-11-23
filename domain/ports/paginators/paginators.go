package paginators

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/resources"
)

type (
	Metadata struct {
		Page    uint `json:"page,omitempty"`
		PerPage uint `json:"per_page,omitempty"`
		Total   uint `json:"total,omitempty"`
	}

	Links struct {
		Next     string `json:"next,omitempty"`
		Previous string `json:"previous,omitempty"`
		First    string `json:"first,omitempty"`
		Last     string `json:"last,omitempty"`
	}

	Paginated struct {
		Data     interface{} `json:"data,omitempty"`
		Metadata Metadata    `json:"_metadata"`
		Links    Links       `json:"_links"`
	}

	Paginator interface {
		Paginate(items []interface{ resources.Resource }) *Paginated
		GetPagination(c echo.Context) (uint, uint)
	}

	Pager struct {
		BaseURL string
		PerPage uint
		Page    uint
		Total   uint
	}
)
