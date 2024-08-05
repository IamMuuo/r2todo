package configs

import "testing"

func TestLoadConfig(t *testing.T) {
	var config Config

	if err := LoadConfig(&config); err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}
