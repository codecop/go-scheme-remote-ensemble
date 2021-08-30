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

	assert.Equal(parser.NewRootNode(), ast)
	assert.NoError(err)
}

func TestBooleanTokenParsesIntoBooleanNode(t *testing.T) {
	assert := assert.New(t)

	tokens := []tokenizer.Token{tokenizer.NewBoolScanner().NewToken("#t")}
	ast, err := parser.Parse(tokens)

	node := ast.GetFirstChild()
	assert.Equal(parser.NewBooleanNode(true), node)
	// TODO assert no other children
	assert.NoError(err)
}

// TODO (later) make these tests pass to continue with logic

// 	Function Evaluation
// 	tokens := []tokenizer.Token{
// 		tokenizer.NewParenthesisScanner().NewToken("("),
// 		tokenizer.NewNameScanner().NewToken("foo"),
// 		tokenizer.NewParenthesisScanner().NewToken(")"),
// 	}

// 	Function Evaluation With Arguments
// 	tokens := []tokenizer.Token{
// 		tokenizer.NewParenthesisScanner().NewToken("("),
// 		tokenizer.NewNameScanner().NewToken("plus"),
// 		tokenizer.NewNumberScanner().NewToken("1"),
// 		tokenizer.NewNumberScanner().NewToken("2"),
// 		tokenizer.NewParenthesisScanner().NewToken(")"),
// 	}
