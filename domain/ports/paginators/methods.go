package paginators

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/resources"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"strconv"
)

// Paginate é o método responsável por gerar a paginação de recursos
func (p *Pager) Paginate(items []interface{ resources.Resource }) *Paginated {
	var data []resources.Resource
	if len(items) > 0 {
		for _, item := range items {
			item.GetEmbedded()
			data = append(data, item)
		}
	} else {
		data = make([]resources.Resource, 1)
	}

	return &Paginated{
		Data: data,
		Metadata: Metadata{
			Page:    p.Page,
			PerPage: p.PerPage,
			Total:   p.Total,
		},
		Links: Links{
			Next:     p.GetNext(),
			Previous: p.GetPrev(),
			First:    p.GetFirst(),
			Last:     p.GetLast(),
		},
	}
}

// GetPagination extrai as paginações vindas do contexto da request
// como "?page=1&per_page=10". Caso não haja paginação será atribuída a padrão
//
// GetPagination retorna o número da página e a quantidade de resultados por página
func (p *Pager) GetPagination(c echo.Context) (uint, uint) {
	page, _ := strconv.Atoi(utilities.GetEnv("DEFAULT_PAGINATOR_PAGE", "1"))
	perPage, _ := strconv.Atoi(utilities.GetEnv("DEFAULT_PAGINATOR_PER_PAGE", "20"))

	if pp := c.QueryParam("per_page"); pp != "" {
		perPage, _ = strconv.Atoi(pp)
	}

	if qp := c.QueryParam("page"); qp != "" {
		page, _ = strconv.Atoi(qp)
	}

	p.Page = uint(page)
	p.PerPage = uint(perPage)

	return p.Page, p.PerPage
}

// MountUrl monta a url de base para Paginated.Links
func (p *Pager) MountUrl(path string) string {
	base := utilities.GetEnv("APPLICATION_HOST", "http://localhost:8080")
	return fmt.Sprintf("%s%s?%s", base, p.BaseURL, path)
}

// MountPage monta os query params de paginação
//
// Ex.: ?page=1&per_page=10
func (p *Pager) MountPage(page, perPage uint) string {
	return fmt.Sprintf("page=%d&per_page=%d", page, perPage)
}

// GetNext calcula a próxima página de resultados.
// Pode não existir caso não haja próxima página
func (p *Pager) GetNext() string {
	next := p.Page + 1
	numPages := p.Total / p.PerPage

	if (p.Total % p.PerPage) > 0 {
		numPages = numPages + 1
	}

	if next <= numPages {
		return p.MountUrl(p.MountPage(next, p.PerPage))
	}

	return ""
}

// GetPrev calcula qual a página anterior de resultados.
// Pode não existir caso seja a primeira página
func (p *Pager) GetPrev() string {
	prev := p.Page - 1
	if prev <= 0 {
		return ""
	}

	return p.MountUrl(p.MountPage(prev, p.PerPage))
}

// GetFirst monta o link para a primeira página de resultados
func (p *Pager) GetFirst() string {
	return p.MountUrl(p.MountPage(1, p.PerPage))
}

// GetLast calcula qual a última página de resultados
func (p *Pager) GetLast() string {
	last := p.Total / p.PerPage
	remaining := p.Total % p.PerPage
	if remaining == 0 {
		return p.MountUrl(p.MountPage(last, p.PerPage))
	}

	return p.MountUrl(p.MountPage(last+1, p.PerPage))
}
