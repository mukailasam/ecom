package routers

import (
	"github.com/ftsog/ecom/handlers"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Route   *chi.Mux
	Handler *handlers.Handler
}
