package lipisha

//Lipisha is a struct defining the main LIPISHA configurations.
type Lipisha struct {
	//APIKey available on https://lipisha.com
	APIKey string
	//APISignature available on https://lipisha.com
	APISignature string
	//IsProduction defines whether to use live account or sandbox account
	IsProduction bool
	//Debug defines whether to print out responses and error
	Debug bool
}



