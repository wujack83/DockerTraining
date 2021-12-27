package internal

import (
	"os"
	"gitlab.com/andersph/docker-master-class/api/internal/config"
)

type Employee struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Mail string `json:"email"`
	DateOfBirth string `json:"birth"`
	Department string `json:"department"`
	JobTitle string `json:"job_title"`
}

var configFile = "config.yaml"

func LoadConfig () config.Config {

	var c config.Config
	c.GetConfig(configFile)
	return c
}

func GetEnv(key, fallback string) string {
	
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	} else {
		return val
	}
}