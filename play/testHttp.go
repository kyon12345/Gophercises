package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "hello")
	})

	mux.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "world")
	})

	mux.HandleFunc("/node/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "node list")
	})

	mux.Handle("/middleware", middlewareOne(middlewareTwo(http.HandlerFunc(final))))

	http.ListenAndServe(":9090", mux)
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Print("final")
}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareTwo")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareTwo again")
	})
}
