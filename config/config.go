package config

import (
	"encoding/json"
	"log"
	"os"
	"rest-api/internal/types"
)

// LoadCfg gets and decode config data from file (config.json)
func LoadCfg(path string) (types.Config, error) {

	var c types.Config

	file, errF := os.Open(path)
	if errF != nil {
		log.Println(errF)
		return c, errF
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	c = types.Config{}
	errDec := decoder.Decode(&c)
	if errDec != nil {
		log.Println(errDec)
		return c, errDec
	}

	return c, nil
}
