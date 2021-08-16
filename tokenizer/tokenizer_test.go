package tokenizer_test

import (
	"fmt"
	"testing"

	"codecop.org/scheme/tokenizer"

	"github.com/stretchr/testify/assert"
)

func TestTokenizes(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name           string
		input          string
		expectedTokens []tokenizer.Token
	}{
		{
			name:           "Empty String Tokenizes Into Empty",
			input:          "",
			expectedTokens: nil,
		},
		{
			name:           "number",
			input:          "1",
			expectedTokens: []tokenizer.Token{tokenizer.NumberToken{Value: 1}},
		},
		{
			name:           "boolean",
			input:          "#t",
			expectedTokens: []tokenizer.Token{tokenizer.BooleanToken{Value: true}},
		},
		{
			name:           "parenthesis",
			input:          "(",
			expectedTokens: []tokenizer.Token{tokenizer.ParenthesisToken{Value: "("}},
		},
		{
			name:           "name",
			input:          "foo",
			expectedTokens: []tokenizer.Token{tokenizer.NameToken{Value: "foo"}},
		},
		{
			name:           "splitting at the blank",
			input:          "foo bar",
			expectedTokens: []tokenizer.Token{tokenizer.NameToken{Value: "foo"}, tokenizer.NameToken{Value: "bar"}},
		},
		{
			name:  "splitting by parenthesis",
			input: "(foo)",
			expectedTokens: []tokenizer.Token{
				tokenizer.ParenthesisToken{Value: "("},
				tokenizer.NameToken{Value: "foo"},
				tokenizer.ParenthesisToken{Value: ")"},
			},
		},
		{
			name:  "splitting on mixed space/parenthesis",
			input: "  (\t foo\n ( -1 ))  ",
			expectedTokens: []tokenizer.Token{
				tokenizer.ParenthesisToken{Value: "("},
				tokenizer.NameToken{Value: "foo"},
				tokenizer.ParenthesisToken{Value: "("},
				tokenizer.NumberToken{Value: -1},
				tokenizer.ParenthesisToken{Value: ")"},
				tokenizer.ParenthesisToken{Value: ")"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tokens, err := tokenizer.Scan(testCase.input)
			assert.Equal(testCase.expectedTokens, tokens)
			assert.NoError(err)
		})
	}
}

func TestTokenizesWithErrors(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name          string
		input         string
		expectedError error
	}{
		{
			name:          "unhandled token",
			input:         "#tt",
			expectedError: fmt.Errorf("no valid token #tt"),
		},
		// {
		// 	name:          "unhandled token in middle",
		// 	input:         "(#tt)",
		// 	expectedError: fmt.Errorf("no valid token #tt"),
		// },
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := tokenizer.Scan(testCase.input)
			assert.Equal(testCase.expectedError, err)
		})
	}
}
