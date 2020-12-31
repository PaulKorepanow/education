package main

import (
	"bookLibrary/internal/server"
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

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := server.ConnectToDB(config.Store)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(db)
	if err := s.Start(config); err != nil {
		log.Fatal(err)
	}
}
