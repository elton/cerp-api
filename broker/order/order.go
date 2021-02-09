package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/elton/cerp-api/models"
	"github.com/elton/cerp-api/utils/signatures"
	"github.com/go-acme/lego/v3/log"
	"github.com/joho/godotenv"
)

// A Response struct to map the Entity Response
type Response struct {
	Success    bool       `json:"success"`
	ErrorDesc  string     `json:"errorDesc"`
	Total      int        `json:"total"`
	Orders     []Order    `json:"orders"`
	Deliveries []Delivery `json:"deliverys"`
}

// A Order struct to map the Entity Order
type Order struct {
	Code                 string     `json:"code"`
	PlatformCode         string     `json:"platform_code"`
	OrderTypeName        string     `json:"order_type_name"`
	ShopName             string     `json:"shop_name"`
	ShopCode             string     `json:"shop_code"`
	VIPName              string     `json:"vip_name"`
	VIPCode              string     `json:"vip_code"`
	VIPRealName          string     `json:"vipRealName"`
	BusinessMan          string     `json:"business_man"`
	Qty                  int8       `json:"qty"`
	Amount               float64    `json:"amount"`
	Payment              float64    `json:"payment"`
	WarehouseName        string     `json:"warehouse_name"`
	WarehouseCode        string     `json:"warehouse_code"`
	DeliveryState        int8       `json:"delivery_state"`
	ExpressName          string     `json:"express_name"`
	ExpressCode          string     `json:"express_code"`
	ReceiverArea         string     `json:"receiver_area"`
	PlatformTradingState string     `json:"platform_trading_state"`
	Deliveries           []Delivery `json:"deliverys"`
	Details              []Detail   `json:"details"`
	Payments             []Payment  `json:"payments"`
	PayTime              string     `json:"paytime"`
	DealTime             string     `json:"dealtime"`
	CreateTime           string     `json:"createtime"`
	ModifyTime           string     `json:"modifytime"`
}

// Delivery struct to map the Entity of the Delivery.
type Delivery struct {
	Delivery      bool   `json:"delivery"`
	Code          string `json:"code"`
	WarehouseName string `json:"warehouse_name"`
	WarehouseCode string `json:"warehouse_code"`
	ExpressName   string `json:"express_name"`
	ExpressCode   string `json:"express_code"`
	MailNo        string `json:"mail_no"`
}

// Detail struct to map the Entity of the item details.
type Detail struct {
	OID              string  `json:"oid"`
	Qty              float64 `json:"qty"`
	Price            float64 `json:"price"`
	Amount           float64 `json:"amount"`
	Refund           int     `json:"refund"`
	Note             string  `json:"note"`
	PlatformItemName string  `json:"platform_item_name"`
	PlatformSkuName  string  `json:"platform_sku_name"`
	ItemCode         string  `json:"item_code"`
	ItemName         string  `json:"item_name"`
	ItemSimpleName   string  `json:"item_simple_name"`
	PostFee          float64 `json:"post_fee"`
	DiscountFee      float64 `json:"discount_fee"`
	AmountAfter      float64 `json:"amount_after"`
}

// Payment struct to map the Entity of the payment.
type Payment struct {
	Payment     float64 `json:"payment"`
	PayCode     string  `json:"payCode"`
	PayTypeName string  `json:"pay_type_name"`
	PayTime     string  `json:"payTime"`
}

// GetOrders returns a list of all orders form specified shop.
func GetOrders(pgNum string, pgSize string, shopCode string, startDate time.Time) (*[]models.Order, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
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
	request["start_date"] = startDate.Format("2006-01-02 15:04:05")
	request["date_type"] = 3

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

	log.Infof("Order request JSON:%s \n", string(reqJSON))

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

	log.Infof("Get %d order information. \n", responseObject.Total)

	orders := []models.Order{}
	var layout string = "2006-01-02 15:04:05"

	for i := 0; i < len(responseObject.Orders); i++ {
		order := models.Order{}
		order.ID = node.Generate().Int64()
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

		for j := 0; j < len(responseObject.Orders[i].Deliveries); j++ {
			delivery := models.Delivery{}
			delivery.ID = node.Generate().Int64()
			delivery.Delivery = responseObject.Orders[i].Deliveries[j].Delivery
			delivery.Code = responseObject.Orders[i].Deliveries[j].Code
			delivery.WarehouseName = responseObject.Orders[i].Deliveries[j].WarehouseName
			delivery.WarehouseCode = responseObject.Orders[i].Deliveries[j].WarehouseCode
			delivery.ExpressName = responseObject.Orders[i].Deliveries[j].ExpressName
			delivery.MailNo = responseObject.Orders[i].Deliveries[j].MailNo

			order.Deliveries = append(order.Deliveries, delivery)
		}

		for m := 0; m < len(responseObject.Orders[i].Details); m++ {
			detail := models.Detail{}
			detail.ID = node.Generate().Int64()
			detail.OID = responseObject.Orders[i].Details[m].OID
			detail.Qty = responseObject.Orders[i].Details[m].Qty
			detail.Price = responseObject.Orders[i].Details[m].Price
			detail.Amount = responseObject.Orders[i].Details[m].Amount
			detail.Refund = responseObject.Orders[i].Details[m].Refund
			detail.Note = responseObject.Orders[i].Details[m].Note
			detail.PlatformItemName = responseObject.Orders[i].Details[m].PlatformItemName
			detail.PlatformSkuName = responseObject.Orders[i].Details[m].PlatformSkuName
			detail.ItemCode = responseObject.Orders[i].Details[m].ItemCode
			detail.ItemName = responseObject.Orders[i].Details[m].ItemName
			detail.ItemSimpleName = responseObject.Orders[i].Details[m].ItemSimpleName
			detail.PostFee = responseObject.Orders[i].Details[m].PostFee
			detail.DiscountFee = responseObject.Orders[i].Details[m].DiscountFee
			detail.AmountAfter = responseObject.Orders[i].Details[m].AmountAfter

			order.Details = append(order.Details, detail)
		}

		for n := 0; n < len(responseObject.Orders[i].Payments); n++ {
			payment := models.Payment{}
			payment.ID = node.Generate().Int64()
			payment.Payment = responseObject.Orders[i].Payments[n].Payment
			payment.PayCode = responseObject.Orders[i].Payments[n].PayCode
			payment.PayTypeName = responseObject.Orders[i].Payments[n].PayTypeName

			if responseObject.Orders[i].Payments[n].PayTime != "" && responseObject.Orders[i].Payments[n].PayTime != "0000-00-00 00:00:00" {
				payment.PayTime, err = time.ParseInLocation(layout, responseObject.Orders[i].Payments[n].PayTime, time.Local)
				if err != nil {
					return nil, err
				}
			}

			order.Payments = append(order.Payments, payment)
		}

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
