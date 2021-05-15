package config

import (
	"encoding/json"
	"log"
	"os"
	t "rest-api/internal/types"
)

// Loads config from file (config.json)
func LoadCfg(path string) (t.Config, error) {

	var c t.Config

	file, errF := os.Open(path)
	if errF != nil {
		log.Println(errF)
		return c, errF
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	c = t.Config{}
	errDec := decoder.Decode(&c)
	if errDec != nil {
		log.Println(errDec)
		return c, errDec
	}

	return c, nil
}
