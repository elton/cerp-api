package basic

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
	Success   bool   `json:"success"`
	ErrorDesc string `json:"errorDesc"`
	Total     int    `json:"total"`
	Shops     []Shop `json:"shops"`
}

// A Shop struct to map every shop information.
type Shop struct {
	ID         string `json:"id"`
	Nick       string `json:"nick"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	CreateDate string `json:"create_date"`
	ModifyDate string `json:"modify_date"`
	Note       string `json:"note"`
	TypeName   string `json:"type_name"`
}

// GetShops returns the list of shops.
func GetShops(pgNum string, pgSize string) (*[]models.Shop, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	apiURL := os.Getenv("apiURL")

	request := make(map[string]interface{})

	request["appkey"] = os.Getenv("appKey")
	request["sessionkey"] = os.Getenv("sessionKey")
	request["method"] = "gy.erp.shop.get"
	request["page_no"] = pgNum
	request["page_size"] = pgSize

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

	fmt.Println(responseObject.Total)

	var shops []models.Shop
	var shop models.Shop
	var layout string = "2006-01-02 15:04:05"

	for i := 0; i < len(responseObject.Shops); i++ {
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
