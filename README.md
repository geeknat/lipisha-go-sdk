# Lipisha Payments Go SDK

This package provides bindings for the Lipisha Payments API (https://developer.lipisha.com/)
 
 - Free software : MIT license
 - Documentation : https://developer.lipisha.com

### Features
 - Get account balance
 - Get float balance
 - Request money
 - Send money
 - Send airtime
 - Authorize Card transaction
 - Complete Card transaction
 - Reverse Card transaction
 - Request settlement
 - Authorize settlement
 - Cancel settlement
 - Acknowledge a transaction
 - Reconcile a transaction
 - Reverse a transaction
 - Get transactions
 - Get customers
 - Create a user
 - Update a user
 - Delete a user
 - Get users


### Install
To install, use `go get`:

```bash
$ go get github.com/geeknat/lipisha-go-sdk/lipisha
```


### Usage
Get your API Key and API Signature from Lipisha dashboard.

Create a Lipisha object with the following parameters
 
1. Your API Key,

2. Your API Signature,

3. A boolean stating whether to use Sandbox (false) or Live (true) accounts,

4. A boolean flag to print out logs when debug mode is true. 

```go

package test

import (
	"github.com/geeknat/lipisha-go-sdk/lipisha"
)

func handleLipisha(){
	
	lipishaApp := lipisha.Lipisha{
    		APIKey:       "YOUR_API_KEY",
    		APISignature: "YOUR_API_SIGNATURE",
    		IsProduction: true,
    		Debug:        true}
	
}
		
```


### Examples


##### Sample ITN implementation

Here's a sample IPN implementation in Go.

Depending on your server implementation, have a POST route to the ITN handler

```go
apiRoutes.Post("/itn", a.ITN)
```

```go

package itn

import (
	"github.com/geeknat/lipisha-go-sdk/lipisha"
    ...
)

type IPNAcknowledgeResponse struct {
	ApiKey                       string `json:"api_key"`
	ApiSignature                 string `json:"api_signature"`
	ApiVersion                   string `json:"api_version"`
	ApiType                      string `json:"api_type"`
	TransactionStatus            string `json:"transaction_status"`
	TransactionReference         string `json:"transaction_reference"`
	TransactionStatusCode        string `json:"transaction_status_code"`
	TransactionStatusDescription string `json:"transaction_status_description"`
	TransactionStatusReason      string `json:"transaction_status_reason"`
	TransactionStatusAction      string `json:"transaction_status_action"`
	TransactionCustomSMS         string `json:"transaction_custom_sms"`
}

// I'm using the iris framework, so some functions may vary.

// Lipisha functions remain the same regardless

func (a *App) ITN(ctx context.Context) {

	log.Println("ITN.......")
	log.Println(ctx.FormValues())

	apiKey := ctx.PostValue("api_key")
	apiSignature := ctx.PostValue("api_signature")

	// Check if Lipisha ITN server has made a callback
	// Confirm if authentication details are genuine else log a fraud attempt
	if apiKey == config.LipishaAPIKey && apiSignature == config.LipishaAPISignature {

		// Process Initiate
		if ctx.PostValue("api_type") == "Initiate" {

			// Extract transaction details
			country := ctx.PostValue("transaction_country")
			transactionType := ctx.PostValue("transaction_type")
			method := ctx.PostValue("transaction_method")
			date := ctx.PostValue("transaction_date")
			currency := ctx.PostValue("transaction_currency")
			amount := ctx.PostValue("transaction_amount")
			name := ctx.PostValue("transaction_name")
			mobile := ctx.PostValue("transaction_mobile")
			email := ctx.PostValue("transaction_email")
			paybill := ctx.PostValue("transaction_paybill")
			paybillType := ctx.PostValue("transaction_paybill_type")
			accountNumber := ctx.PostValue("transaction_account_number")
			accountName := ctx.PostValue("transaction_account_name")
			reference := ctx.PostValue("transaction_reference")
			merchantReference := ctx.PostValue("transaction_merchant_reference")
			code := ctx.PostValue("transaction_code")
			status := ctx.PostValue("transaction_status")

			// Perform action e.g update order/invoice as paid, save to database or log

			ipnTransaction := payments.IPNPayment{
				Country:           country,
				TransactionType:   transactionType,
				Method:            method,
				DateTime:          date,
				Currency:          currency,
				Amount:            amount,
				Name:              name,
				Mobile:            mobile,
				Email:             email,
				Paybill:           paybill,
				PaybillType:       paybillType,
				AccountNumber:     accountNumber,
				AccountName:       accountName,
				Reference:         reference,
				MerchantReference: merchantReference,
				Code:              code,
				Status:            status}

			if err := ipnTransaction.InsertTransaction(a.DB); err != nil {
				fmt.Println(err)
				return
			}

			smsMessage := "Dear " + name + ", your payment of  " + currency + " " + amount + " via " + code + " was received."

			// Acknowledge API call response
			acknowledgeResponse := &IPNAcknowledgeResponse{
				ApiKey:                       apiKey,
				ApiSignature:                 apiSignature,
				ApiType:                      "Receipt",
				ApiVersion:                   ctx.PostValue("api_version"),
				TransactionReference:         reference,
				TransactionStatusCode:        "001",
				TransactionStatusDescription: "Transaction received successfully",
				TransactionStatusAction:      "ACCEPT",
				TransactionStatusReason:      "VALID_TRANSACTION",
				TransactionCustomSMS:         smsMessage,
				TransactionStatus:            "SUCCESS"}

			serverResponse, _ := json.Marshal(acknowledgeResponse)

			fmt.Println(serverResponse)

			w.Header().Set("Content-Type", "application/json")
			w.Write(serverResponse)

		}

		if ctx.PostValue("api_type") == "Acknowledge" {

			reference := ctx.PostValue("transaction_reference")

			fmt.Println("Ack ref , " + reference)

			ipnTransaction := payments.IPNPayment{Reference: reference}

			if err := ipnTransaction.GetTransactionByReference(a.DB); err != nil {
				fmt.Println(err)
				return
			}

		
			//PROCESS THE TRANSACTION
			
		}

	} else {
		log.Println("FRAUD ATTEMPT")
	}

}


```


##### Sample Methods
More methods can be found by calling the Lipisha object.

Get account balance :

```go

        response, err := lipishaApp.GetAccountBalance()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(response)
```

Get float balance :

```go
        accountNumber := 12345

	response, err := lipishaApp.GetAccountFloat(accountNumber)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
```

Request money :

```go

		merchantAccNumber := 15373
		mobileNumber := 254712345678
		currency := "KES"
		amount := 10
		
		// Your unique identifier for this transaction, it will be sent to your IPN
		merchantReference := "12"

		serverResponse, err := a.Lipisha.RequestMoney(
			merchantAccNumber,
			mobileNumber,
			amount,
			"Paybill (M-Pesa)",
			currency,
			merchantReference)

		if err != nil {
			fmt.Println(err)
			response.RespondWithError(ctx.ResponseWriter(), http.StatusOK, "We encountered an error")
			return
		}

		fmt.Println(serverResponse)

		var responseMap map[string]*json.RawMessage

		if err := json.Unmarshal([]byte(serverResponse), &responseMap); err != nil {
			fmt.Println(err)
			response.RespondWithError(ctx.ResponseWriter(), http.StatusOK, "We encountered an error")
			return
		}

		status, err := utils.GetValueByUnmarshalToInterface("status", responseMap["status"])
		if err != nil {
			fmt.Println(err)
			response.RespondWithError(ctx.ResponseWriter(), http.StatusOK, "We encountered an error")
			return
		}

		if status == "SUCCESS" {

			response.RespondWithJSON(ctx.ResponseWriter(), http.StatusOK, config.CodeSuccess, "Success")
			return

		}

		response.RespondWithError(ctx.ResponseWriter(), http.StatusOK, "We encountered an error")


```

Send money :

```go
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
```

### Licence

[MIT](https://choosealicense.com/licenses/mit/)

### Author

[Geek Nat](http://geeknat.com)

