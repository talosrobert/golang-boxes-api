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
