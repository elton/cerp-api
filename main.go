package main

import (
	"net/http"

	"github.com/elton/cerp-api/consume/basic"
	_ "github.com/elton/cerp-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	basic.GetShops("1", "20")

	r.Run()
}
