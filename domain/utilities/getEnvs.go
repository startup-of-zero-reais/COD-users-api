package utilities

import (
	"log"
	"os"
	"strconv"
)

// GetEnv é uma função para recuperar as variáveis a em os.GetEnv,
// porém com um fallback de valor default caso a variável não esteja definida
func GetEnv(key, _default string) string {
	if e := os.Getenv(key); e != "" {
		return e
	}

	return _default
}

// GetEnvUint é uma função assim como GetEnv, porém retorna o valor em uint
func GetEnvUint(key string, _default uint) uint {
	if e := os.Getenv(key); e != "" {
		intEnv, err := strconv.Atoi(e)
		if err != nil {
			log.Fatalln("Erro ao recuperar env var")
		}

		return uint(intEnv)
	}

	return _default
}
