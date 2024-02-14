package main

import (
	"fmt"
	"github.com/mmaxim2710/url-shortener/internal/config"
	"github.com/mmaxim2710/url-shortener/pkg/logger"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.Env)

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting URL shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// storage

	// router

	// start server
}
