package main

import (
	"net/mail"

	"github.com/google/uuid"
)

// Email validation
func isEmail(e string) bool {

	var err error

	_, err = mail.ParseAddress(e)
	if err != nil {
		return false
	}

	return true
}

// Email validation
func isUUID(u string) bool {

	var err error

	_, err = uuid.Parse(u)
	if err != nil {
		return false
	}

	return true
}
