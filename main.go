package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ftsog/ecom/settings"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	router := settings.NewRouter(r)
	router.Routing()

	server := &http.Server{
		Handler:      router.Route,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("listening on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
