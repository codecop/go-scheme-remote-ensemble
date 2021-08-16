package parser

import "codecop.org/scheme/tokenizer"

type Ast interface {
}

func Parse(tokens []tokenizer.Token) (Ast, error) {
	return nil, nil
}
