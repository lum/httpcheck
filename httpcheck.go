package main

import (
	"github.com/lum/httpcheck/cassandracheck"
	"github.com/pivotal-golang/lager"
	"net/http"
)

var (
	logger           lager.Logger
	cassandrachecker *cassandracheck.Cassandrachecker
)

func handler(w http.ResponseWriter, r *http.Request, logger lager.Logger) {
	result := cassandrachecker.Check()
	if result {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func main() {
	cassandrachecker = cassandracheck.New(logger)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, logger)
	})
	http.ListenAndServe(":8080", nil)
}
