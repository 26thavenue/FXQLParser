package repository

import (
	"fmt"
	"log"

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

func Create (input string) (*[]Response, error){

	if input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	vr,err := parser.Parse(input)

	if err != nil{
		return nil, fmt.Errorf("%s", err)
	}

	currencyMap := make(map[string]Response)

	for _, item:= range vr {
		log.Printf(" %v", item)
	}

	for _, fx := range vr {
		key := fx.SourceCurrency + "-" + fx.DestinationCurrency

		currencyMap[key] = Response{
			SourceCurrency:      fx.SourceCurrency,
			DestinationCurrency: fx.DestinationCurrency,
			BuyPrice:            fx.Buy,
			SellPrice:           fx.Sell,
			CapAmount:           fx.Cap,
		}
		// log.Printf("%v - Inside",currencyMap)
	}

	var responses []Response
	for _, value := range currencyMap {
		responses = append(responses, value)
	}

	return &responses, nil

}

func CheckCurrencyPair(des, source string) bool{
	return false
}

func Update(input string) (*Response, error){

	return nil, nil
}