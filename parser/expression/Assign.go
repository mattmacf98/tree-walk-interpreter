package expression

import "tree-walk-interpreter/token"

type Assign struct {
	Name  token.Token
	Value Expr
}

func NewAssign(name token.Token, value Expr) Assign {
	return Assign{
		Name:  name,
		Value: value,
	}
}

func (a Assign) Accept(visitor ExprVisitor) any {
	return visitor.VisitAssignExpr(a)
}
