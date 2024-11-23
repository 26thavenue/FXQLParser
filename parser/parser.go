package parser

import (
	"fmt"
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

func (r *Response) ValidateCurrencyPair() error {

	err := ValidateCurrencyPair(r.CURR1)
	if err != nil{
		return fmt.Errorf( "%s",err)
	}

	err = ValidateCurrencyPair(r.CURR2)
	if err != nil{
		return fmt.Errorf( "%s",err)
	}

	return nil
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