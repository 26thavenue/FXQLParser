package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func ValidateCurrencyPair(s string) error {
	err := checkUpperCase(s)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	err = hasThreeLetters(s)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func checkUpperCase(s string) error {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return fmt.Errorf("Invalid: '%s' should be '%s' ", s, strings.ToUpper(s))
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
	if count != 3 {
		return fmt.Errorf("Currency must be exactly 3 uppercase characters")
	}
	return nil
}

func CheckIntValue(s ,Prefix string) (int,error ){
	trim := strings.TrimSpace(s)
	tr := strings.TrimPrefix(trim, Prefix)
	value, err := strconv.Atoi(strings.TrimSpace(tr))
	if err != nil {
		return  0, fmt.Errorf("Invalid: '%s' is not a valid numeric amount , %v , %v ", Prefix, err, tr)
	}
	
	if value < 0 {
		return 0,fmt.Errorf("Invalid: CAP cannot be a negative number")
	}

	
	return value, nil
}