package main

import (
	"github.com/elton/cerp-api/api/controllers"
	"github.com/elton/cerp-api/config"
)

func main() {
	var server = controllers.Server{}

	server.Run(config.Config("SERVER_PORT"))
}
