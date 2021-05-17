package config

import (
	"log"
	"os"
	"rest-api/internal/types"
	"strconv"
)

// GetCfg parses and checks config data
func GetCfg() (types.DbConfig, error) {

	var (
		c      types.DbConfig
		port   string
		dbPort int
		err    error
	)

	port = os.Getenv("POSTGRES_PORT")
	dbPort, err = strconv.Atoi(port)
	if err != nil {
		log.Print("Failed to parse database port")
		return c, err
	}

	c.User = os.Getenv("POSTGRES_USER")
	c.Pass = os.Getenv("POSTGRES_PASSWORD")
	c.Host = os.Getenv("POSTGRES_HOST")
	c.Port = dbPort
	c.DbName = os.Getenv("POSTGRES_DB")

	return c, nil
}
