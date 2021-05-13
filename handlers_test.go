package main

import (
	"testing"
)

// Tests email validation
func TestIsEmail(t *testing.T) {

	var (
		vGot, vExpd bool
		e           string
	)

	// testing correct email #1
	e = "qwerty@gmail.com"
	vExpd = true

	vGot = isEmail(e)
	if vGot != vExpd {
		t.Error("Expected:", vExpd, "got:", vGot)
	}

	// testing correct email #2
	e = "qwer234@mail.ru"
	vExpd = true

	vGot = isEmail(e)
	if vGot != vExpd {
		t.Error("Expected:", vExpd, "got:", vGot)
	}

	// testing incorrect email #1
	e = "qweriopmail.ru"
	vExpd = false

	vGot = isEmail(e)
	if vGot != vExpd {
		t.Error("Expected:", vExpd, "got:", vGot)
	}

	// testing incorrect email #2
	e = "qwerty@proton,com"
	vExpd = false

	vGot = isEmail(e)
	if vGot != vExpd {
		t.Error("Expected:", vExpd, "got:", vGot)
	}
}

// Tests UUID validation
func TestIsUUID(t *testing.T) {

	var (
		vGot, vExpd bool
		u           string
	)

	// testing correct UUID #1
	u = "a5657a25-b62d-45f8-96f6-41aab04f9ec0"
	vExpd = true

	vGot = isUUID(u)
	if vGot != vExpd {
		t.Error("Expected:", vExpd, "got:", vGot)
	}

	// testing correct UUID #2
	u = "c30946e2-eb83-45fb-99e1-44e5ea1528ac"
	vExpd = true

	vGot = isUUID(u)
	if vGot != vExpd {
		t.Error("Expected:", vExpd, "got:", vGot)
	}

	// testing incorrect UUID #1
	u = "qweqw-342423-mttyu-45"
	vExpd = false

	vGot = isUUID(u)
	if vGot != vExpd {
		t.Error("Expected:", vExpd, "got:", vGot)
	}

	// testing incorrect UUID #2
	u = "76.34%ter&das32HEWfew"
	vExpd = false

	vGot = isUUID(u)
	if vGot != vExpd {
		t.Error("Expected:", vExpd, "got:", vGot)
	}
}
