# Lipisha Payments Go SDK

This package provides bindings for the Lipisha Payments API (http://developer.lipisha.com/)
 
 - Free software : MIT license
 - Documentation : http://developer.lipisha.com

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
$ go get -d github.com/geeknat/lipisha-go-sdk/lipisha
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

package ipn

import (
	"github.com/geeknat/lipisha-go-sdk/lipisha"
	"github.com/kataras/iris/context"
	"encoding/json"
	"strconv"
	"strings"
	"log"
)

type IPNAcknowledgeResponse struct {
	ApiKey                       string
	ApiSignature                 string
	ApiVersion                   string
	ApiType                      string
	TransactionStatus            string
	TransactionReference         string
	TransactionStatusCode        string
	TransactionStatusDescription string
	TransactionStatusReason      string
	TransactionStatusAction      string
	TransactionCustomSMS         string
}


// I'm using the iris framework, so some functions may vary.

// Lipisha functions remain the same regardless

func (a *App) ITN(ctx context.Context) {

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

			w := ctx.ResponseWriter()

			// Acknowledge API call
			acknowledgeResponse := IPNAcknowledgeResponse{
				ApiKey:                       apiKey,
				ApiSignature:                 apiSignature,
				ApiType:                      "Receipt",
				ApiVersion:                   ctx.PostValue("api_version"),
				TransactionReference:         reference,
				TransactionStatusCode:        "001",
				TransactionStatusDescription: "Transaction received successfully",
				TransactionStatusAction:      "ACCEPT",
				TransactionStatusReason:      "VALID_TRANSACTION",
				TransactionCustomSMS:         "Dear " + name + ", your payment of  " + currency + " " + amount + " via " + code + " was received.",
				TransactionStatus:            "SUCCESS"}

			serverResponse, _ := json.Marshal(acknowledgeResponse)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-Auth-Token, X-FILENAME, X-FILESIZE, Content-Disposition")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Write(serverResponse)

		}

		if ctx.PostValue("api_type") == "Acknowledge" {

			reference := ctx.PostValue("transaction_reference")

			ipnTransaction := payments.IPNPayment{Reference: reference}

			if err := ipnTransaction.GetTransactionByReference(a.DB); err != nil {
				fmt.Println(err)
				return
			}

			// Complete the transaction

			fmt.Println("Success")
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
        accountNumber := "12345"

	response, err := lipishaApp.GetAccountFloat(accountNumber)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
```

Request money :

```go

	accountNumber := "12345"
    	mobileNumber := "0718353279"
    	method := "Paybill (M-Pesa)"
    	currency := "KES"
    	amount := "1000"
    
    	// Your unique identifier for this transaction, it will be sent to your IPN
    	merchantReference := "1"
    
    	response, err := lipishaApp.RequestMoney(
    		accountNumber,
    		mobileNumber,
    		method,
    		amount,
    		currency,
    		merchantReference)
    
    	if err != nil {
    		fmt.Println(err)
    	}
    
    	fmt.Println(response)

```

Send money :

```go
        accountNumber := "15189"
	mobileNumber := "0718353279"
	currency := "KES"
	amount := "1000"

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

