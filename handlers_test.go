package main

import (
	"testing"
)

// Tests email validation
func TestIsEmail(t *testing.T) {

	var (
		got, expd bool
		e         string
	)

	// testing correct email #1
	e = "qwerty@gmail.com"
	expd = true

	got = isEmail(e)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing correct email #2
	e = "qwer234@mail.ru"
	expd = true

	got = isEmail(e)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing incorrect email #1
	e = "qweriopmail.ru"
	expd = false

	got = isEmail(e)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing incorrect email #2
	e = "qwerty@proton,com"
	expd = false

	got = isEmail(e)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}
}

// Tests UUID validation
func TestIsUUID(t *testing.T) {

	var (
		got, expd bool
		u         string
	)

	// testing correct UUID #1
	u = "a5657a25-b62d-45f8-96f6-41aab04f9ec0"
	expd = true

	got = isUUID(u)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing correct UUID #2
	u = "c30946e2-eb83-45fb-99e1-44e5ea1528ac"
	expd = true

	got = isUUID(u)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing incorrect UUID #1
	u = "qweqw-342423-mttyu-45"
	expd = false

	got = isUUID(u)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}

	// testing incorrect UUID #2
	u = "76.34%ter&das32HEWfew"
	expd = false

	got = isUUID(u)
	if got != expd {
		t.Error("Expected:", expd, "got:", got)
	}
}
