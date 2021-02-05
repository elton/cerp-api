package cron

import (
	"fmt"
	"log"

	"github.com/elton/cerp-api/broker/basic"
	"github.com/elton/cerp-api/models"
	"github.com/robfig/cron"
)

func init() {
	c := cron.New()
	shop := models.Shop{}
	c.AddFunc("@midnight", func() {
		shops := basic.GetShops("1", "20")
		shopCreated, err := shop.SaveAll(shops)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Save the shops %v\n", shopCreated)
	})
	c.Start()
}
