package postgres

import(
	model "allyapps.com/bequant/models"
	"strings"
	"fmt"
	db "allyapps.com/bequant/db"
)

func GetPrice(fsyms string, tsyms string) model.Data{
	var data model.Data
	fsysmsRaw := &model.FsysmsRaw{}
	fsysmsDisplay := &model.FsysmsDisplay{}
	currenciesRawBtc := &model.CurrenciesRaw{}
	currenciesDisplayBtc := &model.CurrenciesDisplay{}
	currenciesRawLink := &model.CurrenciesRaw{}
	currenciesDisplayLink := &model.CurrenciesDisplay{}

	if fsyms == "" {
		return data
	}

	if tsyms == "" {
		return data
	}

	fsymsArr := parseFsyms(fsyms)
	tsymsArr := parseTsyms(tsyms)

	for _, v := range fsymsArr{
		switch v {
			case "BTC":
				for _, y := range tsymsArr{
					fmt.Println(y)
					switch y {
						case "USD":
							currenciesRawBtc.USD = getPriceRawFromDB("USD", v)
							currenciesDisplayBtc.USD =  getPriceDisplayFromDB("USD", v)
						case "EUR":
							currenciesRawBtc.EUR = getPriceRawFromDB("EUR", "BTC")
							currenciesDisplayBtc.EUR = getPriceDisplayFromDB("EUR", v)
									
					}
				}
		
				fsysmsRaw.BTC = currenciesRawBtc
				fsysmsDisplay.BTC = currenciesDisplayBtc
     		case "LINK":
				for _, y := range tsymsArr{
					fmt.Println(y)
					switch y {
						case "USD":
							currenciesRawLink.USD = getPriceRawFromDB("USD", v)
							currenciesDisplayLink.USD =  getPriceDisplayFromDB("USD", v)
						case "EUR":
							currenciesRawLink.EUR = getPriceRawFromDB("EUR", v)
							currenciesDisplayLink.EUR = getPriceDisplayFromDB("EUR", v)
									
					}
				}
		
				fsysmsRaw.LINK = currenciesRawLink
				fsysmsDisplay.LINK = currenciesDisplayLink			
			
		}
	}

	data.Raw = fsysmsRaw
	data.Display = fsysmsDisplay

	return data

}

func ClearRaw() {
	sqlStatement := `delete from raw`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec()

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}
}


func ClearDisplay() {
	sqlStatement := `delete from display`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec()

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}
}	

