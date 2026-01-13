package main

import (
	"log/slog"
	"os"

	"github.com/s4bb4t/url-shortener/internal/config"
	"github.com/s4bb4t/url-shortener/internal/storage/postgres"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fatal(err)
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	_ = log

	storage, err := postgres.New(cfg.Postgres)
	if err != nil {
		fatal(err)
	}
	_ = storage
	// TODO: init server

	// TODO: run server
}

func fatal(err error) {
	println(err.Error())
	os.Exit(1)
}
