package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/youngjae-lim/gosnel"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.get("/", a.Handlers.Home)

	// static files
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// routes from gosnel
	a.App.Routes.Mount("/gosnel", gosnel.Routes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
