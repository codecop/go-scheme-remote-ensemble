package parser

import "codecop.org/scheme/tokenizer"

type Ast interface {
	GetFirstChild() Ast
	addChild(child Ast) error
}

type Root struct {
	children []Ast
}

func NewRoot() Ast {
	return &Root{children: make([]Ast, 0)}
}

func (r Root) GetFirstChild() Ast {
	return r.children[0]
}

func (r *Root) addChild(child Ast) error {
	r.children = append(r.children, child)
	return nil
}

func Parse(tokens []tokenizer.Token) (Ast, error) {
	root := NewRoot()
	if len(tokens) > 0 {
		root.addChild(NewBooleanNode(true))
	}
	return root, nil
}

func NewBooleanNode(value bool) Ast {
	return &BooleanNode{value: value}
}

type BooleanNode struct {
	value bool
}

func (bn BooleanNode) GetFirstChild() Ast {
	return nil
}

func (bn *BooleanNode) addChild(child Ast) error {
	panic("not implemented")
}
