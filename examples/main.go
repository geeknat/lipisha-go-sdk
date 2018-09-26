package main

import (
	"github.com/geeknat/lipisha-go-sdk/lipisha"
	"fmt"
)

// I'm probably not feeling so good, so I won't write real test files, just do actual examples
// Replace the config options with your own
func main() {

	lipishaApp := lipisha.Lipisha{
		APIKey:       "",
		APISignature: "",
		IsProduction: true,
		Debug:        true}

	accountNumber := 15189
	mobileNumber := 254718353279
	currency := "KES"
	amount := 1000

	// Your unique identifier for this transaction, it will be sent to your IPN
	merchantReference := "1"

	response, err := lipishaApp.SendMoney(
		accountNumber,
		mobileNumber,
		amount,
		currency,
		merchantReference)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}
