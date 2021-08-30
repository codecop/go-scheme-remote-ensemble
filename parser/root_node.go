package parser

type Root struct {
	children []Ast
}

func NewRoot() Ast {
	return &Root{children: make([]Ast, 0)}
}

func (r Root) GetFirstChild() Ast {
	return r.children[0]
}

func (r *Root) addChild(child Ast) error {
	r.children = append(r.children, child)
	return nil
}
