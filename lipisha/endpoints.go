package lipisha

import (
	"net/url"
	"strconv"
)

const (
	getAccountBalance        = "/get_balance"
	getAccountFloat          = "/get_float"
	sendMoney                = "/send_money"
	requestMoney             = "/request_money"
	sendAirtime              = "/send_airtime"
	sendSMS                  = "/send_sms"
	authorizeCardTransaction = "/authorize_card_transaction"
	completeCardTransaction  = "/complete_card_transaction"
	reverseCardTransaction   = "/reverse_card_transaction"
	voidCardTransaction      = "/void_card_transaction"
	requestSettlement        = "/request_settlement"
	authorizeSettlement      = "/authorize_settlement"
	cancelSettlement         = "/cancel_settlement"
	acknowledgeTransaction   = "/acknowledge_transaction"
	reconcileTransaction     = "/reconcile_transaction"
	reverseTransaction       = "/reverse_transaction"
	getTransactions          = "/get_transactions"
	getCustomers             = "/get_customers"
	createUser               = "/create_user"
	updateUser               = "/update_user"
	deleteUser               = "/delete_user"
	getUsers                 = "/get_users"
)

// GetAccountBalance returns the balance in your main Lipisha account.
func (app *Lipisha) GetAccountBalance() (string, error) {
	return app.getURLResponse(getAccountBalance, url.Values{})
}

// GetAccountFloat returns the float in a given account number
func (app *Lipisha) GetAccountFloat(accountNumber int) (string, error) {

	data := url.Values{}
	data.Set("account_number", strconv.Itoa(accountNumber))

	return app.getURLResponse(getAccountFloat, data)
}

// RequestMoney initiates a direct debit from the mobile money wallet e.g M-pesa,
// of your customer into your account.
// The customer gets an alert on their phone asking them to enter their PIN to confirm payment.
func (app *Lipisha) RequestMoney(accountNumber, mobileNumber, amount int, method, currency, myReference string) (string, error) {

	data := url.Values{}
	data.Set("account_number", strconv.Itoa(accountNumber))
	data.Set("mobile_number", strconv.Itoa(mobileNumber))
	data.Set("amount", strconv.Itoa(amount))
	data.Set("method", method)
	data.Set("currency", currency)
	data.Set("reference", myReference)

	return app.getURLResponse(requestMoney, data)
}

// SendMoney credits the mobile money wallet e.g M-PESA, of your customer from the float in your payout account.
// The customer gets a mobile money SMS confirmation on their phone on receiving the money.
// Once completed, it sends an Instant Transaction Notification or webhook event to your callback URL for you to process the transaction.
func (app *Lipisha) SendMoney(accountNumber, mobileNumber, amount int, currency, myReference string) (string, error) {

	data := url.Values{}
	data.Set("account_number", strconv.Itoa(accountNumber))
	data.Set("mobile_number", strconv.Itoa(mobileNumber))
	data.Set("amount", strconv.Itoa(amount))
	data.Set("currency", currency)
	data.Set("reference", myReference)

	return app.getURLResponse(sendMoney, data)
}

// SendAirtime recharges or tops up the airtime credit on your customer's phone from the float in your airtime account.
// The customer gets an airtime topup SMS confirmation on their phone on recharge.
// Once completed, it sends an Instant Transaction Notification or webhook event to your callback URL for you to process the transaction.
func (app *Lipisha) SendAirtime(accountNumber, mobileNumber, amount int, currency, myReference string) (string, error) {

	data := url.Values{}
	data.Set("account_number", strconv.Itoa(accountNumber))
	data.Set("mobile_number", strconv.Itoa(mobileNumber))
	data.Set("amount", strconv.Itoa(amount))
	data.Set("currency", currency)
	data.Set("reference", myReference)

	return app.getURLResponse(sendAirtime, data)
}

// SendSMS sends a text message via SMS to your customer.
// Once completed, it sends an Instant Transaction Notification or webhook event to your callback URL for you to process the transaction.
func (app *Lipisha) SendSMS(accountNumber, mobileNumber int, message, myReference string) (string, error) {

	data := url.Values{}
	data.Set("account_number", strconv.Itoa(accountNumber))
	data.Set("mobile_number", strconv.Itoa(mobileNumber))
	data.Set("message", message)
	data.Set("reference", myReference)

	return app.getURLResponse(sendSMS, data)
}

