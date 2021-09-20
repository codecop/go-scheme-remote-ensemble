package parser

type FunctionNode struct {
	value    string
	children []Ast
}

func NewFunctionNode(value string) astNode {
	return &FunctionNode{value: value}
}

func (fn FunctionNode) GetFirstChild() Ast {
	if len(fn.children) > 0 {
		return fn.children[0]
	}
	return nil
}

func (fn *FunctionNode) AddChild(child Ast) error {
	fn.children = append(fn.children, child)
	return nil
}
