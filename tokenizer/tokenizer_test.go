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
			name:           "number",
			cleanedString:  "1",
			expectedTokens: []tokenizer.Token{tokenizer.NumberToken{Value: 1}},
		},
		{
			name:           "boolean",
			cleanedString:  "#t",
			expectedTokens: []tokenizer.Token{tokenizer.BooleanToken{Value: true}},
		},
		{
			name:           "parenthesis",
			cleanedString:  "(",
			expectedTokens: []tokenizer.Token{tokenizer.ParenthesisToken{Value: "("}},
		},
		{
			name:           "name",
			cleanedString:  "foo",
			expectedTokens: []tokenizer.Token{tokenizer.NameToken{Value: "foo"}},
		},
		{
			name:           "splitting at the blank",
			cleanedString:  "foo bar",
			expectedTokens: []tokenizer.Token{tokenizer.NameToken{Value: "foo"}, tokenizer.NameToken{Value: "bar"}},
		},
		{
			name:          "splitting by parenthesis",
			cleanedString: "(foo)",
			expectedTokens: []tokenizer.Token{
				tokenizer.ParenthesisToken{Value: "("},
				tokenizer.NameToken{Value: "foo"},
				tokenizer.ParenthesisToken{Value: ")"},
			},
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
