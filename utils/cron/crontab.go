package cron

import (
	"strconv"
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
	day, _ := strconv.Atoi(config.Config("DASHBOARD_INTERVAL_DAY"))
	_, _ = s.Every(day).Day().At(config.Config("DASHBOARD_INTERVAL_TIME")).Do(amountTask)
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
