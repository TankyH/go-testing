package main

import (
	"github.com/ru-rocker/go-testing/xrates"
	"log"
	"fmt"
)

func main() {
	req := xrates.RatesQuotedRequest{Base: "USD"}

	rates := xrates.Rates{Request: req}
	resp, err := xrates.GetCurrentRates(rates)
	if err != nil {
		log.Fatal(err)
		return
	}


	fmt.Printf("Base: %s\n", resp.Base)
	fmt.Printf("Date %s\n", resp.Date)
	fmt.Printf("AUD: %f\n", resp.Rates["AUD"])
	fmt.Printf("NZD: %f\n", resp.Rates["NZD"])
	fmt.Printf("IDR: %f\n", resp.Rates["IDR"])
	fmt.Printf("CAD: %f\n", resp.Rates["CAD"])
	fmt.Printf("SGD: %f\n", resp.Rates["SGD"])
}
