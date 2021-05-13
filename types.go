package main

import (
	"time"

	"github.com/google/uuid"
)

type tConfig struct {
	Login string `json:"db_login"`
	Pass  string `json:"db_password"`
	Ip    string `json:"db_ip"`
	Port  int    `json:"db_port"`
	Table string `json:"table_name"`
}

type tUser struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
	Email     string
	Age       uint64
	Created   time.Time
}
