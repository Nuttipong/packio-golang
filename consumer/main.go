package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func main() {
	product, err := CallGet("http://localhost:5000", 35)
	if err != nil {
		fmt.Println(err)
	}
	// productList := make([]Product, 0)
	// err = json.Unmarshal(data, &productList)
	// product := &Product{}
	// err = json.Unmarshal(data, product)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(product)
}

func CallGet(host string, productId int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s/api/products/%d", host, productId)
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req.Request.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	var product Product
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func callPost(ssoid string) error {
	payload, _ := json.Marshal(map[string]interface{}{
		"campaign_code":              "VIRTUAL_CONCERT_20211225",
		"vendor_code":                "TMH",
		"identifier_type":            "account",
		"identifier_number":          ssoid,
		"physical_identifier_number": nil,
		"receipt_require":            false,
		"first_charge":               false,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequest("POST",
		"https://subscription-b2b-api.trueid-preprod.net/subscription/v3/orders",
		bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "5b18e526f0463656f7c4329f4c312a8f419645d1833725841baffc47")
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	all, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(all)
	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New("Failed response")
	}
}
