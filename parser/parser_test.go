package parser_test

import (
	"testing"

	"codecop.org/scheme/parser"
	"codecop.org/scheme/tokenizer"

	"github.com/stretchr/testify/assert"
)

func TestParses(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name        string
		tokens      []tokenizer.Token
		expectedAst parser.Ast
	}{
		{
			name:        "Empty tokens parses into empty",
			tokens:      []tokenizer.Token{},
			expectedAst: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ast, err := parser.Parse(testCase.tokens)
			assert.Equal(testCase.expectedAst, ast)
			assert.NoError(err)
		})
	}
}
