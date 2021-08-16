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
			name:          "boolean false",
			cleanedString: "#f",
			scanner:       tokenizer.NewBoolScanner(),
			expectedToken: tokenizer.BooleanToken{Value: false},
		},
		{
			name:          "number 1",
			cleanedString: "1",
			scanner:       tokenizer.NewNumberScanner(),
			expectedToken: tokenizer.NumberToken{Value: 1},
		},
		{
			name:          "number 3",
			cleanedString: "3",
			scanner:       tokenizer.NewNumberScanner(),
			expectedToken: tokenizer.NumberToken{Value: 3},
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
		{
			name:          "is name token with value camelCase",
			cleanedString: "camelCase",
			scanner:       tokenizer.NewNameScanner(),
			expectedToken: tokenizer.NameToken{Value: "camelCase"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			a.True(tt.scanner.IsToken(tt.cleanedString))
			a.Equal(tt.expectedToken, tt.scanner.NewToken(tt.cleanedString))
		})
	}
}
