package main

import (
	"net/http"

	server "github.com/dklassen/ready_racer/proxy_server/pkg"
)

func main() {
	mux := http.NewServeMux()

	resultsHandler := http.HandlerFunc(server.RetrieveRaceResults)

	mux.Handle("/", server.RequestLoggerMiddleware(resultsHandler))

	http.ListenAndServe(":8080", mux)
}
