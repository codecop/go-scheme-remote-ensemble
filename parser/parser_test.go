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
		// TODO make these tests pass
		{
			name:        "Empty tokens parses into empty",
			tokens:      []tokenizer.Token{},
			expectedAst: parser.NewRoot(),
		},
		// root -> slice -> nodes -> slice -> boolean node
		// {
		// 	name: "Function evaluation (foo)",
		// 	tokens: []tokenizer.Token{
		// 		tokenizer.NewParenthesisToken("("),
		// 		tokenizer.NewNameToken("foo"),
		// 		tokenizer.NewParenthesisToken(")"),
		// 	},
		// 	expectedAst: nil,
		// },
		// {
		// 	name: "Function evaluation with arguments (+ 1 2)",
		// 	tokens: []tokenizer.Token{
		// 		tokenizer.NewParenthesisToken("("),
		// 		tokenizer.NewNameToken("plus"),
		// 		tokenizer.NewNumberToken("1"),
		// 		tokenizer.NewNumberToken("2"),
		// 		tokenizer.NewParenthesisToken(")"),
		// 	},
		// 	expectedAst: nil,
		// },
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ast, err := parser.Parse(testCase.tokens)
			assert.Equal(testCase.expectedAst, ast)
			assert.NoError(err)
		})
	}
}

func TestBooleanTokenParsesIntoBooleanNode(t *testing.T) {
	assert := assert.New(t)

	tokens := []tokenizer.Token{tokenizer.NewBooleanToken("#t")}
	ast, err := parser.Parse(tokens)

	node := ast.GetFirstChild()
	assert.Equal(parser.NewBooleanNode(true), node)
	// assert no other children
	assert.NoError(err)
}
