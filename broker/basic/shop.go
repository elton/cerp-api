package basic

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
	Success   bool   `json:"success"`
	ErrorDesc string `json:"errorDesc"`
	Total     int    `json:"total"`
	Shops     []Shop `json:"shops"`
}

// A Shop struct to map every shop information.
type Shop struct {
	ID         int    `json:"id"`
	Nick       string `json:"nick"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	CreateDate string `json:"create_date"`
	ModifyDate string `json:"modify_date"`
	Note       string `json:"note"`
	TypeName   string `json:"type_name"`
}

// GetShops returns the list of shops.
func GetShops(pgNum string, pgSize string, startDate time.Time) (*[]models.Shop, error) {
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
	request["method"] = "gy.erp.shop.get"
	request["page_no"] = pgNum
	request["page_size"] = pgSize
	request["modify_start_date"] = startDate.Format("2006-01-02 15:04:05")

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

	log.Infof("Shop request JSON:%s \n", string(reqJSON))

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

	log.Infof("Get %d shop information. \n", responseObject.Total)

	var shops []models.Shop
	var shop models.Shop
	var layout string = "2006-01-02 15:04:05"

	for i := 0; i < len(responseObject.Shops); i++ {
		shop.ID = node.Generate().Int64()
		shop.ShopID = responseObject.Shops[i].ID
		shop.Name = responseObject.Shops[i].Name
		shop.Nick = responseObject.Shops[i].Nick
		shop.Code = responseObject.Shops[i].Code
		shop.Note = responseObject.Shops[i].Note
		shop.TypeName = responseObject.Shops[i].TypeName

		if responseObject.Shops[i].CreateDate != "" && responseObject.Shops[i].CreateDate != "0000-00-00 00:00:00" {
			shop.CreateDate, err = time.ParseInLocation(layout, responseObject.Shops[i].CreateDate, time.Local)
			if err != nil {
				return nil, err
			}
		}

		if responseObject.Shops[i].ModifyDate != "" && responseObject.Shops[i].ModifyDate != "0000-00-00 00:00:00" {
			shop.ModifyDate, err = time.ParseInLocation(layout, responseObject.Shops[i].ModifyDate, time.Local)
			if err != nil {
				return nil, err
			}
		}

		shops = append(shops, shop)
	}

	return &shops, nil
}
