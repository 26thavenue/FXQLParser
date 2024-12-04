package repository

import (
	"fmt"
	"log"
	_ "log"

	"github.com/26thavenue/FXQLParser/database"
	"github.com/26thavenue/FXQLParser/parser"
)

type Response struct {
	EntryId string
	SourceCurrency string
	DestinationCurrency string
	SellPrice int 
	BuyPrice int 
	CapAmount int
}

func Transform (input string) (*[]Response, error){
	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}
	vr,err := parser.Parse(input)

	if err != nil{
		return nil, fmt.Errorf("%s", err)
	}

	if vr == nil {
		return nil, fmt.Errorf("parser returned nil data")
	}

	// MAP OVER THE VALUES OF VR ,APPEND IT TO A MAP AND PASS IT TO THE RESPONSE 
	pairs := make(map[string]parser.FXQLData)

	for _, data := range vr {
		key := fmt.Sprintf("%s-%s", data.SourceCurrency, data.DestinationCurrency)
		pairs[key] = data 
	}

	var responses []Response
	for _, data := range pairs {
		responses = append(responses, Response{
			SourceCurrency:    data.SourceCurrency,
			DestinationCurrency: data.DestinationCurrency,
			SellPrice:         data.Sell,
			BuyPrice:          data.Buy,
			CapAmount:         data.Cap,
		})
	}

	return &responses, nil

}

func Create(input string) error {
	responses, err := Transform(input)
	if err != nil {
		return fmt.Errorf("failed to transform input: %v", err)
	}

	for _, response := range *responses {
		transaction := &database.Transaction{
			SourceCurrency:     response.SourceCurrency,
			DestinationCurrency: response.DestinationCurrency,
			SellPrice:          response.SellPrice,
			BuyPrice:           response.BuyPrice,
			CapAmount:          response.CapAmount,
		}

		err := database.DBInstance.Instance.Create(transaction).Error
		if err != nil {
			log.Printf("Error inserting transaction: %v", err)
			return fmt.Errorf("failed to insert data: %v", err)
		}
	}

	return nil
}

func CheckCurrencyPair(source, destination string) bool {
	var count int64
	err := database.DBInstance.Instance.Model(&database.Transaction{}).
		Where("source_currency = ? AND destination_currency = ?", source, destination).
		Count(&count).Error

	if err != nil {
		log.Printf("Error checking currency pair: %v", err)
		return false
	}

	return count > 0
}

// func Update(input string) (*Response, error){

// 	return nil, nil
// }