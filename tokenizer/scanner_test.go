package tokenizer_test

import (
	"testing"

	"codecop.org/scheme/tokenizer"

	"github.com/stretchr/testify/assert"
)

func TestScanners(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		name          string
		token         string
		usingScanner  tokenizer.Scanner
		expectedToken tokenizer.Token
	}{
		{
			name:          "is bool token with value true",
			token:         "#t",
			usingScanner:  tokenizer.NewBoolScanner(),
			expectedToken: tokenizer.BooleanToken{Value: true},
		},
		{
			name:          "boolean false",
			token:         "#f",
			usingScanner:  tokenizer.NewBoolScanner(),
			expectedToken: tokenizer.BooleanToken{Value: false},
		},
		{
			name:          "number 1",
			token:         "1",
			usingScanner:  tokenizer.NewNumberScanner(),
			expectedToken: tokenizer.NumberToken{Value: 1},
		},
		{
			name:          "number 3",
			token:         "3",
			usingScanner:  tokenizer.NewNumberScanner(),
			expectedToken: tokenizer.NumberToken{Value: 3},
		},
		{
			name:          "is number token with value 666",
			token:         "666",
			usingScanner:  tokenizer.NewNumberScanner(),
			expectedToken: tokenizer.NumberToken{Value: 666},
		},
		{
			name:          "is parenthesis token with value (",
			token:         "(",
			usingScanner:  tokenizer.NewParenthesisScanner(),
			expectedToken: tokenizer.ParenthesisToken{Value: "("},
		},
		{
			name:          "is parenthesis token with value )",
			token:         ")",
			usingScanner:  tokenizer.NewParenthesisScanner(),
			expectedToken: tokenizer.ParenthesisToken{Value: ")"},
		},
		{
			name:          "is name token with value camelCase",
			token:         "camelCase",
			usingScanner:  tokenizer.NewNameScanner(),
			expectedToken: tokenizer.NameToken{Value: "camelCase"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a.True(testCase.usingScanner.IsToken(testCase.token))
			a.Equal(testCase.expectedToken, testCase.usingScanner.NewToken(testCase.token))
		})
	}
}
