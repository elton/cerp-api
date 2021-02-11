package batch

import (
	"os"
	"strconv"
	"time"

	"github.com/elton/cerp-api/broker"
	"github.com/elton/cerp-api/models"
	"github.com/go-acme/lego/v3/log"
	"github.com/joho/godotenv"
)

func init() {

	shops, err := getShops()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, shop := range *shops {
		if err := getOrders(shop.Code); err != nil {
			log.Fatal(err)
			return
		}
	}
}

// getShops save all shop information.
func getShops() (*[]models.Shop, error) {
	var (
		shop               models.Shop
		shops, shopCreated *[]models.Shop
		lastUpdateAt       time.Time
		err                error
	)

	if lastUpdateAt, err = shop.GetLastUpdatedAt(); err != nil {
		return nil, err
	}

	if shops, err = broker.GetShops("1", "20", lastUpdateAt); err != nil {
		return nil, err
	}

	if len(*shops) > 0 {
		if shopCreated, err = shop.SaveAll(shops); err != nil {
			return nil, err
		}
		log.Infof("Save %d shops information\n", len(*shopCreated))
	}

	return shops, nil
}

// getOrders save all the orders of specified shop.
func getOrders(shopCode string) error {
	var (
		orderDb              models.Order
		orders, orderCreated *[]models.Order
		lastUpdateAt         time.Time
		totalOrder           int
		err                  error
	)
	godotenv.Load()
	pgSize, _ := strconv.Atoi(os.Getenv("PAGE_SIZE"))

	if lastUpdateAt, err = orderDb.GetLastUpdatedAt(shopCode); err != nil {
		return err
	}

	if totalOrder, err = broker.GetTotalOfOrders(shopCode, lastUpdateAt); err != nil {
		return err
	}

	totalPg := totalOrder / pgSize
	if totalOrder%pgSize != 0 {
		totalPg = totalPg + 1
	}

	log.Infof("Total Order: %d, page size: %d, total page: %d", totalOrder, pgSize, totalPg)

	for i := 0; i < totalPg; i++ {
		if orders, err = broker.GetOrders(strconv.Itoa(i+1), strconv.Itoa(pgSize), shopCode, lastUpdateAt); err != nil {
			return err
		}

		if len(*orders) > 0 {
			if orderCreated, err = orderDb.SaveAll(orders); err != nil {
				return err
			}
			log.Infof("Save %d orders information\n", len(*orderCreated))
		}
	}
	return nil
}
