package cron

import (
	"fmt"
	"log"

	"github.com/elton/cerp-api/broker/basic"
	"github.com/elton/cerp-api/broker/order"
	"github.com/elton/cerp-api/models"
	"github.com/robfig/cron"
)

func init() {
	c := cron.New()
	shop := models.Shop{}
	c.AddFunc("00 * * * * ?", func() {
		shops, err := basic.GetShops("1", "20")
		if err != nil {
			log.Fatal(err)
			return
		}

		shopCreated, err := shop.SaveAll(shops)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Save the shops %v\n", shopCreated)
	})

	c.AddFunc("00 * * * * ?", func() {
		orders, err := order.GetOrders("1", "20", "011")
		if err != nil {
			log.Fatal(err)
			return
		}

		order := models.Order{}
		orderCreated, err := order.SaveAll(orders)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Save the orders %v\n", orderCreated)
	})

	c.Start()
}