// AuthorizeCardTransaction authorizes a credit card transaction locking in the specified amount in the card holder's bank account.
// The transaction then needs to be completed using the Complete Card Transaction API call to effect settlement of funds into the merchant's account
// or reversed using the Reverse Card Transaction API call.
// This function reserves funds on the cardholder's account and if successful then you must call the complete_card_transaction
// function with the transaction_index and transaction_reference returned by this
// function to actually move the money to your account.
// Kindly note that in some cases, debit card transactions may be settled before the Complete Card Transaction API call is completed and may NOT be reversible depending on the issuing bank.
func (app *Lipisha) AuthorizeCardTransaction(accountNumber, amount int, cardNumber, addressOne, addressTwo, expiry, cardHolderName, email, mobileNumber, country, state, zip, securityCode, currency string) (string, error) {

	data := url.Values{}
	data.Set("account_number", strconv.Itoa(accountNumber))
	data.Set("mobile_number", mobileNumber)
	data.Set("card_number", cardNumber)
	data.Set("address1", addressOne)
	data.Set("address2", addressTwo)
	data.Set("expiry", expiry)
	data.Set("name", cardHolderName)
	data.Set("email", email)
	data.Set("country", country)
	data.Set("state", state)
	data.Set("zip", zip)
	data.Set("amount", strconv.Itoa(amount))
	data.Set("security_code", securityCode)
	data.Set("currency", currency)

	return app.getURLResponse(authorizeCardTransaction, data)
}

// CompleteCardTransaction completes a credit card transaction and initiates settlement of funds from the cardholder bank account into the merchant's account.
// This function moves already reserved funds on the cardholder's account into your account.
// It's called with the transaction_index and transaction_reference returned by the authorize_card_transaction
// function to actually move the money to your account.
func (app *Lipisha) CompleteCardTransaction(transactionIndex, transactionReference string) (string, error) {

	data := url.Values{}
	data.Set("transaction_index", transactionIndex)
	data.Set("transaction_reference", transactionReference)

	return app.getURLResponse(completeCardTransaction, data)
}

// ReverseCardTransaction reverses an authorized credit card transaction.
// This function unreserves funds previously authorized.
// It's called with the transaction_index and transaction_reference returned by the authorize_card_transaction function to reverse the authorization.
func (app *Lipisha) ReverseCardTransaction(transactionIndex, transactionReference string) (string, error) {

	data := url.Values{}
	data.Set("transaction_index", transactionIndex)
	data.Set("transaction_reference", transactionReference)

	return app.getURLResponse(reverseCardTransaction, data)
}

// VoidCardTransaction cancels a card transaction for which funds have already been charged.
// It's called with the transaction_index and transaction_reference returned by the authorize_card_transaction function to reverse the authorization.
func (app *Lipisha) VoidCardTransaction(transactionIndex, transactionReference string) (string, error) {

	data := url.Values{}
	data.Set("transaction_index", transactionIndex)
	data.Set("transaction_reference", transactionReference)

	return app.getURLResponse(voidCardTransaction, data)
}

// RequestSettlement requests a settlement of the account balance to a transaction or withdrawal account.
func (app *Lipisha) RequestSettlement(toAccountNumber, amount int) (string, error) {

	data := url.Values{}
	data.Set("account_number", strconv.Itoa(toAccountNumber))
	data.Set("amount", strconv.Itoa(amount))

	return app.getURLResponse(requestSettlement, data)
}

// AuthorizeSettlement authorizes a settlement transaction that had been requested.
func (app *Lipisha) AuthorizeSettlement(transactionCode string) (string, error) {

	data := url.Values{}
	data.Set("transaction", transactionCode)

	return app.getURLResponse(authorizeSettlement, data)
}

// CancelSettlement cancels a settlement transaction that had been requested.
func (app *Lipisha) CancelSettlement(transactionCode string) (string, error) {

	data := url.Values{}
	data.Set("transaction", transactionCode)

	return app.getURLResponse(cancelSettlement, data)
}

// AcknowledgeTransaction flags a transaction as having been processed by your application.
// This API call is particularly useful when processing e-commerce transactions and validating a payment.
func (app *Lipisha) AcknowledgeTransaction(commaSeparatedTransactionCode string) (string, error) {

	data := url.Values{}
	data.Set("transaction", commaSeparatedTransactionCode)

	return app.getURLResponse(acknowledgeTransaction, data)
}

// ReconcileTransaction searches and reconciles transactions that have been sent with an invalid
// or missing account and moves it to the specified account from the reconciliations queue.
// If successful, it triggers the Instant Payment Notification process sending relevant SMS, email and IPN URL notification.
func (app *Lipisha) ReconcileTransaction(transactionCode, mobileNumber, accountNumber, transactionReference string) (string, error) {

	data := url.Values{}
	data.Set("transaction", transactionCode)
	data.Set("transaction_mobile_number", mobileNumber)
	data.Set("transaction_account_number", accountNumber)
	data.Set("transaction_reference", transactionReference)

	return app.getURLResponse(reconcileTransaction, data)
}

