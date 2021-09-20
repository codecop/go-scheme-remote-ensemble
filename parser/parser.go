package parser

import (
	"fmt"

	"codecop.org/scheme/tokenizer"
)

type ParseError struct {
	message string
}

func NewParseError(message string) ParseError {
	return ParseError{message: message}
}

func (p ParseError) Error() string {
	return p.message
}

type Ast interface {
	// TODO FIRST PRIO add value
	GetFirstChild() Ast
}

type astNode interface {
	Ast
	AddChild(child Ast) error // TODO handle error or remove error on addChild()
}

func Parse(tokens []tokenizer.Token) (Ast, error) {
	root := newRootNode()
	if len(tokens) > 0 {
		if _, ok := tokens[0].(tokenizer.ParenthesisToken); ok {
			nameToken, ok := tokens[1].(tokenizer.NameToken)
			if !ok {
				return nil, ParseError{message: fmt.Sprintf(" %s", nameToken)}
			}
			functionName := nameToken.Value
			functionNode := NewFunctionNode(functionName)
			root.AddChild(functionNode)

			if len(tokens) < 3 {
				return nil, ParseError{message: fmt.Sprintf("missing closing parenthesis after %s", functionName)}
			}
			if _, ok := tokens[2].(tokenizer.ParenthesisToken); !ok {
				functionNode.AddChild(NewNumberNode(1))
			}

		}

		if booleanToken, isBooleanToken := tokens[0].(tokenizer.BooleanToken); isBooleanToken {
			root.AddChild(NewBooleanNode(booleanToken.Value))
		}

		if numberToken, isNumberToken := tokens[0].(tokenizer.NumberToken); isNumberToken {
			root.AddChild(NewNumberNode(numberToken.Value))
		}

	}
	return root, nil
}
