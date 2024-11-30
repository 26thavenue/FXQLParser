package parser

import (
	"fmt"
	"log"
	_ "log"
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

func Parse(input string) ([]FXQLData, error) {
	// Split the input by single newline and process each block
	blocks := strings.Split(input, "\n")

	var results []FXQLData
	var currentBlock []string

	for _, line := range blocks {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if len(currentBlock) > 0 && strings.HasPrefix(line, "{") {
			return nil, fmt.Errorf("Invalid: Multiple FXQL statements should be separated by a single newline character")
		}
		
		currentBlock = append(currentBlock, line)
		if strings.HasSuffix(line, "}") {
			data, err := processBlock(currentBlock)
			if err != nil {
				return nil, err
			}
			results = append(results, data)
			currentBlock = []string{} 
		}
	}

	return results, nil
}

func processBlock(lines []string) (FXQLData, error) {
	if len(lines) < 2 {
		return FXQLData{}, fmt.Errorf("insufficient lines in block")
	}

	header := strings.TrimSpace(lines[0])
	if !strings.Contains(header, " ") {
		return FXQLData{}, fmt.Errorf("missing space after currency pair")
	}

	parts := strings.SplitN(header, " ", 2)
	currencyPair := parts[0]
	cP := strings.Split(currencyPair, "-")
	if len(cP) != 2 {
		return FXQLData{}, fmt.Errorf("currency pair format incorrect")
	}

	before, after := cP[0], cP[1]

	if err := utils.ValidateCurrencyPair(before); err != nil {
		return FXQLData{}, fmt.Errorf("%s before, %s", err, before)
	}
	if err := utils.ValidateCurrencyPair(after); err != nil {
		return FXQLData{}, fmt.Errorf("%s after, %s", err, after)
	}

	data := FXQLData{
		SourceCurrency:      before,
		DestinationCurrency: after,
	}

	
	for i, line := range lines[1:] {

		if strings.Contains(line, "\n\n") {
			return FXQLData{}, fmt.Errorf("Invalid: Multiple newlines within a single FXQL statement")
		}

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "BUY") {
				value, err := utils.CheckIntValue(line, "BUY")
				if err != nil {
					return FXQLData{}, fmt.Errorf("error in BUY value in block %d: %s", i+1, err)
				}
				data.Buy = value
		} else if strings.HasPrefix(line, "SELL") {
				value, err := utils.CheckIntValue(line, "SELL")
				if err != nil {
					return FXQLData{}, fmt.Errorf("error in SELL value in block %d: %s", i+1, err)
				}
				data.Sell = value
		} else if strings.HasPrefix(line, "CAP") {
				value, err := utils.CheckIntValue(line, "CAP")
				if err != nil {
					return FXQLData{}, fmt.Errorf("error in CAP value in block %d: %s", i+1, err)
				}
				data.Cap = value
		}else{
			    log.Printf("%v",line)
				return FXQLData{}, fmt.Errorf("Invalid: Empty FXQL statement")
		}
		
	}

	return data, nil
}



// Validate remaining lines
		// hasValidContent := false
		// for _, line := range lines[1:] {
		// 	nV := strings.TrimSpace(strings.Trim(line, "{}"))
		// 	if nV == "" {
		// 		continue
		// 	}

		// 	hasValidContent = true

		// 	if strings.HasPrefix(nV, "BUY") {
		// 		value, err := utils.CheckIntValue(nV, "BUY")
		// 		if err != nil {
		// 			return nil, fmt.Errorf("error in BUY value in block %d: %s", i+1, err)
		// 		}
		// 		data.Buy = value
		// 	} else if strings.HasPrefix(nV, "SELL") {
		// 		value, err := utils.CheckIntValue(nV, "SELL")
		// 		if err != nil {
		// 			return nil, fmt.Errorf("error in SELL value in block %d: %s", i+1, err)
		// 		}
		// 		data.Sell = value
		// 	} else if strings.HasPrefix(nV, "CAP") {
		// 		value, err := utils.CheckIntValue(nV, "CAP")
		// 		if err != nil {
		// 			return nil, fmt.Errorf("error in CAP value in block %d: %s", i+1, err)
		// 		}
		// 		data.Cap = value
		// 	}
		// }

		// if !hasValidContent {
		// 	return nil, fmt.Errorf("Invalid: Empty FXQL statement in block %d", i+1)
		// }
