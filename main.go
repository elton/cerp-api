package main

import (
	"net/http"

	_ "github.com/elton/cerp-api/cron"
	_ "github.com/elton/cerp-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	r.Run()
}
