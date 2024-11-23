package main

import (
    "log"
    "github.com/gorilla/mux"
    "net/http"
)

func main() {

    //Routes
    r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handleReceipts).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlePoints).Methods("GET")
    
    http.Handle("/", r)
    log.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)

}