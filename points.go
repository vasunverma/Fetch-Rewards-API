package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func CalculatePoints(id string) (int, error){

	//Checking if receipt for this ID exists
	receipt, ok := receiptsDatabase[id]
	if !ok {
		return 0, fmt.Errorf("receipt with ID %s not found", id)
	}

	totalPoints := 0
	totalPoints += CalculateAplhaNumPoints(receipt.Retailer)

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil{
		return 0, fmt.Errorf("Error while parsing total amount: %v", err)
	}

	totalPoints += CalculateTotalCentAndMultiplePoints(total)
	
	itemPoints, err := CalculatePointsPerItem(receipt.Items)
	if err != nil {
		return 0, fmt.Errorf("Error calculating item points: %v", err)
	}
	totalPoints += itemPoints


	purchaseDatePoints, err := CalculatePurchaseDatePoints(receipt.PurchaseDate)
	if err != nil {
		return 0, fmt.Errorf("Error calculating purchase date points: %v", err)
	}
	totalPoints += purchaseDatePoints

	purchaseTimePoints, err := CalculatePurchaseTimePoints(receipt.PurchaseTime)
	if err != nil {
		return 0, fmt.Errorf("Error calculating purchase time points: %v", err)
	}
	totalPoints += purchaseTimePoints

	//Storing in map
	pointsDatabase[id] = totalPoints

	return totalPoints, nil
}

//Counting AlphaNumeric Characters and adding Points
func CalculateAplhaNumPoints(retailer string) int{
	var nAlphaNuChar int
	for _, char := range retailer{
		if(isAlphaNumeric(char)){
			nAlphaNuChar++
		}
	}
	return nAlphaNuChar
}

//Adding Points if no cents or if multiple of 0.25
func CalculateTotalCentAndMultiplePoints(total float64) int{
	var points int = 0
	if total == math.Floor(total){
		points += 50
	}
	if math.Mod(total*4, 1) == 0{
		points += 25
	}
	return points
}

func CalculatePointsPerItem(items []Item) (int, error){
	var points int = 0
	nItems := len(items)
	points += 5*(nItems/2)

	for _, item := range items {
		//Trimming String
		itemLen := len(strings.TrimSpace(item.ShortDescription))
		if itemLen % 3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, fmt.Errorf("Error while parsing item price: %v", err)
			}
			//Rounding
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points, nil
}

//Adding points if date is odd
func CalculatePurchaseDatePoints(purchaseDate string) (int, error){
	var points int = 0
	purchaseDateArr := strings.Split(purchaseDate, "-")
	purchaseDay := purchaseDateArr[2]

	day, err := strconv.Atoi(purchaseDay)
	if err != nil {
		return 0, fmt.Errorf("Error while parsing purchase day to int: %v", err)
	}

	if day%2 != 0 {
		points += 6
	}
	
	return points, nil
}

//Adding points if purchase after 2:00pm and before 4:00pm
func CalculatePurchaseTimePoints(purchaseTime string) (int, error){
	var points int = 0
	purchaseTimeArr := strings.Split(purchaseTime, ":")
	purchaseHour, err := strconv.Atoi(purchaseTimeArr[0])
	if err != nil {
		return 0, fmt.Errorf("Error while parsing purchase hour to int: %v", err)
	}
	purchaseMin, err := strconv.Atoi(purchaseTimeArr[1])
	if err != nil {
		return 0, fmt.Errorf("Error while parsing purchase minutes to int: %v", err)
	}

	if (purchaseHour == 14 && purchaseMin >= 1) || purchaseHour == 15{
		points += 10
	}
	return points, nil
}