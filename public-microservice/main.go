package main

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

const (
	INTERNAL_SERVICE_ADDRESS_ENV_VAR = "INTERNAL_SERVICE_ADDRESS"
)

func main() {
	internalSvcAddress := os.Getenv(INTERNAL_SERVICE_ADDRESS_ENV_VAR)

	if internalSvcAddress == "" {
		log.Fatal().Msgf("%s env var is not set", INTERNAL_SERVICE_ADDRESS_ENV_VAR)
	}

	http.HandleFunc("/", directHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal().Msgf("ListenAndServe: ", err)
	}
}

func directHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("Redirecting call to %s", INTERNAL_SERVICE_ADDRESS_ENV_VAR)
	http.Redirect(w, r, INTERNAL_SERVICE_ADDRESS_ENV_VAR, 302)
}
