package parser

import (
	_ "encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Response struct {
	CURR1 string
	CURR2 string
	BUY   uint64
	SELL  uint64
	CAP   uint64
}

type FXQLData struct {
    CurrencyPair string
    Buy          int
    Sell         int
    Cap          int
}

func Parse(input string) (*FXQLData, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) < 4 {
		return nil, errors.New("invalid input: insufficient lines")
	}

	// Validate and extract the currency pair
	header := strings.TrimSpace(lines[0])
	if !strings.Contains(header, " ") {
		return nil, errors.New("invalid input: missing space after currency pair")
	}
	parts := strings.SplitN(header, " ", 2)

	currencyPair := parts[0]

    cP := strings.Split(currencyPair, "-")

	if len(cP) != 2 {
		return nil , fmt.Errorf("Invalid input format")
	}

	before := cP[0]
	after := cP[1]

	err := ValidateCurrencyPair(before) 

	if err !=nil{
		return nil , fmt.Errorf("Error %s", err)
	}

	err = ValidateCurrencyPair(after) 

	if err !=nil{
		return nil , fmt.Errorf("Error %s", err)
	}

	// Extract BUY, SELL, CAP
	data := FXQLData{
		CurrencyPair: currencyPair,
	}
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "BUY") {
			value, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "BUY")))
			if err != nil {
				return nil, fmt.Errorf("invalid BUY value: %w", err)
			}
			data.Buy = value
		} else if strings.HasPrefix(line, "SELL") {
			value, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "SELL")))
			if err != nil {
				return nil, fmt.Errorf("invalid SELL value: %w", err)
			}
			data.Sell = value
		} else if strings.HasPrefix(line, "CAP") {
			value, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "CAP")))
			if err != nil {
				return nil, fmt.Errorf("invalid CAP value: %w", err)
			}
			data.Cap = value
		} else {
			return nil, fmt.Errorf("unexpected line: %s", line)
		}
	}

	return &data, nil
}

func ValidateCurrencyPair(s string ) error {
	err := checkUpperCase(s)
	if err != nil{
		return fmt.Errorf( "%s",err)
	}

	err = hasThreeLetters(s)
	if err != nil{
		return fmt.Errorf( "%s",err)
	}

	return nil
}

func checkUpperCase(s string) error {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return fmt.Errorf("Invalid: %s should be %s", s, strings.ToUpper(s))
		}
	}
	return nil
}

func hasThreeLetters(s string) error {
	count := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			count++
		}
	}
	if count == 3 {
		return fmt.Errorf("Currency must be exactly 3 uppercase characters")
	}
	return nil
}