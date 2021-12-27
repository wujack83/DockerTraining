package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	
	var c Config
	c.GetConfig("config_test.yaml")

	strConfig := fmt.Sprintf("%v", c)

	expected := `{{8080} {/employee} {technical_user postgres 5432 my_company} {true}}`
	if strConfig != expected {
		t.Errorf("Config not loaded correctly: got %v want %v",
		strConfig, expected)
	}
}