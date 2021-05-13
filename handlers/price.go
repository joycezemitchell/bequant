package handler

import (
    "net/http"
    "encoding/json"
	pg "allyapps.com/bequant/data/postgres"
)

func ReturnAllPrice(w http.ResponseWriter, r *http.Request){
	f := r.URL.Query().Get("fsyms")
	t := r.URL.Query().Get("tsyms")
	data := pg.GetPrice(f,t)
   	json.NewEncoder(w).Encode(data)
}
