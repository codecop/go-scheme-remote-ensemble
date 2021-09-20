package parser

type NumberNode struct {
	value int
}

func NewNumberNode(value int) Ast {
	return &NumberNode{value: value}
}

func (nn NumberNode) GetFirstChild() Ast {
	return nil
}
