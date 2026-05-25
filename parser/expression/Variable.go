package expression

import "tree-walk-interpreter/token"

type Variable struct {
	Name token.Token
}

func NewVariable(name token.Token) Variable {
	return Variable{
		Name: name,
	}
}

func (v Variable) Accept(visitor ExprVisitor) any {
	return visitor.VisitVariableExpr(v)
}
