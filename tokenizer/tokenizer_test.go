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

func TestNumberOneTokenizes(t *testing.T) {
	assert := assert.New(t)
	number := "1"
	tokens, err := tokenizer.Scan(number)
	assert.Equal(
		[]tokenizer.Token{tokenizer.NumberToken{Value: 1}},
		tokens,
	)
	assert.NoError(err)
}

func TestNumberTokenizes(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name           string
		cleanedString  string
		expectedTokens []tokenizer.Token
	}{
		{
			name:           "param 3",
			cleanedString:  "3",
			expectedTokens: []tokenizer.Token{tokenizer.NumberToken{Value: 3}},
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

// TODO: refactor to generalize the function scan
