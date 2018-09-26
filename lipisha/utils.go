package lipisha

import (
	"net/http"
	"time"
	"strings"
	"strconv"
	"net/url"
	"log"
	"io/ioutil"
)

var netClient = &http.Client{
	Timeout: time.Second * 30,
}

func (app *Lipisha) getURLResponse(endPoint string, data url.Values) (string, error) {
	var apiUrl string

	if app.IsProduction {
		apiUrl = liveUrl
	} else {
		apiUrl = sandBoxUrl
	}

	data.Set("api_key", app.APIKey)
	data.Set("api_signature", app.APISignature)

	urlStr := apiUrl + endPoint

	if app.Debug {
		log.Println(urlStr)
	}

	if app.Debug {
		log.Println(data.Encode())
	}

	client := netClient
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)

	if err != nil {
		if app.Debug {
			log.Println(err)
		}
		return "", err
	}

	f, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		if app.Debug {
			log.Println(err)
		}
		return "", err
	}

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if app.Debug {
		log.Println(string(f))
	}

	return string(f), nil
}
