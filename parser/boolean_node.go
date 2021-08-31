package parser

type BooleanNode struct {
	value bool
}

func NewBooleanNode(value bool) Ast {
	return &BooleanNode{value: value}
}

func (bn BooleanNode) GetFirstChild() Ast {
	return nil
}

func (bn *BooleanNode) addChild(child Ast) error {
    // TODO empty implementations are LSP violations, can we get rid of this method?
	panic("not implemented")
}
