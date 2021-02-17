package main

import (
	_ "github.com/dimiro1/banner/autoload"

	"github.com/elton/cerp-api/api/controllers"
	"github.com/elton/cerp-api/config"
	"github.com/elton/cerp-api/utils/cron"
)

func main() {
	cron.SyncData()
	var server = controllers.Server{}
	server.Run(config.Config("SERVER_PORT"))
}