// ReverseTransaction reverses one or more transactions and refunds the payments back to the customers.
// The API call requires that there is enough balance in the account to meet the amount to be reversed.
func (app *Lipisha) ReverseTransaction(commaSeparatedTransactionCode string) (string, error) {

	data := url.Values{}
	data.Set("transaction", commaSeparatedTransactionCode)

	return app.getURLResponse(reverseTransaction, data)
}

// GetTransactions returns all transactions matching the specified query parameters.
func (app *Lipisha) GetTransactions(transactionCodeList, transactionType, transactionMethod, startDate, endDate, accountNamesList, accountNumbersList, transactionReferencesList, minimumAmount, maximumAmount, transactionStatus, mobileNumber, email, offset, limit string) (string, error) {

	data := url.Values{}
	data.Set("transaction", transactionCodeList)
	data.Set("transaction_type", transactionType)
	data.Set("transaction_method", transactionMethod)
	data.Set("transaction_date_start", startDate)
	data.Set("transaction_date_end", endDate)
	data.Set("transaction_account_name", accountNamesList)
	data.Set("transaction_account_number", accountNumbersList)
	data.Set("transaction_reference", transactionReferencesList)
	data.Set("transaction_amount_minimum", minimumAmount)
	data.Set("transaction_amount_maximum", maximumAmount)
	data.Set("transaction_status", transactionStatus)
	data.Set("transaction_mobile_number", mobileNumber)
	data.Set("transaction_email", email)
	data.Set("limit", limit)
	data.Set("offset", offset)

	return app.getURLResponse(getTransactions, data)
}

// GetCustomers returns all customers matching the specified query parameters.
func (app *Lipisha) GetCustomers(customerName, customerMobileNumber, customerEmail, firstPaymentStartDate, firstPaymentEndDate, lastPaymentStartDate, lastPaymentEndDate, minimumNumberOfPayments, maximumNumberOfPayments, minimumTotalSpent, maximumTotalSpent, minimumAverageSpent, maximumAverageSpent, offset, limit string) (string, error) {

	data := url.Values{}
	data.Set("customer_name", customerName)
	data.Set("customer_mobile_number", customerMobileNumber)
	data.Set("customer_email", customerEmail)
	data.Set("customer_first_payment_from", firstPaymentStartDate)
	data.Set("customer_first_payment_to", firstPaymentEndDate)
	data.Set("customer_last_payment_from", lastPaymentStartDate)
	data.Set("customer_last_payment_to", lastPaymentEndDate)
	data.Set("customer_payments_minimum", minimumNumberOfPayments)
	data.Set("customer_payments_maximum", maximumNumberOfPayments)
	data.Set("customer_total_spent_minimum", minimumTotalSpent)
	data.Set("customer_total_spent_maximum", maximumTotalSpent)
	data.Set("customer_average_spent_minimum", minimumAverageSpent)
	data.Set("customer_average_spent_maximum", maximumAverageSpent)
	data.Set("limit", limit)
	data.Set("offset", offset)

	return app.getURLResponse(getCustomers, data)
}

// CreateUser creates a user under your Account.
func (app *Lipisha) CreateUser(fullName, role, mobileNumber, email, userName, password string) (string, error) {

	data := url.Values{}
	data.Set("full_name", fullName)
	data.Set("role", role)
	data.Set("mobile_number", mobileNumber)
	data.Set("email", email)
	data.Set("user_name", userName)
	data.Set("password", password)

	return app.getURLResponse(createUser, data)
}

// UpdateUser updates a user.
func (app *Lipisha) UpdateUser(fullName, role, mobileNumber, email, userName, password string) (string, error) {

	data := url.Values{}
	data.Set("full_name", fullName)
	data.Set("role", role)
	data.Set("mobile_number", mobileNumber)
	data.Set("email", email)
	data.Set("user_name", userName)
	data.Set("password", password)

	return app.getURLResponse(updateUser, data)
}

// DeleteUser deletes a user.
func (app *Lipisha) DeleteUser(userName string) (string, error) {

	data := url.Values{}
	data.Set("user_name", userName)

	return app.getURLResponse(deleteUser, data)
}

// GetUsers updates a user.
func (app *Lipisha) GetUsers(fullName, role, mobileNumber, email, userName string) (string, error) {

	data := url.Values{}
	data.Set("full_name", fullName)
	data.Set("role", role)
	data.Set("mobile_number", mobileNumber)
	data.Set("email", email)
	data.Set("user_name", userName)

	return app.getURLResponse(getUsers, data)
}
