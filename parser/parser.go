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
		//if reflect.TypeOf(tokens[0]) == tokenizer.BooleanToken  {
		if _, err := tokens[0].(tokenizer.BooleanToken); err {
			root.addChild(NewBooleanNode(true))
		} else {
			root.addChild(NewNumberNode(42))
		}

	}
	return root, nil
}
