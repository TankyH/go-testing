package xrates

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type RatesQuotedRequest struct {
	Base string `json:"base"`
}

type RatesQuotedResponse struct {
	Base string `json:"base"`
	Date string `json:"date"`
	Rates map[string] float32 `json:"rates"`
}

type AccessRates interface {

	getRates() (RatesQuotedResponse,error)
}

type Rates struct {
	Request RatesQuotedRequest
}

func (r Rates) getRates() (RatesQuotedResponse, error)  {

	// Access the API
	url := fmt.Sprintf("http://api.fixer.io/latest?base=%s", r.Request.Base)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return RatesQuotedResponse{},err
	}

	// Create an HTTP Client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return RatesQuotedResponse{},err
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	var rqResp RatesQuotedResponse;
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&rqResp); err != nil {
		log.Println(err)
		return RatesQuotedResponse{},err
	}

	return rqResp, nil
}

func GetCurrentRates(rates AccessRates) (RatesQuotedResponse, error) {
	return rates.GetRates()
}