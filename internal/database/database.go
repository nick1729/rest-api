package database

import (
	"database/sql"
	"fmt"
	"rest-api/config"
	"rest-api/internal/types"

	_ "github.com/lib/pq"
)

// Open DB and check connection
func Dial() (*sql.DB, error) {

	var (
		cfg     types.Config
		db      *sql.DB
		connStr string
		err     error
	)

	cfg, err = config.LoadCfg("../../config/config.json")
	if err != nil {
		return nil, err
	}

	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Ip, cfg.Port, cfg.Login, cfg.Pass, cfg.Table)

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
