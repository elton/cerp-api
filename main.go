package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elton/cerp-api/api/controllers"
	_ "github.com/elton/cerp-api/cron"
	_ "github.com/elton/cerp-api/models"
	// _ "github.com/elton/cerp-api/utils/batch"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values.")
	}

	var server = controllers.Server{}

	server.Run(os.Getenv("SERVER_PORT"))
}
