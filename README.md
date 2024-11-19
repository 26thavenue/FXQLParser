# Foreign Exchange Query Language (FXQL) Statement Parser Implementation

## Overview
Make a parser that will serve as part of a central federation system for Bureau De Change (BDC) operations, allowing them to submit and standardize their exchange rate information.

## Core Functionlities
The system should:
1. Accept FXQL statements via API calls
2. Validate FXQL syntax
3. Store valid entries in the database
4. Return appropriate responses for both successful and failed operations

## FXQL Statement Specification

### Basic Structure
```
CURR1-CURR2 {
 BUY {AMOUNT}
 SELL {AMOUNT}
 CAP {AMOUNT}
}
```

### Rules and Constraints

- **CURR1** (Source Currency):
    - Must be exactly 3 uppercase characters
    - Valid: USD, GBP, EUR
    - Invalid: usd, USDT, US
- **CURR2** (Destination Currency):
    - Must be exactly 3 uppercase characters
    - Same rules as CURR1
- **BUY**:
    - Numeric amount in CURR2 per unit of CURR1
    - Valid: 300, 450.43, 0.04590
    - Invalid: abda, -138, 0..12039
- **SELL**:
    - Numeric amount in CURR2 per unit of CURR1
    - Same validation rules as BUY
- **CAP**:
    - Maximum transaction amount in CURR1
    - Must be a whole number
    - Can be 0 (indicating no cap)
    - Invalid: negative numbers, decimals

## Example Valid FXQL Statements
```
USD-GBP {
 BUY 100
 SELL 200
 CAP 93800
}
KES-NGN {
 BUY 150
 SELL 210
 CAP 9500
}
```

## Example Invalid FXQL Statements

```
usd-GBP {  # Invalid: 'usd' should be 'USD'
 BUY 100
 SELL 200
 CAP 93800
}
USD-GBP{ # Missing single space after currency pair
 BUY 100
 SELL 200
 CAP 93800
}
USD-GBP {
 BUY abc # Invalid: 'abc' is not a valid numeric amount
 SELL 200
 CAP 93800
}
USD-GBP {
 BUY 100
 SELL 200
 CAP -50 # Invalid: CAP cannot be a negative number
}
USD-GBP {
} # Invalid: Empty FXQL statement
USD-GBP {
 BUY 100
 SELL 200
 CAP 93800
}
EUR-JPY {
 BUY 80
 SELL 90
 CAP 50000
}
# Invalid: Multiple FXQL statements should be separated by a single newline character
USD-GBP {
 BUY 100
 SELL 200
 CAP 93800
} # Invalid: Multiple newlines within a single FXQL statement
```

## API Endpoint Specification

### POST /fxql-statements

Request Body:
```json
{
  "FXQL": "USD-GBP {\\n BUY 100\\n SELL 200\\n CAP 93800\\n}"
}
```

Successful Response -OK 
```
{
  "message": "Rates Parsed Successfully.",
  "code": "FXQL-200",
  "data": [
    {
      "EntryId": 192,
      "SourceCurrency": "USD",
      "DestinationCurrency": "GBP",
      "SellPrice": 200,
      "BuyPrice": 100,
      "CapAmount": 93800
    }
  ]
}
```
Note: The EntryId values in the response are arbitrary and will be determined by your database implementation.


### Bonus Points

- Comprehensive error messages with line numbers and character positions
- Unit and/or integration tests
- API documentation (Swagger/OpenAPI)
- Rate limiting
- Input sanitization
- Logging system
- Docker configuration
- Database design
- API design and implementation


Link to Assesment -[https://miraapp.notion.site/Backend-Developer-Technical-Assessment-a954df277ad34772a261ddfe2dd7210c]

