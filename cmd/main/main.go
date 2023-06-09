package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/sessions"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/server"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := server.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatalln(err)
	}

	session := sessions.NewCookieStore([]byte(config.SessionKey))
	if err := server.New(config, session).Start(); err != nil {
		log.Fatal(err)
	}
}
