package main

import (
	"fmt"
	"github.com/mmaxim2710/url-shortener/internal/config"
	"github.com/mmaxim2710/url-shortener/internal/storage/postgeSQL"
	"github.com/mmaxim2710/url-shortener/pkg/logger"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.Env)

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting URL shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := postgeSQL.New(cfg.DB)
	if err != nil {
		log.Error("failed to initialize storage", logger.Err(err))
		os.Exit(1)
	}
	_ = storage

	log.Info("storage initialized")

	// router

	// start server
}
