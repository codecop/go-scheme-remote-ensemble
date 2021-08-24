package parser_test

import (
	"testing"

	"codecop.org/scheme/parser"
	"codecop.org/scheme/tokenizer"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTokensParseIntoEmpty(t *testing.T) {
	assert := assert.New(t)

	ast, err := parser.Parse([]tokenizer.Token{})

	assert.Equal(parser.NewRoot(), ast)
	assert.NoError(err)
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

// TODO make these tests pass

// 	name: "Function evaluation (foo)",
// 	tokens: []tokenizer.Token{
// 		tokenizer.NewParenthesisToken("("),
// 		tokenizer.NewNameToken("foo"),
// 		tokenizer.NewParenthesisToken(")"),
// 	},

// 	name: "Function evaluation with arguments (+ 1 2)",
// 	tokens: []tokenizer.Token{
// 		tokenizer.NewParenthesisToken("("),
// 		tokenizer.NewNameToken("plus"),
// 		tokenizer.NewNumberToken("1"),
// 		tokenizer.NewNumberToken("2"),
// 		tokenizer.NewParenthesisToken(")"),
// 	},
