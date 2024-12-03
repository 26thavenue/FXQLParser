package repository

import (
	"fmt"
	_ "log"

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

func CheckCurrencyPair(des, source string) bool{
	return false
}

func Update(input string) (*Response, error){

	return nil, nil
}