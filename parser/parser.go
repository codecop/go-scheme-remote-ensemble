package parser

import "codecop.org/scheme/tokenizer"

type Ast interface {
}

type Root struct {
}

func NewRoot() Ast {
	return Root{}
}

func Parse(tokens []tokenizer.Token) (Ast, error) {
	return NewRoot(), nil
}