func AddRawBtcUsdPrice(data model.Data) {
	sqlStatement := `insert into raw(
		change24hour,
		changepct24hour,
		open24hour,
		volume24hour,
		volume24hourto,
		low24hour,
		high24hour,
		price,
		lastupdate,
		supply,
		mktcap,
		currency,
		ptype
	)values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

	`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec(
		data.Raw.BTC.USD.CHANGE24HOUR,
		data.Raw.BTC.USD.CHANGEPCT24HOUR,
		data.Raw.BTC.USD.OPEN24HOUR,
		data.Raw.BTC.USD.VOLUME24HOUR,
		data.Raw.BTC.USD.VOLUME24HOURTO,
		data.Raw.BTC.USD.LOW24HOUR,
		data.Raw.BTC.USD.HIGH24HOUR,
		data.Raw.BTC.USD.PRICE,
		data.Raw.BTC.USD.LASTUPDATE,
		data.Raw.BTC.USD.SUPPLY,
		data.Raw.BTC.USD.MKTCAP,
		"USD",
		"BTC",
	)

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func AddRawBtcEurPrice(data model.Data) {
	sqlStatement := `insert into raw(
		change24hour,
		changepct24hour,
		open24hour,
		volume24hour,
		volume24hourto,
		low24hour,
		high24hour,
		price,
		lastupdate,
		supply,
		mktcap,
		currency,
		ptype
	)values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

	`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec(
		data.Raw.BTC.EUR.CHANGE24HOUR,
		data.Raw.BTC.EUR.CHANGEPCT24HOUR,
		data.Raw.BTC.EUR.OPEN24HOUR,
		data.Raw.BTC.EUR.VOLUME24HOUR,
		data.Raw.BTC.EUR.VOLUME24HOURTO,
		data.Raw.BTC.EUR.LOW24HOUR,
		data.Raw.BTC.EUR.HIGH24HOUR,
		data.Raw.BTC.EUR.PRICE,
		data.Raw.BTC.EUR.LASTUPDATE,
		data.Raw.BTC.EUR.SUPPLY,
		data.Raw.BTC.EUR.MKTCAP,
		"EUR",
		"BTC",
	)

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func AddDisplayBtcUsdPrice(data model.Data) {
	sqlStatement := `insert into display(
		change24hour,
		changepct24hour,
		open24hour,
		volume24hour,
		volume24hourto,
		low24hour,
		high24hour,
		price,
		lastupdate,
		supply,
		mktcap,
		currency,
		ptype
	)values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

	`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec(
		data.Display.BTC.USD.CHANGE24HOUR,
		data.Display.BTC.USD.CHANGEPCT24HOUR,
		data.Display.BTC.USD.OPEN24HOUR,
		data.Display.BTC.USD.VOLUME24HOUR,
		data.Display.BTC.USD.VOLUME24HOURTO,
		data.Display.BTC.USD.LOW24HOUR,
		data.Display.BTC.USD.HIGH24HOUR,
		data.Display.BTC.USD.PRICE,
		data.Display.BTC.USD.LASTUPDATE,
		data.Display.BTC.USD.SUPPLY,
		data.Display.BTC.USD.MKTCAP,
		"USD",
		"BTC",
	)

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func AddDisplayBtcEurPrice(data model.Data) {
	sqlStatement := `insert into display(
		change24hour,
		changepct24hour,
		open24hour,
		volume24hour,
		volume24hourto,
		low24hour,
		high24hour,
		price,
		lastupdate,
		supply,
		mktcap,
		currency,
		ptype
	)values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

	`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec(
		data.Display.BTC.EUR.CHANGE24HOUR,
		data.Display.BTC.EUR.CHANGEPCT24HOUR,
		data.Display.BTC.EUR.OPEN24HOUR,
		data.Display.BTC.EUR.VOLUME24HOUR,
		data.Display.BTC.EUR.VOLUME24HOURTO,
		data.Display.BTC.EUR.LOW24HOUR,
		data.Display.BTC.EUR.HIGH24HOUR,
		data.Display.BTC.EUR.PRICE,
		data.Display.BTC.EUR.LASTUPDATE,
		data.Display.BTC.EUR.SUPPLY,
		data.Display.BTC.EUR.MKTCAP,
		"EUR",
		"BTC",
	)

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func AddRawLinkUsdPrice(data model.Data) {
	sqlStatement := `insert into raw(
		change24hour,
		changepct24hour,
		open24hour,
		volume24hour,
		volume24hourto,
		low24hour,
		high24hour,
		price,
		lastupdate,
		supply,
		mktcap,
		currency,
		ptype
	)values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

	`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec(
		data.Raw.LINK.USD.CHANGE24HOUR,
		data.Raw.LINK.USD.CHANGEPCT24HOUR,
		data.Raw.LINK.USD.OPEN24HOUR,
		data.Raw.LINK.USD.VOLUME24HOUR,
		data.Raw.LINK.USD.VOLUME24HOURTO,
		data.Raw.LINK.USD.LOW24HOUR,
		data.Raw.LINK.USD.HIGH24HOUR,
		data.Raw.LINK.USD.PRICE,
		data.Raw.LINK.USD.LASTUPDATE,
		data.Raw.LINK.USD.SUPPLY,
		data.Raw.LINK.USD.MKTCAP,
		"USD",
		"LINK",
	)

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func AddRawLinkEurPrice(data model.Data) {
	sqlStatement := `insert into raw(
		change24hour,
		changepct24hour,
		open24hour,
		volume24hour,
		volume24hourto,
		low24hour,
		high24hour,
		price,
		lastupdate,
		supply,
		mktcap,
		currency,
		ptype
	)values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

	`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec(
		data.Raw.LINK.EUR.CHANGE24HOUR,
		data.Raw.LINK.EUR.CHANGEPCT24HOUR,
		data.Raw.LINK.EUR.OPEN24HOUR,
		data.Raw.LINK.EUR.VOLUME24HOUR,
		data.Raw.LINK.EUR.VOLUME24HOURTO,
		data.Raw.LINK.EUR.LOW24HOUR,
		data.Raw.LINK.EUR.HIGH24HOUR,
		data.Raw.LINK.EUR.PRICE,
		data.Raw.LINK.EUR.LASTUPDATE,
		data.Raw.LINK.EUR.SUPPLY,
		data.Raw.LINK.EUR.MKTCAP,
		"EUR",
		"LINK",
	)

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func AddDisplayLinkUsdPrice(data model.Data) {
	sqlStatement := `insert into display(
		change24hour,
		changepct24hour,
		open24hour,
		volume24hour,
		volume24hourto,
		low24hour,
		high24hour,
		price,
		lastupdate,
		supply,
		mktcap,
		currency,
		ptype
	)values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

	`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec(
		data.Display.LINK.USD.CHANGE24HOUR,
		data.Display.LINK.USD.CHANGEPCT24HOUR,
		data.Display.LINK.USD.OPEN24HOUR,
		data.Display.LINK.USD.VOLUME24HOUR,
		data.Display.LINK.USD.VOLUME24HOURTO,
		data.Display.LINK.USD.LOW24HOUR,
		data.Display.LINK.USD.HIGH24HOUR,
		data.Display.LINK.USD.PRICE,
		data.Display.LINK.USD.LASTUPDATE,
		data.Display.LINK.USD.SUPPLY,
		data.Display.LINK.USD.MKTCAP,
		"USD",
		"LINK",
	)

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func AddDisplayLinkEurPrice(data model.Data) {
	sqlStatement := `insert into display(
		change24hour,
		changepct24hour,
		open24hour,
		volume24hour,
		volume24hourto,
		low24hour,
		high24hour,
		price,
		lastupdate,
		supply,
		mktcap,
		currency,
		ptype
	)values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)

	`

	sqlStr, err := db.Dbs.Prepare(sqlStatement)
	_, err = sqlStr.Exec(
		data.Display.LINK.EUR.CHANGE24HOUR,
		data.Display.LINK.EUR.CHANGEPCT24HOUR,
		data.Display.LINK.EUR.OPEN24HOUR,
		data.Display.LINK.EUR.VOLUME24HOUR,
		data.Display.LINK.EUR.VOLUME24HOURTO,
		data.Display.LINK.EUR.LOW24HOUR,
		data.Display.LINK.EUR.HIGH24HOUR,
		data.Display.LINK.EUR.PRICE,
		data.Display.LINK.EUR.LASTUPDATE,
		data.Display.LINK.EUR.SUPPLY,
		data.Display.LINK.EUR.MKTCAP,
		"EUR",
		"LINK",
	)

	defer sqlStr.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func parseFsyms(fsyms string) []string{
	return  strings.Split(fsyms, ",")
}

func parseTsyms(tsyms string) []string{
	return  strings.Split(tsyms, ",")
}

func getPriceRawFromDB(currency string, ptype string) *model.Raw {
	
	sqlStr, err :=  db.Dbs.Prepare(`
		SELECT 
			change24hour,
			changepct24hour,
			open24hour,
			volume24hour,
			volume24hourto,
			low24hour,
			high24hour,
			price,
			lastupdate,
			supply,
			mktcap
		FROM
			raw	
		WHERE 
			ptype = $1 AND currency = $2
		ORDER BY id DESC LIMIT 1`)

	if err != nil {
		fmt.Println(err)
	}

	defer sqlStr.Close()

	var change24hour float64
	var changepct24hour float64
	var open24hour float64
	var volume24hour float64
	var volume24hourto float64 
	var low24hour float64 
	var high24hour float64 
	var price float64 
	var lastupdate int64 
	var supply int64 
	var mktcap float64

	err = sqlStr.QueryRow(ptype, currency).Scan(
		&change24hour,
		&changepct24hour,
		&open24hour,
		&volume24hour,
		&volume24hourto, 
		&low24hour, 
		&high24hour, 
		&price, 
		&lastupdate, 
		&supply, 
		&mktcap,
	)

	if err != nil {
		fmt.Println(err)
	}
	
	rd := &model.Raw {
		CHANGE24HOUR : change24hour,
		CHANGEPCT24HOUR : changepct24hour,
		OPEN24HOUR: open24hour,
		VOLUME24HOUR: volume24hour,
		VOLUME24HOURTO: volume24hourto,
		LOW24HOUR: low24hour,
		HIGH24HOUR: high24hour,
		PRICE: price,
		LASTUPDATE: lastupdate,
		SUPPLY: supply,
		MKTCAP: mktcap,
	}	

	return rd
	
}

func getPriceDisplayFromDB(currency string, ptype string) *model.Display{

	sqlStr, err :=  db.Dbs.Prepare(`
		SELECT 
			change24hour,
			changepct24hour,
			open24hour,
			volume24hour,
			volume24hourto,
			low24hour,
			high24hour,
			price,
			lastupdate,
			supply,
			mktcap
		FROM
			display	
		WHERE 
			ptype = $1 AND currency = $2
		ORDER BY id DESC LIMIT 1`)

	if err != nil {
		fmt.Println(err)
	}

	defer sqlStr.Close()

	var change24hour string
	var changepct24hour string
	var open24hour string
	var volume24hour string
	var volume24hourto string 
	var low24hour string
	var high24hour string 
	var price string
	var lastupdate string
	var supply string 
	var mktcap string

	err = sqlStr.QueryRow(ptype, currency).Scan(
		&change24hour,
		&changepct24hour,
		&open24hour,
		&volume24hour,
		&volume24hourto, 
		&low24hour, 
		&high24hour, 
		&price, 
		&lastupdate, 
		&supply, 
		&mktcap,
	)

	if err != nil {
		fmt.Println(err)
	}
	
	di := &model.Display {
		CHANGE24HOUR : change24hour,
		CHANGEPCT24HOUR : changepct24hour,
		OPEN24HOUR: open24hour,
		VOLUME24HOUR: volume24hour,
		VOLUME24HOURTO: volume24hourto,
		LOW24HOUR: low24hour,
		HIGH24HOUR: high24hour,
		PRICE: price,
		LASTUPDATE: lastupdate,
		SUPPLY: supply,
		MKTCAP: mktcap,
	}	
	
	return di
}
