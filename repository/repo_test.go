package repository

import (
	"testing"

	"github.com/26thavenue/FXQLParser/database"
	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	tests := []struct {
		input       string
		name        string
		expected    []Response
		expectError bool
	}{
		{
			name:  "Single Currency Pair",
			input: `USD-GBP { 
						BUY 85 
						SELL 94 
						CAP 12000
					}`,
			expected: []Response{
				{
					SourceCurrency:      "USD",
					DestinationCurrency: "GBP",
					BuyPrice:            85,
					SellPrice:           94,
					CapAmount:           12000,
				},
			},
			expectError: false,
		},
		{
			name: "Multiple Currency Pairs",
			input: `USD-GBP { 
					BUY 85 
					SELL 90 
					CAP 10000 
				}
				EUR-JPY { 
					BUY 130 
					SELL 135 
					CAP 5000 
				}`,
			expected: []Response{
				{
					SourceCurrency:      "USD",
					DestinationCurrency: "GBP",
					BuyPrice:            85,
					SellPrice:           90,
					CapAmount:           10000,
				},
				{
					SourceCurrency:      "EUR",
					DestinationCurrency: "JPY",
					BuyPrice:            130,
					SellPrice:           135,
					CapAmount:           5000,
				},
			},
			expectError: false,
		},
		{
			name: "Duplicate Currency Pair",
			input: `USD-GBP { 
					BUY 85 
					SELL 90 
					CAP 10000 
				}
				USD-GBP { 
					BUY 86 
					SELL 91 
					CAP 12000 
				}`,
			expected: []Response{
				{
					SourceCurrency:      "USD",
					DestinationCurrency: "GBP",
					BuyPrice:            86,
					SellPrice:           91,
					CapAmount:           12000,
				},
			},
			expectError: false,
		},
		{
			name:        "Invalid FXQL Statement",
			input:       `USD-GBP {
						}`,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := Transform(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				t.Logf("%v Logs",r )
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, *r)
			}
		})
	}
}

func init() {
	database.DBInstance = &database.DB{
		Instance: nil,
	}
}

// func TestCreate(t *testing.T) {
	
// 	input := `USD-GBP {
// 						BUY 100
// 						SELL 200
// 		 				CAP 93800
// 			}`
// 	err := Create(input)
	
// 	assert.Nil(t, err, "Expected no error from Create function")
	
// 	exists := CheckCurrencyPair("USD", "EUR")
// 	assert.True(t, exists, "Currency pair USD/EUR should exist in the database")
// }
