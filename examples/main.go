package main

import (
	"github.com/geeknat/lipisha-go-sdk/lipisha"
	"fmt"
	"encoding/json"
	"log"
)

// I'm probably not feeling so good, so I won't write real test files, just do actual examples
// Replace the config options with your own
func main() {

	lipishaApp := lipisha.Lipisha{
		APIKey:       "",
		APISignature: "",
		IsProduction: true,
		Debug:        true}

	accountNumber := 123456
	mobileNumber := 254712345678
	currency := "KES"
	amount := 1000

	// Your unique identifier for this transaction, it will be sent to your IPN
	merchantReference := "1"

	serverResponse, err := lipishaApp.RequestMoney(
		accountNumber,
		mobileNumber,
		amount,
		"Paybill (M-Pesa)",
		currency,
		merchantReference)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(serverResponse)

	var responseMap map[string]*json.RawMessage

	if err := json.Unmarshal([]byte(serverResponse), &responseMap); err != nil {
		log.Fatal(err)
		return
	}

	status, err := getValueByUnmarshalToInterface("status", responseMap["status"])
	if err != nil {
		log.Fatal(err)
		return
	}

	if status == "SUCCESS" {

		return

	}

}

func getValueByUnmarshalToInterface(key string, foo *json.RawMessage) (string, error) {
	var tmp map[string]interface{}
	if err := json.Unmarshal(*foo, &tmp); err != nil {
		return "", err
	}

	if typ, ok := tmp[key].(string); ok {
		return typ, nil
	} else {
		return "", fmt.Errorf(key + " should be a string")
	}
}
