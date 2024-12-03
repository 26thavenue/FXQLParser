package parser

import (
	"fmt"
	"log"
	_ "log"
	"strings"

	utils "github.com/26thavenue/FXQLParser/util"
)

const (
	ErrInsufficientLines      = "Insufficient amount of lines"
	ErrTooManyLines           = "Too many lines"
	ErrMultipleStatements     = "Invalid: Multiple FXQL statements should be separated by a single newline character"
	ErrMultipleNewlinesWithin = "Invalid: Multiple newlines within a single FXQL statement"
)


type FXQLData struct {
	SourceCurrency string
    DestinationCurrency string
    Buy          int
    Sell         int
    Cap          int
}

func ProcessStrings(input string) error{
	blocks := strings.Split(input, "}")
	lines := strings.Count(input, "\n")
	if lines < 4 || len(blocks) < 2 {
		return fmt.Errorf(ErrInsufficientLines)
	}
	if len(blocks) > 2 {
		if len(blocks) > 3 {
			return fmt.Errorf(ErrTooManyLines)
		} else if len(blocks) == 3{
			if lines < 9 {
				return fmt.Errorf(ErrMultipleStatements)
			} else if lines > 9 {
				return fmt.Errorf(ErrTooManyLines)
			}
		}
	}
	if len(blocks) == 2 && lines != 4 {
		return fmt.Errorf(ErrMultipleNewlinesWithin)
	}
	return nil
}

func Parse(input string) ([]FXQLData, error) {
	// Split the input by single newline and process each block

    err := ProcessStrings(input)
	if err != nil {
		return nil, err
	}
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
			data, err := ProcessBlock(currentBlock)
			if err != nil {
				return nil, err
			}
			results = append(results, data)
			currentBlock = []string{} 
		}
	}

	return results, nil
}

func ProcessBlock(lines []string) (FXQLData, error) {

	fmt.Println(len(lines))
	if len(lines) < 2 {
		return FXQLData{}, fmt.Errorf("Invalid: Empty FXQL statement, %d", len(lines))
	}

	fmt.Printf("%v",lines)

	header := strings.TrimSpace(lines[0])
	if !strings.Contains(header, " ") {
		return FXQLData{}, fmt.Errorf("missing space after currency pair")
	}

	fmt.Printf("%v",header)

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

		fmt.Printf("%v",line)

		line = strings.TrimSpace(line)

		fmt.Printf("%v",line)

		if line == "" {
			log.Printf("%v",line)
			return FXQLData{}, fmt.Errorf("Invalid: Empty FXQL statement")
				
		} 
		if strings.HasPrefix(line, "SELL") {
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
		}else if strings.HasPrefix(line, "BUY"){
			    value, err := utils.CheckIntValue(line, "BUY")
				if err != nil {
					return FXQLData{}, fmt.Errorf("error in BUY value in block %d: %s", i+1, err)
				}
				data.Buy = value
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
