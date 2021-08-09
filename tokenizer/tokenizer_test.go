/*
Tokenizer
- input:
    - string
    - cleaned
- output:
    - list of tokens
        - what type is it? (e.g. slice, linked list, interface, struct)
        - brackets, identifiers (func names, variables), values (number, "string", boolean #t#f, other literals)
        - ignore symbols
        - (extend later if necessary)
    - token (struct, interface)
        - value (e.g. identifier -> string, number -> number, bracket -> nil)
        - type
*/
package tokenizer_test

import (
	"fmt"
	"testing"

	"codecop.org/scheme/tokenizer"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStringTokenizesIntoEmpty(t *testing.T) {
	cleanedString := ""
	tokens, err := tokenizer.Scan(cleanedString)
	assert.Nil(t, tokens)
	assert.NoError(t, err)
}

func TestTokenizesWithoutErrors(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name           string
		cleanedString  string
		expectedTokens []tokenizer.Token
	}{
		{
			name:           "number 1",
			cleanedString:  "1",
			expectedTokens: []tokenizer.Token{tokenizer.NumberToken{Value: 1}},
		},
		{
			name:           "number 3",
			cleanedString:  "3",
			expectedTokens: []tokenizer.Token{tokenizer.NumberToken{Value: 3}},
		},
		{
			name:           "boolean true",
			cleanedString:  "#t",
			expectedTokens: []tokenizer.Token{tokenizer.BooleanToken{Value: true}},
		},
		{
			name:           "boolean false",
			cleanedString:  "#f",
			expectedTokens: []tokenizer.Token{tokenizer.BooleanToken{Value: false}},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := tokenizer.Scan(tt.cleanedString)
			assert.Equal(tt.expectedTokens, tokens)
			assert.NoError(err)
		})
	}
}

func TestTokenizesWithErrors(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name          string
		cleanedString string
		expectedError error
	}{
		{
			name:          "unhandled token",
			cleanedString: "#tt",
			expectedError: fmt.Errorf("no valid token #tt"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tokenizer.Scan(tt.cleanedString)
			assert.Equal(tt.expectedError, err)
		})
	}
}

func TestScanner(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		name          string
		cleanedString string
		expectedToken tokenizer.Token
		scanner       tokenizer.Scanner
	}{
		{
			name:          "is bool token with value true",
			cleanedString: "#t",
			scanner:       tokenizer.NewBoolScanner(),
			expectedToken: tokenizer.BooleanToken{Value: true},
		},
		{
			name:          "is number token with value 666",
			cleanedString: "666",
			scanner:       tokenizer.NewNumberScanner(),
			expectedToken: tokenizer.NumberToken{Value: 666},
		},
		{
			name:          "is parenthesis token with value (",
			cleanedString: "(",
			scanner:       tokenizer.NewParenthesisScanner(),
			expectedToken: tokenizer.ParenthesisToken{Value: "("},
		},
		{
			name:          "is parenthesis token with value )",
			cleanedString: ")",
			scanner:       tokenizer.NewParenthesisScanner(),
			expectedToken: tokenizer.ParenthesisToken{Value: ")"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			a.True(tt.scanner.IsToken(tt.cleanedString))
			a.Equal(tt.expectedToken, tt.scanner.NewToken(tt.cleanedString))
		})
	}
}
