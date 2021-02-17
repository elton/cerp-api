package cron

import (
	"time"

	"github.com/elton/cerp-api/config"
	"github.com/elton/cerp-api/models"
	"github.com/elton/cerp-api/utils/logger"

	"github.com/go-co-op/gocron"
	"github.com/golang-module/carbon"
)

// SyncData synchron all the data of shop and order.
func SyncData() {
	tl, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(tl)
	_, _ = s.Every(config.Config("DASHBOARD_INTERVAL")).Do(amountTask)
	s.StartAsync()
}

func amountTask() {
	var start, end string
	order := new(models.Order)
	shop := new(models.Shop)
	mons, _ := order.GetOrderCreatedMon()

	for _, mon := range mons {
		if mon == "0000-00" {
			continue
		}
		mon = mon + "-01"
		start = carbon.Parse(mon).StartOfMonth().ToDateTimeString() // 本月开始时间
		end = carbon.Parse(mon).EndOfMonth().ToDateTimeString()     // 本月结束时间
		_, err := shop.GetAmountByShop(start, end)
		if err != nil {
			logger.Error.Println(err)
			return
		}
	}
}
