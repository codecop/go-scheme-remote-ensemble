package parser

type FunctionNode struct {
	value string
}

func NewFunctionNode(value string) Ast {
	return &FunctionNode{value: value}
}

func (fn FunctionNode) GetFirstChild() Ast {
	return nil
}

func (fn *FunctionNode) addChild(child Ast) error {
	panic("not implemented")
}
