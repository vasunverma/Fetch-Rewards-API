package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func handleReceipts(w http.ResponseWriter, r *http.Request){
	
	//Allowing Only POST Req
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	//Decoding JSON
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if(err != nil){
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//Validating JSON
	err = validateJSON(receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Creating new UUID
	id:= uuid.New()
	//Adding to receiptsDatabase Map
	receiptsDatabase[id.String()] = receipt

	//Sending ID Back
	w.Header().Set("Content-Type", "application/json")
	var out = map[string]string{
		"id" : id.String(),
	}
	json.NewEncoder(w).Encode(out)

}

func handlePoints(w http.ResponseWriter, r *http.Request){
	
	//Allowing Only GET Req
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extracting ID
	id := mux.Vars(r)["id"]

	//Returning Points if ID exists
	if points, ok := pointsDatabase[id]; ok{
		w.Header().Set("Content-Type", "application/json")
		var out = map[string]int{
			"points" : points,
		}

		json.NewEncoder(w).Encode(out)
		return 
	}

	//Calculating Total Points for this Receipt
	totalPoints, err := CalculatePoints(id)

	//Handling Error for Calculating Points
	if err != nil {
		http.Error(w, fmt.Sprintf("Error calculating points: %v", err), http.StatusBadRequest)
		return
	}

	//Returning Points
	w.Header().Set("Content-Type", "application/json")
	var out = map[string]int{
		"points" : totalPoints,
	}
	json.NewEncoder(w).Encode(out)
}
