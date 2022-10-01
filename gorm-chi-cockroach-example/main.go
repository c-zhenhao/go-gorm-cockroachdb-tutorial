package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := "8082"
	log.Println("Running on http://localhost:" + port)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	log.Fatal(http.ListenAndServe(":"+port, r))
}
