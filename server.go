package main

import (
    "log"
    "net/http"
	"github.com/gorilla/mux"
	handler "allyapps.com/bequant/handlers"
    db "allyapps.com/bequant/db" 

)

func handleRequests() {
    db.Init()

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", handler.ReturnHome)
    router.HandleFunc("/price", handler.ReturnAllPrice)
    log.Fatal(http.ListenAndServe(":16000", router))
}

func main() {
	handleRequests()
}


