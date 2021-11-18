package paginators

import (
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/paginators"
	"os"
	"strconv"
)

func NewPaginator(baseURL string) paginators.Pager {
	defaultPerPage := getEnvUint("DEFAULT_PAGINATOR_PER_PAGE", 10)
	defaultPage := getEnvUint("DEFAULT_PAGINATOR_PAGE", 1)

	return paginators.Pager{
		BaseURL: baseURL,
		PerPage: defaultPerPage,
		Page:    defaultPage,
	}
}

//func getEnv(key string, _default string) string {
//	env := os.Getenv(key)
//
//	if env != "" {
//		return env
//	}
//
//	return _default
//}

func getEnvUint(key string, _default uint) uint {
	env := os.Getenv(key)

	if env != "" {
		envInt, _ := strconv.Atoi(env)
		return uint(envInt)
	}

	return _default
}
