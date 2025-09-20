package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "0.0.1"

type config struct {
	port uint
	env  string
}

type application struct {
	cfg    config
	logger *slog.Logger
}

func newApplication(cfg config, logger *slog.Logger) *application {
	return &application{
		cfg:    cfg,
		logger: logger,
	}
}

func main() {
	var cfg config

	flag.UintVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "environment", "dev", "API server environment")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := newApplication(cfg, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", app.healthCheck)
	mux.HandleFunc("GET /v1/boxes/", app.healthCheck)
	mux.HandleFunc("POST /v1/boxes/", app.healthCheck)
	mux.HandleFunc("GET /v1/boxes/:id", app.healthCheck)
	mux.HandleFunc("PUT /v1/boxes/:id", app.healthCheck)
	mux.HandleFunc("DELETE /v1/boxes/:id", app.healthCheck)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting http server", "addr", srv.Addr, "env", cfg.env, "api_version", version)
	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
