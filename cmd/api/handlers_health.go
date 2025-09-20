package main

import (
	"net/http"

	"github.com/go-chi/render"
)

type HealthResponse struct {
	HTTPStatusCode int    `json:"-"`
	Status         string `json:"status"`
	Environment    string `json:"environment"`
	APIVersion     string `json:"api_version"`
}

func (resp *HealthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, resp.HTTPStatusCode)
	return nil
}

func HealthRender(env, apiVersion string) render.Renderer {
	return &HealthResponse{
		HTTPStatusCode: http.StatusOK,
		Status:         "available",
		Environment:    env,
		APIVersion:     apiVersion,
	}
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, HealthRender("dev", "0.0.1"))
}
