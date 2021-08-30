package parser

type RootNode struct {
	children []Ast
}

func NewRootNode() Ast {
	return &RootNode{children: make([]Ast, 0)}
}

func (r RootNode) GetFirstChild() Ast {
	return r.children[0]
}

func (r *RootNode) addChild(child Ast) error {
	r.children = append(r.children, child)
	return nil
}
