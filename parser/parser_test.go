package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T){
	tests := []struct{
		name string
		input string
		expectError bool
		errorMessage string
	}{
		{
			name:"Small Currency Pair Value",
			input : `usd-GBP {  
						BUY 100
						SELL 200
						CAP 93800
						}
						`,
			expectError: true,
			errorMessage: "Invalid: 'usd' should be 'USD'",
		},
		{
			name:"Space After Currency Pair",
			input : `USD-GBP{
						BUY 100
						SELL 200
						CAP 93800
						}`,
			expectError: true,
			errorMessage: "Missing single space after currency pair",
		},
		{
			name:"Invalid Integer",
			input : `USD-GBP {
						BUY abc
						SELL 200
						CAP 93800
						}`,
			expectError: true,
			errorMessage: "Invalid: 'abc' is not a valid numeric amount",
		},
		{
			name:"Negative Integer",
			input : `USD-GBP {
						BUY 100
						SELL 200
						CAP -50 
						}`,
			expectError: true,
			errorMessage: "Invalid: CAP cannot be a negative number",
		},
		{
			name:"Empty FXQL statement",
			input : `USD-GBP {
						}`,
			expectError: true,
			errorMessage: "Invalid: Empty FXQL statement",
		},
		{
			name:"Multiple FXQL Statements error",
			input : `USD-GBP {
						BUY 100
						SELL 200
						CAP 93800
						}EUR-JPY {
						BUY 80
						SELL 90
						CAP 50000
						}`,
			expectError: true,
			errorMessage:"Invalid: Multiple FXQL statements should be separated by a single newline character",
		},
		{
			name:"Multiple newlines ",
			input : `USD-GBP {
						BUY 100

						SELL 200

						CAP 93800
						}`,
			expectError: true,
			errorMessage: "Invalid: Multiple newlines within a single FXQL statement",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.errorMessage, result)
			}
		})
	}

	t.Run("Valid FXQL statement", func(t *testing.T){
		input := `USD-GBP {
						BUY 100
						SELL 200
		 				CAP 93800
			}`

		value, err := Parse(input)
		assert.NoError(t, err)
		assert.Equal(t, []*FXQLData{
			{
				SourceCurrency: "USD",
				DestinationCurrency: "GBP",
				Buy:          100,
				Sell:         200,
				Cap:          93800,
			},
		}, value)
	})
	
	t.Run("Valid Double FXQL statement", func(t *testing.T){
		input := `USD-GBP {
						BUY 100
						SELL 200
		 				CAP 93800
					}
					USD-GBP {
						BUY 100
						SELL 200
		 				CAP 93800
					}`

		value, err := Parse(input)
		assert.NoError(t, err)
		assert.Equal(t, []*FXQLData{
			{
				SourceCurrency: "USD",
				DestinationCurrency: "GBP",
				Buy:          100,
				Sell:         200,
				Cap:          93800,
			},
			{
				SourceCurrency: "USD",
				DestinationCurrency: "GBP",
				Buy:          100,
				Sell:         200,
				Cap:          93800,
			},
		}, value)
	})
}


