package parser

import (
	"fmt"
	"strings"

	utils "github.com/26thavenue/FXQLParser/util"
)


type FXQLData struct {
	SourceCurrency string
    DestinationCurrency string
    Buy          int
    Sell         int
    Cap          int
}

func Parse(input string) ([]*FXQLData, error) {

	// Split input into FXQL blocks, ensuring separation by a single newline
	blocks := strings.Split(input, "\n\n")
	var results []*FXQLData

	for i, block := range blocks {
		block = strings.TrimSpace(block)

		// Ensure no empty block exists
		if block == "" {
			return nil, fmt.Errorf("invalid input: empty FXQL statement at block %d", i+1)
		}

		// Validate each block as an independent FXQL statement
		lines := strings.Split(block, "\n")
		if len(lines) < 4 {
			return nil, fmt.Errorf("invalid input: insufficient lines in block %d", i+1)
		}

		// Validate and extract the currency pair
		header := strings.TrimSpace(lines[0])
		if !strings.Contains(header, " ") {
			return nil, fmt.Errorf("invalid input: missing space after currency pair in block %d", i+1)
		}

		parts := strings.SplitN(header, " ", 2)
		currencyPair := parts[0]

		cP := strings.Split(currencyPair, "-")
		if len(cP) != 2 {
			return nil, fmt.Errorf("invalid input: currency pair format is incorrect in block %d", i+1)
		}

		before := cP[0]
		after := cP[1]

		// Validate currency pair parts
		err := utils.ValidateCurrencyPair(before)
		if err != nil {
			return nil, fmt.Errorf("%s before , %s", err, before)
		}

		err = utils.ValidateCurrencyPair(after)
		if err != nil {
			return nil,fmt.Errorf("%s after , %s", err, after)
		}

		// Initialize FXQLData structure
		data := &FXQLData{
			SourceCurrency: before,
			DestinationCurrency:after,
		}

		

		// Process remaining lines in the block
		for _, line := range lines[1:] {
			

			line = strings.Trim(line, "[]")
			nV := strings.TrimSpace(line)
			nV = strings.Trim(nV, "}")

			if nV == "" {
				break
			}

			if strings.HasPrefix(nV, "BUY") {
				value, err := utils.CheckIntValue(nV, "BUY")
				if err != nil {
					return nil, fmt.Errorf("error in BUY value in block %d: %s", i+1, err)
				}
				data.Buy = value
			} else if strings.HasPrefix(nV, "SELL") {
				value, err := utils.CheckIntValue(nV, "SELL")
				if err != nil {
					return nil, fmt.Errorf("error in SELL value in block %d: %s", i+1, err)
				}
				data.Sell = value
			} else if strings.HasPrefix(nV, "CAP") {
				value, err := utils.CheckIntValue(nV, "CAP")
				if err != nil {
					return nil, fmt.Errorf("error in CAP value in block %d: %s", i+1, err)
				}
				data.Cap = value
			} else  {
				fmt.Printf("Unexpected line format: '%s'\n", nV)
				return nil, fmt.Errorf("Empty FXQL statement, %v", data)
			}
		}

		results = append(results, data)
	}

	return results, nil
}
