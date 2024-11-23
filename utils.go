package main

import (
	"fmt"
	"time"
	"regexp"
)

//To Check Aplha Numer Character in String
func isAlphaNumeric(x rune) bool {
    return (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') || (x >= '0' && x <= '9')
}

//Validating JSON
func validateJSON(receipt Receipt) error {
	if receipt.Retailer == "" {
		return fmt.Errorf("Missing required field: Retailer")
	}
	if receipt.PurchaseDate == "" {
		return fmt.Errorf("Missing required field: PurchaseDate")
	}

	_, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return fmt.Errorf("Invalid PurchaseDate format, expected YYYY-MM-DD: %v", err)
	}

	if receipt.PurchaseTime == "" {
		return fmt.Errorf("Missing required field: PurchaseTime")
	}

	match, _ := regexp.MatchString(`^([01]?[0-9]|2[0-3]):([0-5][0-9])$`, receipt.PurchaseTime)
	if !match {
		return fmt.Errorf("Invalid PurchaseTime format, expected HH:MM in 24-hour format")
	}

	if len(receipt.Items) == 0 {
		return fmt.Errorf("Atleast one item is required in items array")
	}
	for i, item := range receipt.Items {
		if item.ShortDescription == "" {
			return fmt.Errorf("Item %d is missing ShortDescription", i+1)
		}
		if item.Price == "" {
			return fmt.Errorf("Item %d is missing Price", i+1)
		}
	}
	if receipt.Total == "" {
		return fmt.Errorf("Missing required field: Total")
	}
	return nil
}