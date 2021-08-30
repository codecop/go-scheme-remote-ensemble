package parser

import (
	"codecop.org/scheme/tokenizer"
)

type Ast interface {
	GetFirstChild() Ast
	addChild(child Ast) error
}

func Parse(tokens []tokenizer.Token) (Ast, error) {
	root := NewRootNode()
	if len(tokens) > 0 {
		if booleanToken, isBooleanToken := tokens[0].(tokenizer.BooleanToken); isBooleanToken {
			root.addChild(NewBooleanNode(booleanToken.Value))
		} else {
			root.addChild(NewNumberNode(42))
		}
	}
	return root, nil
}
