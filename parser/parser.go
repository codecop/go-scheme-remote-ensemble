package parser

import (
	"fmt"

	"codecop.org/scheme/tokenizer"
)

type Ast interface {
	GetFirstChild() Ast
	addChild(child Ast) error // TODO handle error or remove error on addChild()
}

func Parse(tokens []tokenizer.Token) (Ast, error) {
	root := NewRootNode()
	if len(tokens) > 0 {
		if _, ok := tokens[0].(tokenizer.ParenthesisToken); ok {
			functionName := tokens[1].(tokenizer.NameToken).Value
			root.addChild(NewFunctionNode(functionName))
			if len(tokens) < 3 {
				return nil, fmt.Errorf("missing closing parenthesis after %s", "list")
			}
		}

		if booleanToken, isBooleanToken := tokens[0].(tokenizer.BooleanToken); isBooleanToken {
			root.addChild(NewBooleanNode(booleanToken.Value))
		}

		if numberToken, isNumberToken := tokens[0].(tokenizer.NumberToken); isNumberToken {
			root.addChild(NewNumberNode(numberToken.Value))
		}

	}
	return root, nil
}
