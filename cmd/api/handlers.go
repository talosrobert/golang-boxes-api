package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment %s\n", app.cfg.env)
	fmt.Fprintf(w, "api version %s\n", version)
}

func (app *application) listBoxes(w http.ResponseWriter, r *http.Request) {}
func (app *application) createBox(w http.ResponseWriter, r *http.Request) {}

func (app *application) viewBox(w http.ResponseWriter, r *http.Request) {
	id, err := app.fetchBoxIdFromParams(r.Context())
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "box %d", id)
}

func (app *application) updateBox(w http.ResponseWriter, r *http.Request) {}
func (app *application) deleteBox(w http.ResponseWriter, r *http.Request) {}
