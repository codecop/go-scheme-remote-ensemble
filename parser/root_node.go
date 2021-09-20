package parser

type RootNode struct {
	children []Ast
}

func newRootNode() astNode {
	return &RootNode{children: make([]Ast, 0)}
}

func (r RootNode) GetFirstChild() Ast {
	if len(r.children) > 0 {
		return r.children[0]
	}
	return nil
}

func (r *RootNode) AddChild(child Ast) error {
	r.children = append(r.children, child)
	return nil
}
