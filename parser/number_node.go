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

func (nn *NumberNode) addChild(child Ast) error {
    // TODO empty implementations are LSP violations, can we get rid of this method?
	panic("not implemented")
}
