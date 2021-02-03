package basic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	ID           string `json:"id"`
	Nick         string `json:"nick"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	CreateDate   string `json:"create_date"`
	ModifiedDate string `json:"modified_date"`
	Note         string `json:"note"`
	TypeName     string `json:"type_name"`
}

// GetShops returns the list of shops.
func GetShops(pgNum string, pgSize string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
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
		panic(err)
	}

	sign := signatures.Sign(string(reqJSON), os.Getenv("secret"))

	request["sign"] = sign

	reqJSON, err = json.Marshal(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reqJSON))

	var responseObject Response

	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(reqJSON))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseObject)

	// fmt.Println(string(responseData))

	fmt.Println(responseObject.Total)

	for i := 0; i < len(responseObject.Shops); i++ {
		fmt.Println(responseObject.Shops[i].Name)
	}
}
