package main

import (
	"io"
	"net/http"

	server "github.com/dklassen/ready_racer/proxy_server/pkg"
)

func main() {
	mux := http.NewServeMux()

	resultsHandler := http.HandlerFunc(server.RetrieveRaceResults)
	statusHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	mux.Handle("/", server.RequestLoggerMiddleware(resultsHandler))
	mux.Handle("/status", statusHandler)

	http.ListenAndServe(":8080", mux)
}
