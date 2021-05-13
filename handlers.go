package main

import "net/mail"

// Email validation
func valid(e string) bool {

	var err error

	_, err = mail.ParseAddress(e)
	if err != nil {
		return false
	}

	return true
}
