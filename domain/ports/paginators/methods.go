package paginators

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

func (p *Pager) Paginate(items ...interface{}) *Paginated {
	var data interface{}
	if len(items) > 0 {
		data = items[0]
	} else {
		data = []interface{}{}
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

func (p *Pager) GetPagination(c echo.Context) (uint, uint) {
	page := int(p.Page)
	perPage := int(p.PerPage)
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

func (p *Pager) MountUrl(path string) string {
	base := getEnv("APPLICATION_HOST", "http://localhost:8080")
	return fmt.Sprintf("%s%s?%s", base, p.BaseURL, path)
}

func (p *Pager) MountPage(page, perPage uint) string {
	return fmt.Sprintf("page=%d&per_page=%d", page, perPage)
}

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

func (p *Pager) GetPrev() string {
	prev := p.Page - 1
	if prev <= 0 {
		return ""
	}

	return p.MountUrl(p.MountPage(prev, p.PerPage))
}

func (p *Pager) GetFirst() string {
	return p.MountUrl(p.MountPage(1, p.PerPage))
}

func (p *Pager) GetLast() string {
	last := p.Total / p.PerPage
	remaining := p.Total % p.PerPage
	if remaining == 0 {
		return p.MountUrl(p.MountPage(last, p.PerPage))
	}

	return p.MountUrl(p.MountPage(last+1, p.PerPage))
}

func getEnv(key, _default string) string {
	if e := os.Getenv(key); e != "" {
		return e
	}

	return _default
}
