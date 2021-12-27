package main

import (
	"gitlab.com/andersph/docker-master-class/api/internal/server"
)

func main() {
	server.HandleRequests("config.yaml")
}
