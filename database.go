package main

//Structure for Receipt and Items
type Receipt struct {
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items []Item `json:"items"`
	Total string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

//Map for temp storage
var receiptsDatabase = map[string]Receipt{}
var pointsDatabase = map[string]int{}