package main

import (
	"net/http"

	"github.com/go-chi/render"
)

type Health struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	APIVersion  string `json:"api_version"`
}

type HealthResponse struct {
	*Health        `json:"health"`
	HTTPStatusCode int `json:"-"`
}

func NewHealthResponse(h *Health, statusCode int) *HealthResponse {
	return &HealthResponse{
		Health:         h,
		HTTPStatusCode: statusCode,
	}
}

func (resp *HealthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, resp.HTTPStatusCode)
	return nil
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	resp := NewHealthResponse(
		&Health{"available", "dev", "0.0.1"},
		http.StatusOK,
	)

	if err := render.Render(w, r, resp); err != nil {
		app.logger.Error("health check response failed", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
