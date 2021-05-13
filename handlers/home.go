package handler

import (
    "fmt"
    "net/http"
)

func ReturnHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome Bequant REST API")
    fmt.Println("Endpoint Hit: homePage")
}