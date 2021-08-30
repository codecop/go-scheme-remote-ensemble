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

func TestBooleanTokenParsesIntoTrueBooleanNode(t *testing.T) {
	assert := assert.New(t)

	tokens := []tokenizer.Token{tokenizer.NewBoolScanner().NewToken("#t")}
	ast, err := parser.Parse(tokens)

	node := ast.GetFirstChild()
	assert.Equal(parser.NewBooleanNode(true), node)
	// TODO assert no other children
	assert.NoError(err)
}

func TestBooleanTokenParsesIntoFalseBooleanNode(t *testing.T) {
	assert := assert.New(t)

	tokens := []tokenizer.Token{tokenizer.NewBoolScanner().NewToken("#f")}
	ast, _ := parser.Parse(tokens)

	node := ast.GetFirstChild()
	assert.Equal(parser.NewBooleanNode(false), node)
}

func TestNumberTokenParsesIntoNumberNode(t *testing.T) {
	assert := assert.New(t)

	tokens := []tokenizer.Token{tokenizer.NewNumberScanner().NewToken("42")}
	ast, err := parser.Parse(tokens)

	node := ast.GetFirstChild()
	assert.Equal(parser.NewNumberNode(42), node)
	// TODO assert no other children
	assert.NoError(err)
}

// TODO (later) make these tests pass to continue with logic

func TestParsesFunctionCall(t *testing.T) {
	t.Skip()

	assert := assert.New(t)

	tokens := []tokenizer.Token{
		tokenizer.NewParenthesisScanner().NewToken("("),
		tokenizer.NewNameScanner().NewToken("list"),
		tokenizer.NewParenthesisScanner().NewToken(")"),
	}
	ast, err := parser.Parse(tokens)

	node := ast.GetFirstChild()
	assert.Equal(parser.NewFunctionNode("list"), node)

	// TODO assert node to have no children (due to arguments)
	assert.NoError(err)
}

// 	Function Evaluation
// 	tokens := []tokenizer.Token{
// 		tokenizer.NewParenthesisScanner().NewToken("("),
// 		tokenizer.NewNameScanner().NewToken("list"),
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
