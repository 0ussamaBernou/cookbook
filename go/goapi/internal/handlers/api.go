package handlers

import (
	"cookbook/go/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chiddleware "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	//Global middleware
	r.Use(chiddleware.StripSlashes)

	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization)

		r.Get("/coins", GetCoinBalance)
	})
}
