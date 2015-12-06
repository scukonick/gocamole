package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	Pools []Pool `json:"pools"`
}

func NewConfig(r io.Reader) *Config {
	config := new(Config)
	decoder := json.NewDecoder(r)
	err := decoder.Decode(config)
	if err != nil {
		log.Fatalf("Could not parse config: %v", err)
	}
	return config
}

func NewConfigFromFile(filename string) *Config {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open config file: %v", err)
	}
	return NewConfig(file)
}
