package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	URLs     []string
	Interval time.Duration
	Timeout  time.Duration
	Port     string
}

func LoadConfig() (Config, error) {
	var cfg Config

	// TARGET_URLS is required - check .env.example
	raw := os.Getenv("TARGET_URLS")
	for _, u := range strings.Split(raw, ",") {
		if u = strings.TrimSpace(u); u != "" {
			cfg.URLs = append(cfg.URLs, u)
		}
	}
	if len(cfg.URLs) == 0 {
		return cfg, errors.New("TARGET_URLS is required (comma-separated list of URLs)")
	}

	var err error
	if cfg.Interval, err = durationEnv("CHECK_INTERVAL", 30*time.Second); err != nil {
		return cfg, err
	}
	if cfg.Timeout, err = durationEnv("CHECK_TIMEOUT", 5*time.Second); err != nil {
		return cfg, err
	}

	cfg.Port = getEnv("PORT", "8080")
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func durationEnv(key string, fallback time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("invalid %s %q: %w", key, v, err)
	}
	return d, nil
}