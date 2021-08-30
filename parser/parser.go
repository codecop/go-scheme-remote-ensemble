package parser

import "codecop.org/scheme/tokenizer"

type Ast interface {
	GetFirstChild() Ast
	addChild(child Ast) error
}

// TODO cleanup Parser: inconsistent node names: Root vs. RootNode

func Parse(tokens []tokenizer.Token) (Ast, error) {
	root := NewRoot()
	if len(tokens) > 0 {
		root.addChild(NewBooleanNode(true))
	}
	return root, nil
}
