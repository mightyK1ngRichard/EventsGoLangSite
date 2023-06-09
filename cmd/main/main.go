package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatalln(err)
	}

	if err := apiserver.New(config).Start(); err != nil {
		log.Fatal(err)
	}
}
