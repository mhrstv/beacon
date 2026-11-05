package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		slog.Error("invalid configuration", "err", err)
		os.Exit(1)
	}
	slog.Info("config loaded",
		"urls", cfg.URLs,
		"interval", cfg.Interval,
		"timeout", cfg.Timeout,
		"port", cfg.Port,
	)
}