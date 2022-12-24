package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/dot_product", dotProductHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal().Msgf("internal service exited %v", srv.ListenAndServe())
}

func dotProductHandler(w http.ResponseWriter, r *http.Request) {
	var x, y []int
	vars := mux.Vars(r)

	log.Info().Msgf("Handling request with params %v", vars)

	err := json.Unmarshal([]byte(vars["X"]), &x)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "bad input: %s", err.Error())
		log.Info().Msgf("Bad Request %v", err.Error())
	}

	err = json.Unmarshal([]byte(vars["Y"]), &y)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "bad input: %s", err.Error())
		log.Info().Msgf("Bad Request %v", err.Error())
	}

	result, err := DotProduct(x, y)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "bad input: %s", err.Error())
		log.Info().Msgf("Bad Request %v", err.Error())
	}

	response, err := json.Marshal(struct {
		Result int
	}{
		Result: result,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("InternalServerError: %v", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, response)

	log.Info().Msgf("Responded with %s", string(response))
}
