package parser

type BooleanNode struct {
	value bool
}

func NewBooleanNode(value bool) Ast {
	return &BooleanNode{value: value}
}

func (bn BooleanNode) GetFirstChild() Ast {
	// TODO re-evaluate interface definition for unspecific structs
	return nil
}
