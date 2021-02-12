package main

import (
	"os"

	"github.com/elton/cerp-api/api/controllers"
	"github.com/elton/cerp-api/cron"
	"github.com/elton/cerp-api/models"
	"github.com/elton/cerp-api/utils/batch"
	"github.com/elton/cerp-api/utils/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error.Printf("Error getting env, not comming through %v\n", err)
	} else {
		logger.Info.Println("We are getting the env values.")
	}

	var shop *models.Shop
	shops, err := shop.GetAllShops()
	if err != nil {
		logger.Error.Println(err)
	}

	if len(*shops) <= 0 {
		batch.InitializeData()
	} else {
		cron.SyncData()
	}

	var server = controllers.Server{}

	server.Run(os.Getenv("SERVER_PORT"))
}
