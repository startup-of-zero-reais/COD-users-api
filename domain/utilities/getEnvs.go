package utilities

import (
	"log"
	"os"
	"strconv"
)

func GetEnv(key, _default string) string {
	if e := os.Getenv(key); e != "" {
		return e
	}

	return _default
}

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
