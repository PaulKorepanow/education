package main

import (
	"bookLibrary/internal/server"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
	os.Setenv("db_log_level", "0")
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
