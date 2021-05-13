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

    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", handler.ReturnHome)
    myRouter.HandleFunc("/price", handler.ReturnAllPrice)
    log.Fatal(http.ListenAndServe(":16000", myRouter))
}

func main() {
	handleRequests()
}


