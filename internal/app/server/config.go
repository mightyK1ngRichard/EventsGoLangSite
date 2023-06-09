package server

import "github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/store"

type Config struct {
	BindAddr   string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
	SessionKey string `toml:"session_key"`
	Store      *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
