package parser

import "codecop.org/scheme/tokenizer"

type Ast interface {
	GetFirstChild() Ast
	addChild(child Ast) error
}

func Parse(tokens []tokenizer.Token) (Ast, error) {
	root := NewRootNode()
	if len(tokens) > 0 {
		root.addChild(NewBooleanNode(true))
	}
	return root, nil
}
