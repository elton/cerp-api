package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/elton/cerp-api/models"
	"github.com/elton/cerp-api/utils/signatures"
	"github.com/joho/godotenv"
)

// A Response struct to map the Entity Response
type Response struct {
	Success   bool    `json:"success"`
	ErrorDesc string  `json:"errorDesc"`
	Total     int     `json:"total"`
	Orders    []Order `json:"orders"`
}

// A Order struct to map the Entity Order
type Order struct {
	Code                 string  `json:"code"`
	PlatformCode         string  `json:"platform_code"`
	OrderTypeName        string  `json:"order_type_name"`
	ShopName             string  `json:"shop_name"`
	ShopCode             string  `json:"shop_code"`
	VIPName              string  `json:"vip_name"`
	VIPCode              string  `json:"vip_code"`
	VIPRealName          string  `json:"vipRealName"`
	BusinessMan          string  `json:"business_man"`
	Qty                  int     `json:"qty"`
	Amount               float64 `json:"amount"`
	Payment              float64 `json:"payment"`
	WarehouseName        string  `json:"warehouse_name"`
	WarehouseCode        string  `json:"warehouse_code"`
	DeliveryState        int     `json:"delivery_state"`
	ExpressName          string  `json:"express_name"`
	ExpressCode          string  `json:"express_code"`
	ReceiverArea         string  `json:"receiver_area"`
	PlatformTradingState string  `json:"platform_trading_state"`
	PayTime              string  `json:"paytime"`
	DealTime             string  `json:"dealtime"`
	CreateTime           string  `json:"createtime"`
	ModifyTime           string  `json:"modifytime"`
}

// GetOrders returns a list of all orders form specified shop.
func GetOrders(pgNum string, pgSize string, shopCode string) (*[]models.Order, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	apiURL := os.Getenv("apiURL")

	request := make(map[string]interface{})

	request["appkey"] = os.Getenv("appKey")
	request["sessionkey"] = os.Getenv("sessionKey")
	request["method"] = "gy.erp.trade.get"
	request["page_no"] = pgNum
	request["page_size"] = pgSize
	request["shop_code"] = shopCode

	reqJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	sign := signatures.Sign(string(reqJSON), os.Getenv("secret"))

	request["sign"] = sign

	reqJSON, err = json.Marshal(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(reqJSON))

	var responseObject Response

	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(reqJSON))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(responseData, &responseObject)

	// fmt.Println(string(responseData))

	// fmt.Println(responseObject.Total)

	var orders []models.Order
	var order models.Order
	var layout string = "2006-01-02 15:04:05"

	for i := 0; i < len(responseObject.Orders); i++ {
		order.Code = responseObject.Orders[i].Code
		order.PlatformCode = responseObject.Orders[i].PlatformCode
		order.OrderTypeName = responseObject.Orders[i].OrderTypeName
		order.ShopName = responseObject.Orders[i].ShopName
		order.ShopCode = responseObject.Orders[i].ShopCode
		order.VIPName = responseObject.Orders[i].VIPName
		order.VIPCode = responseObject.Orders[i].VIPCode
		order.VIPRealName = responseObject.Orders[i].VIPRealName
		order.BusinessMan = responseObject.Orders[i].BusinessMan
		order.Qty = responseObject.Orders[i].Qty
		order.Amount = responseObject.Orders[i].Amount
		order.Payment = responseObject.Orders[i].Payment

		order.WarehouseName = responseObject.Orders[i].WarehouseName
		order.WarehouseCode = responseObject.Orders[i].WarehouseCode
		order.DeliveryState = responseObject.Orders[i].DeliveryState
		order.ExpressName = responseObject.Orders[i].ExpressName
		order.ExpressCode = responseObject.Orders[i].ExpressCode
		order.ReceiverArea = responseObject.Orders[i].ReceiverArea
		order.PlatformTradingState = responseObject.Orders[i].PlatformTradingState

		if responseObject.Orders[i].CreateTime != "" && responseObject.Orders[i].CreateTime != "0000-00-00 00:00:00" {
			order.CreateTime, err = time.ParseInLocation(layout, responseObject.Orders[i].CreateTime, time.Local)
			if err != nil {
				return nil, err
			}
		}
		if responseObject.Orders[i].ModifyTime != "" && responseObject.Orders[i].ModifyTime != "0000-00-00 00:00:00" {
			order.ModifyTime, err = time.ParseInLocation(layout, responseObject.Orders[i].ModifyTime, time.Local)
			if err != nil {
				return nil, err
			}
		}
		if responseObject.Orders[i].DealTime != "" && responseObject.Orders[i].DealTime != "0000-00-00 00:00:00" {
			order.DealTime, err = time.ParseInLocation(layout, responseObject.Orders[i].DealTime, time.Local)
			if err != nil {
				return nil, err
			}
		}
		if responseObject.Orders[i].PayTime != "" && responseObject.Orders[i].PayTime != "0000-00-00 00:00:00" {
			order.PayTime, err = time.ParseInLocation(layout, responseObject.Orders[i].PayTime, time.Local)
			if err != nil {
				return nil, err
			}
		}

		orders = append(orders, order)
	}

	return &orders, nil
}
