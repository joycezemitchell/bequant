package model 

type Data struct {
	Raw *FsysmsRaw `json:"RAW,omitempty"`
	Display *FsysmsDisplay `json:"DISPLAY,omitempty"`
}

type FsysmsRaw struct {
	BTC *CurrenciesRaw `json:"BTC,omitempty"`
	LINK  *CurrenciesRaw `json:"LINK,omitempty"`
}

type FsysmsDisplay struct {
	BTC *CurrenciesDisplay `json:"BTC,omitempty"`
	LINK  *CurrenciesDisplay `json:"LINK,omitempty"`
}

type CurrenciesRaw struct {
	USD *Raw `json:"USD,omitempty"`
	EUR *Raw `json:"EUR,omitempty"`
}

type CurrenciesDisplay struct {
	USD *Display `json:"USD,omitempty"`
	EUR *Display `json:"EUR,omitempty"`
}

type Raw struct {
	CHANGE24HOUR float64 `json:"CHANGE24HOUR,omitempty"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR,omitempty"`
	OPEN24HOUR float64 `json:"OPEN24HOUR,omitempty"`
	VOLUME24HOUR float64 `json:"VOLUME24HOUR,omitempty"`
	VOLUME24HOURTO float64 `json:"VOLUME24HOURTO,omitempty"`
	LOW24HOUR float64 `json:"LOW24HOUR,omitempty"`
	HIGH24HOUR float64 `json:"HIGH24HOUR,omitempty"`
	PRICE float64 `json:"PRICE,omitempty"`
	LASTUPDATE int64 `json:"LASTUPDATE,omitempty"`
	SUPPLY int64 `json:"SUPPLY,omitempty"`
	MKTCAP float64 `json:"MKTCAP,omitempty"`
}

type Display struct {
	CHANGE24HOUR string `json:"CHANGE24HOUR,omitempty"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR,omitempty"`
	OPEN24HOUR string `json:"OPEN24HOUR,omitempty"`
	VOLUME24HOUR string `json:"VOLUME24HOUR,omitempty"`
	VOLUME24HOURTO string `json:"VOLUME24HOURTO,omitempty"`
	LOW24HOUR string `json:"LOW24HOUR,omitempty"`
	HIGH24HOUR string `json:"HIGH24HOUR,omitempty"`
	PRICE string `json:"PRICE,omitempty"`
	LASTUPDATE string `json:"LASTUPDATE,omitempty"`
	SUPPLY string `json:"SUPPLY,omitempty"`
	MKTCAP string `json:"MKTCAP,omitempty"`
}


