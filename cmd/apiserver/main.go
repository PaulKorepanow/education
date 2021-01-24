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
	logPath    string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
	if err := os.Setenv("token_password", "secret"); err != nil {
		panic(err)
	}

	if err := os.Setenv("log_path", "./logs/"); err != nil {
		panic(err)
	}

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
