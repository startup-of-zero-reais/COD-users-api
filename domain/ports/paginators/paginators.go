package paginators

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/resources"
)

type (
	// Metadata é a estrutura de metadados na resposta da API
	Metadata struct {
		Page    uint `json:"page,omitempty"`
		PerPage uint `json:"per_page,omitempty"`
		Total   uint `json:"total,omitempty"`
	}

	// Links é a estrutura de links na resposta da API
	Links struct {
		Next     string `json:"next,omitempty"`
		Previous string `json:"previous,omitempty"`
		First    string `json:"first,omitempty"`
		Last     string `json:"last,omitempty"`
	}

	// Paginated é a estrutura final de resposta da API
	Paginated struct {
		Data     interface{} `json:"data,omitempty"`
		Metadata Metadata    `json:"_metadata"`
		Links    Links       `json:"_links"`
	}

	// Paginator é a interface para que os recursos tenham os métodos necessários
	// para gerar paginação de request na API
	Paginator interface {
		Paginate(items []interface{ resources.Resource }) *Paginated
		GetPagination(c echo.Context) (uint, uint)
	}

	// Pager é a estrutura base de paginação
	//
	// BaseURL é a url de base para a referência da paginação
	Pager struct {
		BaseURL string
		PerPage uint
		Page    uint
		Total   uint
	}
)
