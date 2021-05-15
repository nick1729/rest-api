package config

import (
	"encoding/json"
	"os"
	"rest-api/internal/types"
	"testing"
)

// Tests loading and decoding json config
func TestLoadConfigData(t *testing.T) {

	var c types.Config

	file, errF := os.Open("./config_test.json")
	if errF != nil {
		t.Error("Expected:", nil, "got:", errF)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	c = types.Config{}
	errDec := decoder.Decode(&c)
	if errDec != nil {
		t.Error("Expected:", nil, "got:", errDec)
	}
}
