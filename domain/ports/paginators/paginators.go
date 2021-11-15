package paginators

type (
	Metadata struct {
		Page    uint `json:"page,omitempty"`
		PerPage uint `json:"per_page,omitempty"`
		Total   uint `json:"total,omitempty"`
	}

	Links struct {
		Self     string `json:"self,omitempty"`
		Next     string `json:"next,omitempty"`
		Previous string `json:"previous,omitempty"`
		First    string `json:"first,omitempty"`
		Last     string `json:"last,omitempty"`
	}

	Paginated struct {
		Data     []interface{} `json:"data,omitempty"`
		Metadata Metadata      `json:"metadata"`
		Links    Links         `json:"links"`
	}

	Paginator interface {
		Paginate(items []interface{}) *Paginated
	}
)
