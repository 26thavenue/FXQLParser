package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	r := Response{
		CURR1: "usd",
		CURR2: "GER",
		BUY:340,
		SELL:30,
		CAP:30,
	}

	t.Run("Check for Valid Currency Pair", func (t *testing.T){
		r := r
		assert.EqualError(t, r.ValidateCurrencyPair(), "Invalid: usd should be USD" )
	})

	t.Run("Check for Valid Currency Pair", func (t *testing.T){
		r := r
		r.CURR1 = "USDT"
		assert.EqualError(t, r.ValidateCurrencyPair(), "Currency must be exactly 3 uppercase characters" )
	})
}