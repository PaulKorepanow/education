package main

import (
	"bookLibrary/internal/app"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
}

func main() {
	flag.Parse()

	config := app.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := app.InitStore(config.Store)
	if err != nil {
		log.Fatal(err)
	}

	s := app.NewServer(config, db)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
