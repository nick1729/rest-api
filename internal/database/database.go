package database

import (
	"database/sql"
	"fmt"
	"rest-api/config"
	"rest-api/internal/types"
)

// Dial opens DB and checks connection
func Dial() (*sql.DB, error) {

	var (
		c       types.DbConfig
		db      *sql.DB
		connStr string
		err     error
	)

	c, err = config.GetCfg()
	if err != nil {
		return nil, err
	}

	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Pass, c.DbName)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
