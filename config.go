package main

import (
	"encoding/json"
	"log"
	"os"
)

// Loads config from file (config.json)
func loadCfg(path string) (tConfig, error) {

	var c tConfig

	file, errF := os.Open(path)
	if errF != nil {
		log.Println(errF)
		return c, errF
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	c = tConfig{}
	errDec := decoder.Decode(&c)
	if errDec != nil {
		log.Println(errDec)
		return c, errDec
	}

	return c, nil
}
