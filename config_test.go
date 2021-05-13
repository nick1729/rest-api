package main

import (
	"encoding/json"
	"os"
	"testing"
)

// Tests loading and decoding json config
func TestLoadConfigData(t *testing.T) {

	var c tConfig

	file, errF := os.Open("./tests/cfgdec_test.json")
	if errF != nil {
		t.Error("Expected:", nil, "got:", errF)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	c = tConfig{}
	errDec := decoder.Decode(&c)
	if errDec != nil {
		t.Error("Expected:", nil, "got:", errDec)
	}
}
