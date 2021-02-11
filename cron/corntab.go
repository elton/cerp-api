package cron

import (
	"github.com/elton/cerp-api/broker"
	"github.com/elton/cerp-api/models"
	"github.com/go-acme/lego/v3/log"
	"github.com/robfig/cron"
)

func init() {
	// Sync store information
	c := cron.New()	
	shop := models.Shop{}
	orderDb := models.Order{}

	c.AddFunc("00 * * * * ?", func() {
		lastUpdateAt, err := shop.GetLastUpdatedAt()
		if err != nil {
			log.Fatal(err)
			return
		}

		shops, err := broker.GetShops("1", "20", lastUpdateAt)
		if err != nil {
			log.Fatal(err)
			return
		}

		if len(*shops) == 0 {
			return
		}

		shopCreated, err := shop.SaveAll(shops)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Infof("Save %d shops information\n", len(*shopCreated))
	})

	// Sync order information.
	c.AddFunc("00 * * * * ?", func() {

		lastUpdateAt, err := orderDb.GetLastUpdatedAt()
		if err != nil {
			log.Fatal(err)
			return
		}
		orders, err := broker.GetOrders("1", "20", "011", lastUpdateAt)
		if err != nil {
			log.Fatal(err)
			return
		}

		if len(*orders) == 0 {
			return
		}

		orderCreated, err := orderDb.SaveAll(orders)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Infof("Save %d orders information\n", len(*orderCreated))
	})

	c.Start()
}
