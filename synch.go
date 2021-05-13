package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	model "allyapps.com/bequant/models"	
	db "allyapps.com/bequant/db"
	pg "allyapps.com/bequant/data/postgres"
)

func main(){
	
	// Connect to database
	db.Init()
	
	// Run every 2 min
	cronJob(120000*time.Millisecond, runSych)


}
func cronJob(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func runSych(t time.Time) {
	fmt.Printf("%v: Synching data!\n", t)
	// Parse Data
	data := parseData("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=BTC,LINK&tsyms=USD,EUR")
	fmt.Println(data.Raw.BTC.USD.CHANGE24HOUR)

	pg.ClearRaw()
	pg.ClearDisplay()
	
	// Insert BTC to database 
	pg.AddRawBtcUsdPrice(data)
	pg.AddRawBtcEurPrice(data)
	pg.AddDisplayBtcUsdPrice(data)
	pg.AddDisplayBtcEurPrice(data)

	// Insert LINK to database 
	pg.AddRawLinkUsdPrice(data)
	pg.AddRawLinkEurPrice(data)
	pg.AddDisplayLinkUsdPrice(data)
	pg.AddDisplayLinkEurPrice(data)

}
func parseData(url string) model.Data {
	var data model.Data

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return data
}