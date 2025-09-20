package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheck)

		r.Route("/boxes", func(r chi.Router) {
			r.Get("/", app.listBoxes)
			r.Post("/", app.createBox)

			r.Route("/{boxId}", func(r chi.Router) {
				r.Get("/", app.viewBox)
				r.Put("/", app.updateBox)
				r.Delete("/", app.deleteBox)
			})
		})
	})

	return r
}
