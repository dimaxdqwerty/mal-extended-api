package models

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

var (
	dbPort = os.Getenv("DB_PORT")
	dbHost = os.Getenv("DB_HOST")
)

func init() {
	if dbPort == "" {
		dbPort = "6379"
	}
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
}

type Config struct {
	MalClientID string
}

var conf Config

func GetMalClientID() string {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	return conf.MalClientID
}

func GetDBPort() string {
	return dbPort
}

func GetDBHost() string {
	return dbHost
}
